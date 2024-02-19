package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, v interface{})
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{storage: make(map[string]interface{})}
}

type cacheImpl struct {
	storage map[string]interface{}
}

func (c *cacheImpl) Get(k string) (interface{}, bool) {
	value, ok := c.storage[k]
	return value, ok
}

func (c *cacheImpl) Set(k string, v interface{}) {
	c.storage[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %v", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	c.Set("hello", "world")
	c.Set("test", struct{ field string }{"my map"})
	db := newDbImpl(c)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
