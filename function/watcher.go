package function

import "time"

type Watcher struct {
	startTime int64
	stopTime  int64
	excuting  bool
}

func NewWatcher() *Watcher { _ = "STUB: not implemented"; return nil }

func (w *Watcher) Start() { _ = "STUB: not implemented"; return }

func (w *Watcher) Stop() { _ = "STUB: not implemented"; return }

func (w *Watcher) GetElapsedTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (w *Watcher) Reset() { _ = "STUB: not implemented"; return }
