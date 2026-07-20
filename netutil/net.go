package netutil

import (
	"net"
	"net/http"
	"regexp"
)

func GetInternalIp() string { _ = "STUB: not implemented"; return "" }

func GetRequestPublicIp(req *http.Request) string { _ = "STUB: not implemented"; return "" }

func GetPublicIpInfo() (*PublicIpInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func GetIps() []string { _ = "STUB: not implemented"; return nil }

func GetMacAddrs() []string { _ = "STUB: not implemented"; return nil }

type PublicIpInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Ip          string  `json:"query"`
}

func IsPublicIP(IP net.IP) bool { _ = "STUB: not implemented"; return false }

func IsInternalIP(IP net.IP) bool { _ = "STUB: not implemented"; return false }

func EncodeUrl(urlStr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func UploadFile(filepath string, server string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DownloadFile(filepath string, url string) error { _ = "STUB: not implemented"; return nil }

func IsPingConnected(host string) bool { _ = "STUB: not implemented"; return false }

func IsTelnetConnected(host string, port string) bool { _ = "STUB: not implemented"; return false }

func BuildUrl(scheme, host, path string, query map[string][]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var supportedSchemes = map[string]bool{
	"http":   true,
	"https":  true,
	"ftp":    true,
	"file":   true,
	"mailto": true,
	"ws":     true,
	"wss":    true,
	"data":   true,
}

func validateScheme(scheme string) error { _ = "STUB: not implemented"; return nil }

var hostRegex = regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9]?)(\.[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9]?)+$`)
var pathRegex = regexp.MustCompile(`^\/([a-zA-Z0-9%_-]+(?:\/[a-zA-Z0-9%_-]+)*)$`)

var alphaNumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func AddQueryParams(urlStr string, params map[string][]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
