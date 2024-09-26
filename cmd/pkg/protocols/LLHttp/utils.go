package protocols

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)


func (resp *Response) WithHeader(key, value string) *Response {
	resp.Headers = append(resp.Headers, Header{AsTitle(key), value})
	return resp
}

func (r *Request) WithHeader(key, value string) *Request {
	r.Headers = append(r.Headers, Header{AsTitle(key), value})
	return r
}

func AsTitle(key string) string {
    /* design note --- an empty string could be considered 'in title case', 
    but in practice it's probably programmer error. rather than guess, we'll panic.
    */
    if key == "" {
        panic("empty header key")
    }
    if isTitleCase(key) {
        return key
    }
    /* ---- design note: allocation is very expensive, while iteration through strings is very cheap.
    in general, better to check twice rather than allocate once. ----
    */
    return newTitleCase(key)
}

func newTitleCase(key string) string {
    var b strings.Builder
    b.Grow(len(key))
    for i := range key {

        if i == 0 || key[i-1] == '-' {
            b.WriteByte(upper(key[i]))
        } else {
            b.WriteByte(lower(key[i]))
        }
    }
    return b.String()
}


// straight from K&R C, 2nd edition, page 43. some classics never go out of style.
func lower(c byte) byte {
    /* if you're having trouble understanding this:
        the idea is as follows: A..=Z are 65..=90, and a..=z are 97..=122.
        so upper-case letters are 32 less than their lower-case counterparts (or 'a'-'A' == 32).
        rather than using the 'magic' number 32, we use 'a'-'A' to get the same result.
    */
    if c >= 'A' && c <= 'Z' {
        return c + 'a' - 'A'
    }
    return c
}
func upper(c byte) byte {
    if c >= 'a' && c <= 'z' {
        return c + 'A' - 'a'
    }
    return c
}



// isTitleCase returns true if the given header key is already title case; i.e, it is of the form "Content-Type" or "Content-Length", "Some-Odd-Header", etc.
func isTitleCase(key string) bool {
    // check if this is already title case.
    for i := range key {
        if i == 0 || key[i-1] == '-' {
            if key[i] >= 'a' && key[i] <= 'z' {
                return false
            }
        } else if key[i] >= 'A' && key[i] <= 'Z' {
            return false
        }
    }
    return true
}


func (r *Request) WriteTo(w io.Writer) (n int64, err error) {
	printf := func(format string, args ...any) error {
		m, err := fmt.Fprintf(w, format, args...)
		n += int64(m)
		return err
	}

	if err := printf("%s %s HTTP/1.1\r\n", r.Method, r.Path); err != nil {
		return n, err
	}

	for _, h := range r.Headers {
		if err := printf("%s: %s\r\n", h.Key, h.Value); err != nil {
			return n, err
		}
	}
	printf("\r\n")                 // write the empty line that separates the headers from the body
	err = printf("%s\r\n", r.Body) // write the body and terminate with a newline
	return n, err
}

func (resp *Response) WriteTo(w io.Writer) (n int64, err error) {
	printf := func(format string, args ...any) error {
		m, err := fmt.Fprintf(w, format, args...)
		n += int64(m)
		return err
	}
	if err := printf("HTTP/1.1 %d %s\r\n", resp.StatusCode, http.StatusText(resp.StatusCode)); err != nil {
		return n, err
	}
	for _, h := range resp.Headers {
		if err := printf("%s: %s\r\n", h.Key, h.Value); err != nil {
			return n, err
		}

	}
	if err := printf("\r\n%s\r\n", resp.Body); err != nil {
		return n, err
	}
	return n, nil
}


// ParseResponse parses the given HTTP/1.1 response string into the Response. It returns an error if the Response is invalid,
// - not a valid integer
// - invalid status code
// - missing status text
// - invalid headers
// it doesn't properly handle multi-line headers, headers with multiple values, or html-encoding, etc.zzs
func ParseResponse(raw string) (resp *Response, err error) {
    // response has three parts:
    // 1. Response line
    // 2. Headers
    // 3. Body (optional)
    lines := splitRows(raw)
    info.Info(lines)

    // First line is special.
    first := strings.SplitN(lines[0], " ", 3)
    if !strings.Contains(first[0], "HTTP") {
        return nil, fmt.Errorf("malformed response: first line should contain HTTP version")
    }
    resp = new(Response)
    resp.StatusCode, err = strconv.Atoi(first[1])
    if err != nil {
        return nil, fmt.Errorf("malformed response: expected status code to be an integer, got %q", first[1])
    }
    if first[2] == "" || http.StatusText(resp.StatusCode) != first[2] {
        info.Info("missing or incorrect status text for status code %d: expected %q, but got %q", resp.StatusCode, http.StatusText(resp.StatusCode), first[2])
    }
    var bodyStart int
    // then we have headers, up until the an empty line.
    for i := 1; i < len(lines); i++ {
        info.Info(i, lines[i])
        if lines[i] == "" { // empty line
            bodyStart = i + 1
            break
        }
        key, val, ok := strings.Cut(lines[i], ": ")
        if !ok {
            return nil, fmt.Errorf("malformed response: header %q should be of form 'key: value'", lines[i])
        }
        key = AsTitle(key)
        resp.Headers = append(resp.Headers, Header{key, val})
    }
    resp.Body = strings.TrimSpace(strings.Join(lines[bodyStart:], "\r\n")) // recombine the body using normal newlines.
    return resp, nil
}


func ParseRequest(row string) (r *Response, err error) {
	lines := splitRows(row)
    info.Info(lines)


    first := strings.SplitN(lines[0], " ", 3)

    if !strings.Contains(first[0], "HTTP") {
        errMsg.Error("malformed response: first line should contain HTTP version")
        return 
    }

    r  =  new(Response)

    r.StatusCode, err = strconv.Atoi(first[1])
    if err != nil {
        errMsg.Error(err)
    }

    if first[2] == "" || http.StatusText(r.StatusCode) != first[2] {
        info.Info("missing or incorrect status text for status code")
    }

    var body int 

    for i := 1; i < len(lines); i++ {
        info.Info(i, lines[i])
        if lines[i] == "" { // empty line
            body = i + 1
            break
        }
        key, val, ok := strings.Cut(lines[i], ": ")
        if !ok {
            return nil, fmt.Errorf("malformed response: header %q should be of form 'key: value'", lines[i])
        }
        key = AsTitle(key)
        r.Headers = append(r.Headers, Header{key, val})
    }
    r.Body = strings.TrimSpace(strings.Join(lines[body:], "\r\n")) // recombine the body using normal newlines.
    return r, nil

}

func splitRows(s string) []string {
	if s == "" {
		return nil
	}

	var lines []string
	i := 0
	for {
		j := strings.Index(s[i:], "\r\n")
		if j == -1 {
			lines := append(lines, s[i:])
			return lines
		}

		lines = append(lines, s[i:i+j])

		i += j + 2
	}
}
