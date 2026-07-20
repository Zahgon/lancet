package netutil

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

func HttpGet(url string, params ...any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HttpPost(url string, params ...any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HttpPut(url string, params ...any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HttpDelete(url string, params ...any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HttpPatch(url string, params ...any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseHttpResponse(resp *http.Response, obj any) error { _ = "STUB: not implemented"; return nil }

func ConvertMapToQueryString(param map[string]any) string { _ = "STUB: not implemented"; return "" }

type HttpRequest struct {
	RawURL      string
	Method      string
	Headers     http.Header
	QueryParams url.Values
	FormData    url.Values
	File        *File
	Body        []byte
}

type HttpClientConfig struct {
	Timeout          time.Duration
	SSLEnabled       bool
	TLSConfig        *tls.Config
	Compressed       bool
	HandshakeTimeout time.Duration
	ResponseTimeout  time.Duration
	Verbose          bool
	Proxy            *url.URL
}

var defaultHttpClientConfig = &HttpClientConfig{
	Timeout:          50 * time.Second,
	Compressed:       false,
	HandshakeTimeout: 10 * time.Second,
	ResponseTimeout:  10 * time.Second,
}

type HttpClient struct {
	*http.Client
	TLS     *tls.Config
	Request *http.Request
	Config  HttpClientConfig
	Context context.Context
}

func NewHttpClient() *HttpClient { _ = "STUB: not implemented"; return nil }

func NewHttpClientWithConfig(config *HttpClientConfig) *HttpClient {
	_ = "STUB: not implemented"
	return nil
}

func (client *HttpClient) SendRequest(request *HttpRequest) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *HttpClient) AsyncSendRequest(request *HttpRequest, respChan chan *http.Response, errChan chan error) {
	_ = "STUB: not implemented"
	return
}

func (client *HttpClient) DecodeResponse(resp *http.Response, target any) error {
	_ = "STUB: not implemented"
	return nil
}

func (client *HttpClient) setTLS(rawUrl string) { _ = "STUB: not implemented"; return }

func (client *HttpClient) setHeader(req *http.Request, headers http.Header) {
	_ = "STUB: not implemented"
	return
}

func (client *HttpClient) setQueryParam(req *http.Request, reqUrl string, queryParam url.Values) error {
	_ = "STUB: not implemented"
	return nil
}

func (client *HttpClient) setFormData(req *http.Request, values url.Values, setFile SetFileFunc) error {
	_ = "STUB: not implemented"
	return nil
}

type SetFileFunc func(req *http.Request, values url.Values) error

type File struct {
	Content   []byte
	Path      string
	FieldName string
	FileName  string
}

func setFile(f *File) SetFileFunc { _ = "STUB: not implemented"; return *new(SetFileFunc) }

func validateRequest(req *HttpRequest) error { _ = "STUB: not implemented"; return nil }

func StructToUrlValues(targetStruct any) (url.Values, error) {
	_ = "STUB: not implemented"
	return *new(url.Values), nil
}
