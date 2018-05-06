package gotool

import "sync"

type Array struct {
	data []interface{}
	lock sync.Mutex
}

func NewArray(data []interface{}) *Array {
	arr := &Array{}
	if data != nil {
		arr.data = data
	}
	return arr
}

func (m *Array) Length() int {
	return len(m.data)
}

func (m *Array) ToSlice() []interface{} {
	return m.data
}

func (m *Array) Source() interface{} {
	return m.data
}

func (m *Array) Search(fn func(item interface{}, index int) bool) ([]int) {
	defer m.lock.Unlock()
	m.lock.Lock()

	var indices []int
	for index, item := range m.data {
		if fn(item, index) {
			indices = append(indices, index)
			break
		}
	}

	return indices
}

func (m *Array) Push(item interface{}, others ...interface{}) *Array {
	defer m.lock.Unlock()
	m.lock.Lock()

	m.data = append(m.data, item)

	if len(others) > 0 {
		m.data = append(m.data, others...)
	}

	return m
}

func (m *Array) UnShift(item interface{}, others ...interface{}) *Array {
	defer m.lock.Unlock()
	m.lock.Lock()

	total := 1 + len(others) + len(m.data)

	newSlice := make([]interface{}, total)
	newSlice = newSlice[:1]

	newSlice[0] = item
	newSlice = append(newSlice, others...)
	m.data = append(newSlice, m.data...)

	return m
}

func (m *Array) Pop() (interface{}, bool) {
	defer m.lock.Unlock()
	m.lock.Lock()

	data := m.data
	l := len(data)
	if l > 0 {
		item := m.data[l-1]
		m.data = m.data[:l-1]
		return item, true
	} else {
		return nil, false
	}
}

func (m *Array) Shift() (interface{}, bool) {
	defer m.lock.Unlock()
	m.lock.Lock()

	data := m.data
	l := len(data)
	if l > 0 {
		item := m.data[0]
		m.data = m.data[1:]
		return item, true
	} else {
		return nil, false
	}
}

func (m *Array) Splice(from int, count int) ([]interface{}, bool) {
	defer m.lock.Unlock()
	m.lock.Lock()

	data := m.data
	l := len(data)

	if from+count < l {
		items := data[from : from+count]
		m.data = append(data[:from], data[from+count:])
		return items, true
	} else {
		return nil, false
	}
}

func (m *Array) ForEach(fn func(item interface{}, index int, array *Array) bool) {
	for index, item := range m.data {
		if fn(item, index, m) {
			break
		}
	}
}
