/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:28 AM
 */

package workflow

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Add struct {
	i int32
}

func (r *Add) Run(i interface{}) {
	atomic.AddInt32(i.(*int32), r.i)
}

type Multiply struct {
	i int32
}

func (r *Multiply) Run(i interface{}) {
	//time.Sleep(time.Second)
	*(i.(*int32)) = *(i.(*int32)) * r.i
}

type Print struct {
}

func (r *Print) Run(i interface{}) {
	//fmt.Println(*i.(*int32))
}

func TestWorkFlow(t *testing.T) {
	add1 := &Add{i: 1,}
	add2 := &Add{i: 2,}
	multiply := &Multiply{i: 3}

	add1Task := &TaskNode{
		Task: add1,
	}

	add2Task := &TaskNode{
		Task: add2,
	}

	multiplyTask := &TaskNode{
		Task: multiply,
	}

	printTask := &TaskNode{
		Task: &Print{},
	}
	add1Task2 := NewTaskNode(add1)
	add2Task2 := NewTaskNode(add2)

	// 实现 (i + 1 + 2) * 3 + 1 + 2，然后打印
	wf := NewWorkFlow()
	wf.AddStartNode(add1Task)
	wf.AddStartNode(add2Task)
	wf.AddEdge(add1Task, multiplyTask)
	wf.AddEdge(add2Task, multiplyTask)
	wf.AddEdge(multiplyTask, add1Task2)
	wf.AddEdge(multiplyTask, add2Task2)
	wf.AddEdge(add1Task2, printTask)
	wf.AddEdge(add2Task2, printTask)
	wf.ConnectToEnd(printTask)

	isDAG := wf.CheckDAG()
	assert.True(t, isDAG)

	if !wf.CheckDAG() {
		t.Errorf("check DAG failed")
	}

	var i int32 = 10
	wf.Start(&i)
	wf.WaitDone()

	wf.AddEdge(printTask, multiplyTask)
	isDAG = wf.CheckDAG()
	assert.False(t, isDAG)
}

func BenchmarkTaskNode_ExecuteTask(b *testing.B) {
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			wf := buildWorkflow()
			var p int32 = 10
			wf.Start(&p)
			wf.WaitDone()
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkTaskNode_SubmitTask(b *testing.B) {
	//b.ResetTimer()
	queue := NewDefaultJobQueue(&QueueOption{
		WorkerCount: 50,
		QueueSize:   1000,
		PushTimeout: time.Second * 3,
	})
	queue.Start()
	defer queue.Stop()
	for i := 0; i < b.N; i++ {
		wf := buildWorkflow()

		var p int32 = 10
		wf.StartWithJobQueue(queue, context.Background(), &p)
		wf.WaitDone()
	}
}

func TestWorkFlowWithContext(t *testing.T) {
	var i int32 = 10
	var wf *WorkFlow
	var ctx context.Context
	var cancel context.CancelFunc

	wf = buildWorkflow()
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*50)
	time.Sleep(time.Millisecond * 60)
	wf.StartWithContext(ctx, &i)
	wf.WaitDone()
	assert.True(t, ctx.Err() == context.DeadlineExceeded)
	assert.Error(t, ctx.Err())
	cancel()

	wf = buildWorkflow()
	ctx, cancel = context.WithCancel(context.Background())
	wf.StartWithContext(ctx, &i)
	cancel()
	wf.WaitDone()
	assert.True(t, ctx.Err() == context.Canceled)
	assert.Error(t, ctx.Err())
	cancel()
}

func buildWorkflow() *WorkFlow {
	add1 := &Add{i: 1,}
	add2 := &Add{i: 2,}
	multiply := &Multiply{i: 3}

	add1Task := &TaskNode{
		Task: add1,
	}

	add2Task := &TaskNode{
		Task: add2,
	}

	multiplyTask := &TaskNode{
		Task: multiply,
	}

	printTask := &TaskNode{
		Task: &Print{},
	}
	add1Task2 := NewTaskNode(add1)
	add2Task2 := NewTaskNode(add2)

	// 实现 (i + 1 + 2) * 3 + 1 + 2，然后打印
	wf := NewWorkFlow()
	wf.AddStartNode(add1Task)
	wf.AddStartNode(add2Task)
	wf.AddEdge(add1Task, multiplyTask)
	wf.AddEdge(add2Task, multiplyTask)
	wf.AddEdge(multiplyTask, add1Task2)
	wf.AddEdge(multiplyTask, add2Task2)
	wf.AddEdge(add1Task2, printTask)
	wf.AddEdge(add2Task2, printTask)
	wf.ConnectToEnd(printTask)

	return wf
}
