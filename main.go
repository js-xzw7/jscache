package main

import (
	"fmt"
	"jscache/jscache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	jscache.NewGroup("js", jscache.GetterFunc(func(key string) ([]byte, error) {
		fmt.Printf("db query key [%s]\n", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}

		return nil, fmt.Errorf("%s not exist", key)
	}), 2<<10)

	addr := "127.0.0.1:8089"
	peers := jscache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
