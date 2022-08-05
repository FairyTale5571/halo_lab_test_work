package memory

import (
	"time"

	"github.com/fairytale5571/privat_test/pkg/errs"
)

const (
	ttl = 15 * time.Second
)

type Memory struct {
	data map[string]interface{}

	stop chan struct{}
}

func New() *Memory {
	return &Memory{
		data: make(map[string]interface{}),
	}
}

func (m *Memory) Get(key string) (interface{}, error) {
	value, ok := m.data[key]
	if !ok {
		return nil, errs.ErrorNotFound
	}
	return value, nil
}

func (m *Memory) Set(key string, value interface{}, ttl int64) error {
	m.data[key] = value
	return nil
}

func (m *Memory) Delete(key string) error {
	delete(m.data, key)
	return nil
}

func (m *Memory) CleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			m.Cleanup()
		case <-m.stop:
			return
		}
	}
}

func (m *Memory) Cleanup() {
}
