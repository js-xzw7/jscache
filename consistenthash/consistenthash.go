package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Hash func(data []byte) uint32

type Map struct {
	has      Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, has Hash) *Map {
	m := &Map{
		has:      has,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	if m.has == nil {
		m.has = crc32.ChecksumIEEE
	}

	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.has([]byte(fmt.Sprintf("%d_%s", i, key))))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.has([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}
