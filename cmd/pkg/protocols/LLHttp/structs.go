package protocols

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
