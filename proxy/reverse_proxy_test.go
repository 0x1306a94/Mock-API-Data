package proxy

import (
	"Mock-API-Data/constant"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
)

type _MockServer struct {
	proxy *ReverseProxy
}

func (m *_MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mockHost := r.Header.Get(constant.HTTPHeaderMockProjectKey)
	if mockHost == "" {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(constant.NotHeaderMockProjectKeyErrorReason))
		return
	}
	if m.proxy == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	m.proxy.ServeHTTP(w, r)
}

func TestReverseProxy(t *testing.T) {

	proxyURL := os.Getenv("PROXY_URL")
	targetURL, err := url.Parse(proxyURL)
	if err != nil {
		t.Fatal(err)
		return
	}
	// TODO: 1.多应用匹配 2.模拟数据LRU缓存 3.查询配置的模拟数据 如有匹配则直接返回 ...
	proxy := NewSingleHostReverseProxy(targetURL)
	m := &_MockServer{
		proxy: proxy,
	}
	if err := http.ListenAndServe(":9999", m); err != nil {
		fmt.Println(err)
	}
}
