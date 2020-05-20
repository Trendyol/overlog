package over

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	_globalMdc  = InitGlobalMdcAdapter()
	_MdcAdapter *MdcAdapter
)

type MdcAdapter struct {
	items map[string]interface{}
	sync.RWMutex
}

func InitGlobalMdcAdapter() *MdcAdapter {
	if _MdcAdapter == nil {
		_MdcAdapter = &MdcAdapter{items: make(map[string]interface{})}
	}

	return _MdcAdapter
}

func ResetGlobalMdcAdapter() {
	_MdcAdapter.RLock()
	_MdcAdapter.items = make(map[string]interface{})
	_MdcAdapter.RUnlock()
}

func MDC() *MdcAdapter {
	_MdcAdapter.RLock()
	s := _globalMdc
	_MdcAdapter.RUnlock()
	return s
}

func (m *MdcAdapter) Set(key string, value interface{}) {
	uniqueKey := m.getUniqueKey(key)
	m.Lock()
	m.items[uniqueKey] = value
	m.Unlock()
}

func (m *MdcAdapter) SetMap(data map[string]interface{}) {
	for key, value := range data {
		uniqueKey := m.getUniqueKey(key)
		m.Lock()
		m.items[uniqueKey] = value
		m.Unlock()
	}
}

func (m *MdcAdapter) SetIfAbsent(key string, value interface{}) bool {
	uniqueKey := m.getUniqueKey(key)
	m.Lock()
	_, ok := m.items[uniqueKey]
	if !ok {
		m.items[uniqueKey] = value
	}
	m.Unlock()
	return !ok
}

func (m *MdcAdapter) Get(key string) (interface{}, bool) {
	uniqueKey := m.getUniqueKey(key)
	m.RLock()
	v, ok := m.items[uniqueKey]
	m.RUnlock()
	return v, ok
}

func (m *MdcAdapter) GetString(key string) string {
	uniqueKey := m.getUniqueKey(key)
	m.RLock()
	v, ok := m.items[uniqueKey]
	m.RUnlock()
	if !ok {
		v = ""
	}
	return fmt.Sprintf("%v", v)
}

func (m *MdcAdapter) Count() int {
	m.RLock()
	count := len(m.items)
	m.RUnlock()
	return count
}

func (m *MdcAdapter) Has(key string) bool {
	uniqueKey := m.getUniqueKey(key)
	m.RLock()
	_, ok := m.items[uniqueKey]
	m.RUnlock()
	return ok
}

func (m *MdcAdapter) Remove(key string) {
	uniqueKey := m.getUniqueKey(key)
	m.Lock()
	delete(m.items, uniqueKey)
	m.Unlock()
}

func (m *MdcAdapter) IsEmpty() bool {
	m.RLock()
	count := len(m.items)
	m.RUnlock()
	return count == 0
}

func (m *MdcAdapter) Keys() []string {
	count := m.Count()
	ch := make(chan string, count)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func(m *MdcAdapter) {
			m.RLock()
			for key := range m.items {
				ch <- key
			}
			m.RUnlock()
			wg.Done()
		}(m)
		wg.Wait()
		close(ch)
	}()

	keys := make([]string, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

func (m *MdcAdapter) getUniqueKey(key string) string {
	if key == "" {
		panic("MDC key cannot be empty")
	}
	return key + "-" + strconv.FormatUint(GetGoroutineID(), 10)
}
