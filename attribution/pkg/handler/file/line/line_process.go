/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/9/20, 5:26 PM
 */

package line

import (
	"bufio"
	"errors"
	"os"
	"sync"
	"time"
)

var (
	ErrQueueTimeout = errors.New("queue full")
)

type LineProcess struct {
	filename    string
	processFunc Func
	errCallback ErrCallback

	parallelism  int
	queueTimeout time.Duration
	queueSize    int
	dataQueue    chan string
	once         sync.Once

	doneWG sync.WaitGroup
}

type Func func(line string) error
type ErrCallback func(line string, err error)

const defaultParallelism = 5
const defaultQueueSize = 100
const defaultQueueTimeout = time.Millisecond * 10

func NewLineProcess(filename string, processFunc Func, errCallback ErrCallback) *LineProcess {
	return &LineProcess{
		filename:     filename,
		processFunc:  processFunc,
		errCallback:  errCallback,
		parallelism:  defaultParallelism,
		queueSize:    defaultQueueSize,
		queueTimeout: defaultQueueTimeout,
	}
}

func (p *LineProcess) WithParallelism(parallel int) *LineProcess {
	p.parallelism = parallel
	return p
}

func (p *LineProcess) WithQueueTimeout(timeout time.Duration) *LineProcess {
	p.queueTimeout = timeout
	return p
}

func (p *LineProcess) WithQueueSize(size int) *LineProcess {
	p.queueSize = size
	return p
}

func (p *LineProcess) init() error {
	p.dataQueue = make(chan string, p.queueSize)
	return p.startWorker()
}

func (p *LineProcess) Run() error {
	if err := p.init(); err != nil {
		return err
	}
	defer p.close()

	file, err := os.Open(p.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err = p.addLine(line); err != nil {
			return err
		}
	}
	return nil
}

func (p *LineProcess) startWorker() error {
	p.doneWG.Add(p.parallelism)
	for i := 0; i < p.parallelism; i++ {
		go func() {
			defer p.doneWG.Done()
			for line := range p.dataQueue {

				if err := p.processFunc(line); err != nil && p.errCallback != nil {
					p.errCallback(line, err)
				}
			}
		}()
	}
	return nil
}

func (p *LineProcess) addLine(line string) error {
	if line == "" {
		return nil
	}

	tm := time.NewTimer(p.queueTimeout)

	select {
	case p.dataQueue <- line:
		return nil
	case <-tm.C:
		return ErrQueueTimeout
	}
}

func (p *LineProcess) close() {
	p.once.Do(
		func() {
			close(p.dataQueue)
		})
}

func (p *LineProcess) WaitDone() {
	p.doneWG.Wait()
}
