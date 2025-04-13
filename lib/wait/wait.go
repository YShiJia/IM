/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-04 17:11:05
 */

package wait

import (
	"time"
)

type Waiter interface {
	//重置
	Reset()
	// 等待一会
	Wait()
}

var _ Waiter = (*waiterByExponentialBackoff)(nil)

type waiterByExponentialBackoff struct {
	backoff    int
	maxBackoff int
}

var defaultWaiterExponential = 16
var defaultWaiterUnitTime = 10 * time.Millisecond

type waiterByExponentialBackoffOption func(*waiterByExponentialBackoff)

func NewWaiterByExponentialBackoff(pts ...waiterByExponentialBackoffOption) *waiterByExponentialBackoff {
	w := waiterByExponentialBackoff{
		backoff:    1,
		maxBackoff: defaultWaiterExponential,
	}
	for _, opt := range pts {
		opt(&w)
	}
	w.repair()

	return &w
}

func (w *waiterByExponentialBackoff) Reset() {
	w.backoff = 1
}

func (w *waiterByExponentialBackoff) Wait() {
	time.Sleep(time.Duration(w.backoff) * defaultWaiterUnitTime)
	// 翻倍递增，到顶为止
	w.backoff = min(w.maxBackoff, w.backoff<<1)

}

func (w *waiterByExponentialBackoff) repair() {
	if w.maxBackoff <= 0 {
		w.maxBackoff = defaultWaiterExponential / 2
	}
	if w.maxBackoff > 16 {
		w.maxBackoff = defaultWaiterExponential
	}
}

func WithMaxBackoff(maxBackoff int) waiterByExponentialBackoffOption {
	return func(backoff *waiterByExponentialBackoff) {
		backoff.maxBackoff = maxBackoff
	}
}
