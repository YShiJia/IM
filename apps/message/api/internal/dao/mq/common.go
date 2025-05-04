/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-04 14:27:27
 */

package mq

import (
	"github.com/segmentio/kafka-go"
	"sync"
)

var recvMsgQueueWriter map[string]*kafka.Writer
var SendMsgQueueReader *kafka.Reader

var rwLock sync.RWMutex

func init() {
	recvMsgQueueWriter = make(map[string]*kafka.Writer)
}

func GetRecvMsgQueueWriter(serviceName string) *kafka.Writer {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if w, exists := recvMsgQueueWriter[serviceName]; exists {
		return w
	}
	return nil
}

func SetRecvMsgQueueWriter(serviceName string, w *kafka.Writer) {
	rwLock.Lock()
	defer rwLock.Unlock()
	recvMsgQueueWriter[serviceName] = w
}

func DelRecvMsgQueueWriter(serviceName string) *kafka.Writer {
	rwLock.Lock()
	defer rwLock.Unlock()
	delete(recvMsgQueueWriter, serviceName)
	return nil
}
