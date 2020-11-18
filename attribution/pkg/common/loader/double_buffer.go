package loader

import (
	"sync"
	"sync/atomic"
	"time"
)

var m FileLoadManager

type FileLoader interface {
	DetectNewFile() (string, bool)
	Load(filePath string, i interface{}) error
	Alloc() interface{}
	Reset()
}

type allocFunc func() interface{}

type FileDoubleBuffer struct {
	Loader     FileLoader
	bufferData []interface{}
	curIndex   int32
	notify     func(error)
}

func NewFileDoubleBuffer(loader FileLoader) *FileDoubleBuffer {
	b := &FileDoubleBuffer{
		Loader:   loader,
		curIndex: 0,
	}
	b.bufferData = append(b.bufferData, loader.Alloc(), loader.Alloc())
	register(b)
	return b
}

func (b *FileDoubleBuffer) SetNotify(fn func(error))  {
	b.notify = fn
}

func (b *FileDoubleBuffer) load() {
	file, newFound := b.Loader.DetectNewFile()
	if newFound {
		ci := 1 - atomic.LoadInt32(&b.curIndex)
		err := b.Loader.Load(file, b.bufferData[ci])
		if err == nil {
			atomic.StoreInt32(&b.curIndex, ci)
		}

		if b.notify != nil {
			b.notify(err)
		}
	}
}

func (b *FileDoubleBuffer) Data() interface{} {
	ci := atomic.LoadInt32(&b.curIndex)
	return b.bufferData[ci]
}

type FileLoadManager struct {
	doubleBuffers []*FileDoubleBuffer
	mutex         sync.Mutex
}

func (m *FileLoadManager) reset() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, b := range m.doubleBuffers {
		b.Loader.Reset()
		b.load()
	}
}

func (m *FileLoadManager) load() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, b := range m.doubleBuffers {
		b.load()
	}
}

func (m *FileLoadManager) reload(reloadInterval int) {
	for {
		time.Sleep(time.Duration(reloadInterval) * time.Second)
		m.load()
	}
}

func register(doubleBuffer *FileDoubleBuffer) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.doubleBuffers = append(m.doubleBuffers, doubleBuffer)
}

func ResetDoubleBuffer() {
	m.reset()
}

func StartDoubleBufferLoad(reloadInterval int) {
	m.load()
	go m.reload(reloadInterval)
}
