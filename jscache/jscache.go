package jscache

import (
	"errors"
	"fmt"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}

type group struct {
	name      string
	getter    Getter
	maincache cache
}

var (
	mu     sync.Mutex
	groups = make(map[string]*group)
)

func NewGroup(name string, getter Getter, cacheBytes int64) *group {
	if getter == nil {
		panic("nil Getter")
	}

	mu.Lock()
	defer mu.Unlock()

	g := &group{
		name:      name,
		getter:    getter,
		maincache: cache{cacheBytes: cacheBytes},
	}

	groups[name] = g
	return g
}

func GetGroup(key string) *group {
	mu.Lock()
	defer mu.Unlock()

	if v, ok := groups[key]; ok {
		return v
	}
	return nil
}

func (g *group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, errors.New("key is require")
	}

	if v, ok := g.maincache.Get(key); ok {
		fmt.Printf("cache get [%s]\n", key)
		return v, nil
	}

	return g.load(key)
}

func (g *group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

func (g *group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	value := ByteView{b: copyBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *group) populateCache(key string, value ByteView) {
	g.maincache.Add(key, value)
}
