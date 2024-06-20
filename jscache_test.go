package jscache

import (
	"errors"
	"fmt"
	"testing"
)

// func TestGetter(t *testing.T) {
// 	var f Getter = GetterFunc(func(key string) ([]byte, error) { return []byte(key), nil })

// 	expect := []byte("hello")
// 	if v, _ := f.Get("hello"); !reflect.DeepEqual(v, expect) {
// 		t.Errorf("callback failed")
// 	}
// }

func TestGroupGet(t *testing.T) {
	var db = map[string]string{
		"willian": "威廉",
		"tony":    "托尼",
		"lili":    "莉莉",
	}

	flag := make(map[string]int, len(db))

	g := NewGroup("test", GetterFunc(func(key string) ([]byte, error) {
		if v, ok := db[key]; ok {
			if _, ok := flag[key]; !ok {
				flag[key] = 1
			}
			fmt.Printf("flag[%s] = %d\n", key, flag[key])
			return []byte(v), nil
		}

		return nil, errors.New("locally is not found")
	}), 2<<10)

	for i := 0; i < 2; i++ {
		for k, v := range db {
			if bv, err := g.Get(k); err != nil {
				t.Errorf("g.Get(%s) error:%v", k, err)
			} else {
				if v != bv.String() {
					t.Errorf("bv.String(%s) != v(%s)", bv.String(), v)
				}
			}
		}
	}

}
