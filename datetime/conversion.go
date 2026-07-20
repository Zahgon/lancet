package datetime

type theTime struct {
	unix int64
}

func NewUnixNow() *theTime { _ = "STUB: not implemented"; return nil }

func NewUnix(unix int64) *theTime { _ = "STUB: not implemented"; return nil }

func NewFormat(t string) (*theTime, error) { _ = "STUB: not implemented"; return nil, nil }

func NewISO8601(iso8601 string) (*theTime, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *theTime) ToUnix() int64 { _ = "STUB: not implemented"; return 0 }

func (t *theTime) ToFormat() string { _ = "STUB: not implemented"; return "" }

func (t *theTime) ToFormatForTpl(tpl string) string { _ = "STUB: not implemented"; return "" }

func (t *theTime) ToIso8601() string { _ = "STUB: not implemented"; return "" }
