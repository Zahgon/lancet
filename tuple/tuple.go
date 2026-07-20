package tuple

type Tuple2[A any, B any] struct {
	FieldA A
	FieldB B
}

func (t Tuple2[A, B]) Unbox() (A, B) { _ = "STUB: not implemented"; return *new(A), *new(B) }

func NewTuple2[A any, B any](a A, b B) Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

func Zip2[A any, B any](a []A, b []B) []Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

func Unzip2[A any, B any](tuples []Tuple2[A, B]) ([]A, []B) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Tuple3[A any, B any, C any] struct {
	FieldA A
	FieldB B
	FieldC C
}

func (t Tuple3[A, B, C]) Unbox() (A, B, C) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C)
}

func NewTuple3[A any, B any, C any](a A, b B, c C) Tuple3[A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

func Zip3[A any, B any, C any](a []A, b []B, c []C) []Tuple3[A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip3[A any, B any, C any](tuples []Tuple3[A, B, C]) ([]A, []B, []C) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Tuple4[A any, B any, C any, D any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
}

func (t Tuple4[A, B, C, D]) Unbox() (A, B, C, D) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D)
}

func NewTuple4[A any, B any, C any, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

func Zip4[A any, B any, C any, D any](a []A, b []B, c []C, d []D) []Tuple4[A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip4[A any, B any, C any, D any](tuples []Tuple4[A, B, C, D]) ([]A, []B, []C, []D) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

type Tuple5[A any, B any, C any, D any, E any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
}

func (t Tuple5[A, B, C, D, E]) Unbox() (A, B, C, D, E) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E)
}

func NewTuple5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

func Zip5[A any, B any, C any, D any, E any](a []A, b []B, c []C, d []D, e []E) []Tuple5[A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip5[A any, B any, C any, D any, E any](tuples []Tuple5[A, B, C, D, E]) ([]A, []B, []C, []D, []E) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

type Tuple6[A any, B any, C any, D any, E any, F any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
}

func (t Tuple6[A, B, C, D, E, F]) Unbox() (A, B, C, D, E, F) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F)
}

func NewTuple6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	_ = "STUB: not implemented"
	return nil
}

func Zip6[A any, B any, C any, D any, E any, F any](a []A, b []B, c []C, d []D, e []E, f []F) []Tuple6[A, B, C, D, E, F] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip6[A any, B any, C any, D any, E any, F any](tuples []Tuple6[A, B, C, D, E, F]) ([]A, []B, []C, []D, []E, []F) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil
}

type Tuple7[A any, B any, C any, D any, E any, F any, G any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
}

func (t Tuple7[A, B, C, D, E, F, G]) Unbox() (A, B, C, D, E, F, G) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G)
}

func NewTuple7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	_ = "STUB: not implemented"
	return nil
}

func Zip7[A any, B any, C any, D any, E any, F any, G any](a []A, b []B, c []C, d []D, e []E, f []F, g []G) []Tuple7[A, B, C, D, E, F, G] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip7[A any, B any, C any, D any, E any, F any, G any](tuples []Tuple7[A, B, C, D, E, F, G]) ([]A, []B, []C, []D, []E, []F, []G) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil
}

type Tuple8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
	FieldH H
}

func (t Tuple8[A, B, C, D, E, F, G, H]) Unbox() (A, B, C, D, E, F, G, H) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H)
}

func NewTuple8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	_ = "STUB: not implemented"
	return nil
}

func Zip8[A any, B any, C any, D any, E any, F any, G any, H any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H) []Tuple8[A, B, C, D, E, F, G, H] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip8[A any, B any, C any, D any, E any, F any, G any, H any](tuples []Tuple8[A, B, C, D, E, F, G, H]) ([]A, []B, []C, []D, []E, []F, []G, []H) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil
}

type Tuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
	FieldH H
	FieldI I
}

func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unbox() (A, B, C, D, E, F, G, H, I) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), *new(I)
}

func NewTuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	_ = "STUB: not implemented"
	return nil
}

func Zip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I) []Tuple9[A, B, C, D, E, F, G, H, I] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](tuples []Tuple9[A, B, C, D, E, F, G, H, I]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil, nil
}

type Tuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
	FieldH H
	FieldI I
	FieldJ J
}

func (t Tuple10[A, B, C, D, E, F, G, H, I, J]) Unbox() (A, B, C, D, E, F, G, H, I, J) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), *new(I), *new(J)
}

func NewTuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Tuple10[A, B, C, D, E, F, G, H, I, J] {
	_ = "STUB: not implemented"
	return nil
}

func Zip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, j []J) []Tuple10[A, B, C, D, E, F, G, H, I, J] {
	_ = "STUB: not implemented"
	return nil
}

func Unzip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](tuples []Tuple10[A, B, C, D, E, F, G, H, I, J]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I, []J) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
}

func getByIndex[T any](slice []T, index int) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}
