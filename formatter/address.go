package formatter

type AddressInfo struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	IDN      string `json:"idn"`
	Postcode string `json:"postcode"`
	Province string `json:"province"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Street   string `json:"street"`
	Addr     string `json:"addr"`
}

type fuzzyResult struct {
	A1     string
	A2     string
	A3     string
	Street string
}

func ParseCNAddress(str string, withUser bool) *AddressInfo { _ = "STUB: not implemented"; return nil }

func ParsePersonInfo(str string) *AddressInfo { _ = "STUB: not implemented"; return nil }

func fuzz(addr string) *fuzzyResult { _ = "STUB: not implemented"; return nil }

func parse(a1, a2, a3 string) *AddressInfo { _ = "STUB: not implemented"; return nil }
