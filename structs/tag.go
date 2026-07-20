package structs

type Tag struct {
	Name    string
	Options []string
}

func newTag(tag string) *Tag { _ = "STUB: not implemented"; return nil }

func (t *Tag) HasOption(opt string) bool { _ = "STUB: not implemented"; return false }

func (t *Tag) IsEmpty() bool { _ = "STUB: not implemented"; return false }
