package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu sync.Mutex // 排他制御のためのミューテックス?
	records []Record
}


func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock() // 遅延実行される関数の引数は即時評価される。後入先出法で実行される。
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}


// START:types
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

//END:types

var ErrOffsetNotFound = fmt.Errorf("offset not found")


