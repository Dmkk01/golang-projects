package cache

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type mockWriter response

func newMockWriter() *mockWriter {
	return &mockWriter{
		body:   []byte{},
		header: http.Header{},
	}
}

func (mw *mockWriter) Write(b []byte) (int, error) {
	mw.body = make([]byte, len(b))
	for k, v := range b {
		mw.body[k] = v
	}
	return len(b), nil
}

func (mw *mockWriter) WriteHeader(code int) { mw.code = code }

func (mw *mockWriter) Header() http.Header { return mw.header }

func TestWriter(t *testing.T) {
	mw := newMockWriter()

	res := "/test/url?with=query"

	u, err := url.Parse(res)

	if err != nil {
		t.Fatal("Invalid url")
	}

	req := &http.Request{
		URL: u,
	}

	t.Log("Testing NewWriter")

	w := NewWriter(mw, req)

	if w.resource != res {
		t.Error("Invalid resource")
	}

	if w.writer != mw {
		t.Fatal("Invalid writer")
	}

	t.Log("Test Header")
	h := w.Header()
	h.Add("test", "test")

	h2 := w.response.header

	if h2.Get("test") != "test" {
		t.Error("Invalid header")
	}

	t.Log("Test WriteHeader")

	c := 201
	w.WriteHeader(c)

	if w.response.code != c {
		t.Error("Invalid code")
	}

	if mw.code != c {
		t.Error("Invalid code")
	}

	h2 = mw.header

	if h2.Get("test") != "test" {
		t.Error("Header not written")
	}

	t.Log("Test Write")
	bd := []byte{1, 2, 3, 4, 5}
	n, err := w.Write(bd)

	if err != nil {
		t.Error("Error writing")
	}

	if n != len(bd) {
		t.Error("Invalid length")
	}

	if &w.response.body == &bd {
		t.Error("Invalid body")
	}

	if !reflect.DeepEqual(w.response.body, bd) {
		t.Error("Invalid body")
	}

	if &mw.body == &bd {
		t.Error("Invalid body")
	}

	if !reflect.DeepEqual(mw.body, bd) {
		t.Error("Invalid body")
	}

}
