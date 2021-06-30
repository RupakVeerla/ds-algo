package hashtable

import "fmt"

type hashtable struct {
	data [][]keyvalue
}

type keyvalue struct {
	key   string
	value interface{}
}

func New(size int) *hashtable {
	h := new(hashtable)
	h.data = make([][]keyvalue, size, size)
	return h
}

func (h *hashtable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(h.data)
	}
	return hash
}

func (h *hashtable) Set(key string, value interface{}) {
	hash := h.hash(key)
	if h.data[hash] != nil {
		for _, v := range h.data[hash] {
			if v.key == key {
				v.value = value
				return
			}
		}
	}
	h.data[hash] = append(h.data[hash], keyvalue{key, value})
}

func (h *hashtable) Get(key string) interface{} {
	hash := h.hash(key)
	if h.data[hash] == nil {
		return nil
	}
	fmt.Println(h.data)
	for _, v := range h.data[hash] {
		if v.key == key {
			return v.value
		}
	}
	return nil
}

func (h *hashtable) Keys() (keys []string) {
	for _, v := range h.data {
		if v != nil {
			for _, val := range v {
				keys = append(keys, val.key)
			}
		}
	}
	return
}
