package netutil

import (
	"net/http"
)

func doHttpRequest(method, reqUrl string, params ...any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setHeaderAndQueryParam(req *http.Request, reqUrl string, header, queryParam any) error {
	_ = "STUB: not implemented"
	return nil
}

func setHeaderAndQueryAndBody(req *http.Request, reqUrl string, header, queryParam, body any) error {
	_ = "STUB: not implemented"
	return nil
}

func setHeader(req *http.Request, header any) error { _ = "STUB: not implemented"; return nil }

func setUrl(req *http.Request, reqUrl string) error { _ = "STUB: not implemented"; return nil }

func setQueryParam(req *http.Request, reqUrl string, queryParam any) error {
	_ = "STUB: not implemented"
	return nil
}

func setBodyByte(req *http.Request, body any) error { _ = "STUB: not implemented"; return nil }

func getClient(client any) (*http.Client, error) { _ = "STUB: not implemented"; return nil, nil }
