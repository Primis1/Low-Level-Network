package test

import (
	protocols "low-level-tools/cmd/pkg/protocols/LLHttp"
	"reflect"
	"testing"
)

type Header struct{ Key, Value string }

type Request struct {
	Method, Path, Body string
	Headers            []Header
}

type Response struct {
	StatusCode int
	Headers    []Header
	Body       string
}

func TestHTTPResponse(t *testing.T) {
	for name, tt := range map[string]struct {
		input string
		want  *Response
	}{
		"200 OK (no body)": {
			input: "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n",
			want: &Response{
				StatusCode: 200,
				Headers: []Header{
					{"Content-Length", "0"},
				},
			},
		},
		"404 Not Found (w/ body)": {
			input: "HTTP/1.1 404 Not Found\r\nContent-Length: 11\r\n\r\nHello World\r\n",
			want: &Response{
				StatusCode: 404,
				Headers: []Header{
					{"Content-Length", "11"},
				},
				Body: "Hello World",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := protocols.ParseResponse(tt.input)
			if err != nil {
				t.Errorf("ParseResponse(%q) returned error: %v", tt.input, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseResponse(%q) = %#+v, want %#+v", tt.input, got, tt.want)
			}

			gotStr := string(got.Body)
			got2, err := protocols.ParseResponse(gotStr)
			if err != nil {
				t.Errorf("ParseResponse(%q) returned error: %v", gotStr, err)
			}
			if !reflect.DeepEqual(got2, got) {
				t.Errorf("ParseResponse(%q) = %#+v, want %#+v", gotStr, got2, got)
			}
		})
	}
}

func TestHTTPRequest(t *testing.T) {
	for name, tt := range map[string]struct {
		input string
		want  Request
	}{
		"GET (no body)": {
			input: "GET / HTTP/1.1\r\nHost: www.example.com\r\n\r\n",
			want: Request{
				Method: "GET",
				Path:   "/",
				Headers: []Header{
					{"Host", "www.example.com"},
				},
			},
		},
		"POST (w/ body)": {
			input: "POST / HTTP/1.1\r\nHost: www.example.com\r\nContent-Length: 11\r\n\r\nHello World\r\n",
			want: Request{
				Method: "POST",
				Path:   "/",
				Headers: []Header{
					{"Host", "www.example.com"},
					{"Content-Length", "11"},
				},
				Body: "Hello World",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := protocols.ParseRequest(tt.input)
			if err != nil {
				t.Errorf("ParseRequest(%q) returned error: %v", tt.input, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRequest(%q) = %#+v, want %#+v", tt.input, got, tt.want)
			}

			gotStr := string(got.Body)
			got2, err := protocols.ParseRequest(gotStr)
			if err != nil {
				t.Errorf("ParseRequest(%q) returned error: %v", gotStr, err)
			}
			if !reflect.DeepEqual(got, got2) {
				t.Errorf("ParseRequest(%q) = %+v, want %+v", gotStr, got2, got)
			}
		})
	}
}
