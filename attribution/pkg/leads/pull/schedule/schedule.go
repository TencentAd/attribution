/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/15/20, 11:17 AM
 */

package schedule

type Task struct {
	meta *TaskMeta
}

func NewTask(meta *TaskMeta) *Task {
	return &Task{
		meta: meta,
	}
}

func (t *Task) Start() error {
	return nil
}

func (t *Task) Stop() error {
	return nil
}

func (t *Task) FlushStatus() error {

	return nil
}

type TaskMeta struct {

}

type TaskProgress struct {
	CompleteTimestamp int64 // 已经拉取完成的时间戳
}

type TaskManager struct {

}

type
