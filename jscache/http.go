package jscache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_jscache/"

type HTTPPool struct {
	self     string
	basePath string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (h *HTTPPool) Log(form string, v ...interface{}) {
	log.Printf("[Server %s] %s\n", h.self, fmt.Sprintf(form, v...))
}

func (h *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.Write([]byte("ok"))
		return
	}
	if !strings.HasPrefix(r.URL.Path, h.basePath) {
		panic(fmt.Sprintf("HTTPPool serving unexpected path:" + r.URL.Path))
	}

	h.Log("%s:%s", r.Method, r.URL.Path)

	parts := strings.SplitN(r.URL.Path[len(h.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, fmt.Sprintf("no such group:%s", groupName), http.StatusNotFound)
		return
	}

	value, err := group.Get(key)
	if err != nil {
		http.Error(w, fmt.Sprintf("get value error:%s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/octet-stream")
	w.Write(value.ByteSlice())
}
