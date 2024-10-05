/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 13:13:19
 */

package lock

import "sync"

type emptyLock struct{}

func NewEmptyLock() sync.Locker {
	return &emptyLock{}
}

func (e *emptyLock) Lock() {}

func (e *emptyLock) Unlock() {}
