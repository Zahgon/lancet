package xerror

import (
	"context"
)

type TryCatch struct {
	ctx         context.Context
	tryFunc     func(ctx context.Context) error
	catchFunc   func(ctx context.Context, err error)
	finallyFunc func(ctx context.Context)
}

func NewTryCatch(ctx context.Context) *TryCatch { _ = "STUB: not implemented"; return nil }

func (tc *TryCatch) Try(tryFunc func(ctx context.Context) error) *TryCatch {
	_ = "STUB: not implemented"
	return nil
}

func (tc *TryCatch) Catch(catchFunc func(ctx context.Context, err error)) *TryCatch {
	_ = "STUB: not implemented"
	return nil
}

func (tc *TryCatch) Finally(finallyFunc func(ctx context.Context)) *TryCatch {
	_ = "STUB: not implemented"
	return nil
}

func (tc *TryCatch) Do() { _ = "STUB: not implemented"; return }

type CatchError struct {
	Msg   string
	File  string
	Line  int
	Cause error
}

func (e *CatchError) Error() string { _ = "STUB: not implemented"; return "" }

func WrapCatchError(err error, msg string) *CatchError { _ = "STUB: not implemented"; return nil }
