package requestfiles

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

type dummyResponseWriter struct {
	saveb []byte
	h     http.Header
}

func (self *dummyResponseWriter) Write(b []byte) (int, error) {
	self.saveb = b
	return 0, nil
}

func (self *dummyResponseWriter) Header() http.Header {
	if self.h == nil {
		self.h = make(http.Header)
	}
	return self.h
}

func (self *dummyResponseWriter) WriteHeader(int) {
}

func TestGetHandle(t *testing.T) {
	f := GetHandler("./testdata/", 6)
	if f == nil {
		t.Errorf("nil Handle")
	}

	req := new(http.Request)
	res := new(dummyResponseWriter)

	req.URL = new(url.URL)
	req.URL.Path = "/home/sample.html"
	f(res, req)

	if res.Header().Get(ContentType) != "text/html" {
		t.Errorf("Content Type is not text/html:%v", res.Header().Get(ContentType))
	}
}

func TestNonHandle(t *testing.T) {
	f := GetHandler("./testdata/", 6)
	req := new(http.Request)
	res := new(dummyResponseWriter)

	req.URL = new(url.URL)
	req.URL.Path = "/home/detarame.html"

	f(res, req)

	if res.Header().Get(ContentType) != "text/plain; charset=utf-8" {
		t.Errorf("NonHandle:%v", res.Header().Get(ContentType))
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("Hello\n")
	}
}
