/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 2:34 PM
 */

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/oauth"
	"github.com/TencentAd/attribution/attribution/pkg/storage/leads"

	"github.com/avast/retry-go"
	"github.com/golang/glog"
)

// 负责从线索平台拉取线索信息
type LeadsPullClient struct {
	config     *PullConfig
	httpClient *http.Client
	storage    leads.Storage
	token      *oauth.Token

	// 进度
	beginTime    time.Time     // 拉取开始的时间
	endTime      time.Time     // 拉取结束的时间戳
	pullInterval time.Duration // 拉取间隔
	intervalDone bool          // 拉取时间段已经完成

	nextPage     int       // 下次拉取第几页
	currentCount int       // 当前时间段总共的线索数
	lastSearch   [2]string // 最后查到的线索，用户深度翻页
}

type PullConfig struct {
	AccountId int64                    `json:"account_id"`
	Filtering []map[string]interface{} `json:"filtering"`
	PageSize  int                      `json:"page_size"`
}

const defaultPageSize = 100
func NewPullConfig(conf string) (*PullConfig, error) {
	var pc PullConfig
	if err := json.Unmarshal([]byte(conf), &pc); err != nil {
		return nil, err
	}

	if pc.PageSize == 0 {
		pc.PageSize = defaultPageSize
	}

	return &pc, nil
}

func NewLeadsPullClient(config *PullConfig) *LeadsPullClient {
	return &LeadsPullClient{
		config:       config,
		httpClient:   &http.Client{},
		beginTime:    time.Unix(0, 0),
		endTime:      time.Now(),
		intervalDone: false,
		nextPage:     0,
		currentCount: 0,
		lastSearch:   [2]string{"", ""},
	}
}
func (c *LeadsPullClient) WithStorage(storage leads.Storage) *LeadsPullClient {
	c.storage = storage
	return c
}
func (c *LeadsPullClient) WithPullInterval(interval time.Duration) *LeadsPullClient {
	c.pullInterval = interval
	return c
}

func (c *LeadsPullClient) formatRequestUrl() (string, error) {
	// token := c.token.Get()
	//
	// return fmt.Sprintf(`https://api.e.qq.com/v1.3/lead_clues/get?access_token=%s&timestamp=%d&nonce=%s`,
	// 	token, time.Now().Unix(), oauth.GenNonce()), nil
	return "", nil
}

func (c *LeadsPullClient) Pull() error {
	if err := c.pullRoutine(); err != nil {
		return err
	}
	return nil
}

func (c *LeadsPullClient) pullRoutine() error {
	for {
		if time.Since(c.beginTime) >= c.pullInterval {
			if err := c.requestInterval(); err != nil {
				glog.Errorf("failed to pull interval, begin: %v, end: %v", c.beginTime, c.endTime)
			} else {
				c.onIntervalDone()
			}
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (c *LeadsPullClient) onIntervalDone() error {
	c.beginTime = c.endTime
	c.endTime = c.endTime.Add(c.pullInterval)

	c.nextPage = 0
	c.currentCount = 0
	c.intervalDone = false
	c.lastSearch = [2]string{"", ""}
	return nil
}

func (c *LeadsPullClient) requestInterval() error {
	c.intervalDone = false
	for !c.intervalDone {
		if err := c.requestPageWithRetry(); err != nil {
			glog.Errorf("failed to request page, err: %v", err)
			return err
		}
	}

	return nil
}

func (c *LeadsPullClient) requestPageWithRetry() error {
	return retry.Do(
		func() error {
			return c.requestPage()
		},
		retry.Attempts(3))
}

func (c *LeadsPullClient) requestPage() error {
	url, err := c.formatRequestUrl()
	if err != nil {
		return err
	}

	config := c.config
	req := &protocal.LeadsRequest{
		AccountId: config.AccountId,
		TimeRange: &protocal.TimeRange{
			StartTime: c.beginTime.Unix(),
			EndTime:   c.endTime.Unix(),
		},
		Filtering: config.Filtering,
		PageSize:  config.PageSize,
	}

	if c.currentCount+config.PageSize <= 5000 {
		req.Page = c.nextPage
	} else {
		req.LastSearchAfterValues = c.lastSearch
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
	if err != nil {
		glog.Errorf("failed to create request, err: %v", err)
		return err
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	if httpResp.StatusCode != 200 {
		return fmt.Errorf("http response status[%s] not valid", httpResp.Status)
	}
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	var resp protocal.LeadsResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("leads response code[%d] not valid, msg: %s", resp.Code, resp.Message)
	}
	if len(resp.Data) == 0 {
		return errors.New("leads data empty")
	}
	if resp.PageInfo == nil {
		return errors.New("response page info nil")
	}
	if err := c.store(&resp); err != nil {
		return err
	}
	c.onPageRequestSuccess(&resp)
	return nil
}

func (c *LeadsPullClient) store(resp *protocal.LeadsResponse) error {
	for _, info := range resp.Data {
		if err := c.storage.Store(info, time.Hour * time.Duration(*define.LeadsExpireHour)); err != nil {
			return err
		}
	}
	return nil
}

func (c *LeadsPullClient) onPageRequestSuccess(resp *protocal.LeadsResponse) {
	last := resp.Data[len(resp.Data)-1]
	c.lastSearch = [2]string{
		strconv.FormatInt(last.LeadsActionTime, 64),
		strconv.FormatInt(last.LeadsId, 64),
	}

	c.nextPage++

	// 检查是否已经拉取完成
	c.currentCount += len(resp.Data)
	if c.currentCount == resp.PageInfo.TotalNumber {
		c.intervalDone = true
	}
}
