/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:28 AM
 */

package workflow

import (
	"fmt"
	"sync"
	"time"
)

type job func()

var (
	PushQueueTimeout = fmt.Errorf("push queue timeout")
)

type JobQueue interface {
	PushJob(j job) error
}

type DefaultJobQueue struct {
	queue  chan job
	once   sync.Once
	option *QueueOption
}

type QueueOption struct {
	WorkerCount int
	QueueSize   int
	PushTimeout time.Duration
}

func NewDefaultJobQueue(option *QueueOption) *DefaultJobQueue {
	queue := &DefaultJobQueue{
		queue:  make(chan job, option.QueueSize),
		option: option,
	}

	return queue
}

func (queue *DefaultJobQueue) Start() error {
	for i := 0; i < queue.option.WorkerCount; i++ {
		go func() {
			for {
				select {
				case job, ok := <-queue.queue:
					if !ok {
						return
					}

					job()
				}
			}
		}()
	}

	return nil
}

func (queue *DefaultJobQueue) PushJob(j job) error {
	timer := time.NewTimer(queue.option.PushTimeout)

	select {
	case queue.queue <- j:
		return nil
	case <-timer.C:
		return PushQueueTimeout
	}
}

func (queue *DefaultJobQueue) Stop() {
	queue.once.Do(func() {
		close(queue.queue)
	})
}
