package constraints

type Comparator interface {
	Compare(v1, v2 any) int
}
