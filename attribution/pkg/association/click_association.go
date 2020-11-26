package association

import (
	"sort"

	"github.com/TencentAd/attribution/attribution/pkg/association/validation"
	"github.com/TencentAd/attribution/attribution/pkg/common/key"
	"github.com/TencentAd/attribution/attribution/pkg/data/user"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"
	"github.com/TencentAd/attribution/attribution/proto/conv"
)

// 将点击信息关联到转化日志
type ClickAssociation struct {
	clickIndex clickindex.ClickIndex
	validation validation.ClickLogValidation
}

func NewClickAssociation() *ClickAssociation {
	return &ClickAssociation{}
}

func (ass *ClickAssociation) WithClickIndex(clickIndex clickindex.ClickIndex) *ClickAssociation {
	ass.clickIndex = clickIndex
	return ass
}

func (ass *ClickAssociation) WithValidation(validation validation.ClickLogValidation) *ClickAssociation {
	ass.validation = validation
	return ass
}

type AssocContext struct {
	ConvLog     *conv.ConversionLog // 转化数据
	Candidates  []*conv.MatchClick  // 关联上的点击数据
	SelectClick *conv.MatchClick    // 最终选择的点击数据
}

func NewAssocContext(convLog *conv.ConversionLog) *AssocContext {
	return &AssocContext{
		ConvLog: convLog,
	}
}

func (ass *ClickAssociation) Association(c *AssocContext) error {
	if err := ass.queryClicks(c); err != nil {
		return err
	}
	ass.filterInvalidClick(c)
	ass.selectOneClick(c)

	c.ConvLog.MatchClick = c.SelectClick

	return nil
}

func (ass *ClickAssociation) queryClicks(c *AssocContext) error {
	convLog := c.ConvLog
	uids, err := user.GenerateNormalIdsByPriority(convLog.UserData)
	if err != nil {
		return err
	}

	candidates := make([]*conv.MatchClick, 0)
	for _, uid := range uids {
		v, err := ass.clickIndex.Get(uid.T, key.FormatClickLogKey(convLog.AppId, uid.Id))
		if err != nil {
			return err
		}
		if v != nil {
			candidates = append(candidates, &conv.MatchClick{
				ClickLog:    v,
				MatchIdType: uid.T,
			})
		}
	}
	c.Candidates = candidates
	return nil
}

func (ass *ClickAssociation) filterInvalidClick(c *AssocContext) {
	valid := make([]*conv.MatchClick, 0, len(c.Candidates))
	convLog := c.ConvLog
	for _, cand := range c.Candidates {
		if ass.validation.Check(convLog, cand.ClickLog) {
			valid = append(valid, cand)
		}
	}
	c.Candidates = valid
}

func (ass *ClickAssociation) selectOneClick(c *AssocContext) {
	if len(c.Candidates) == 0 {
		return
	}
	sort.Slice(c.Candidates, func(i, j int) bool {
		return c.Candidates[i].ClickLog.ClickTime > c.Candidates[j].ClickLog.ClickTime
	})
	c.SelectClick = c.Candidates[0]
}
