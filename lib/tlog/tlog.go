/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-05 15:39:17
 */

package tlog

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var begin time.Time
var tmp time.Time
var file *os.File

func Start() {
	begin = time.Now()
	tmp = begin

	filePath := "/work/goproject/IM/logs/tlog.log"
	var err error
	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	_, fileName, line, _ := runtime.Caller(1)
	fmt.Fprintf(file, "\ntime:%s  起始点===%s:%d\n", tmp.Format("2006-01-02 15:04:05.000"), fileName, line)
}

func Close() {
	d := time.Since(begin).Seconds()
	_, fileName, line, _ := runtime.Caller(1)
	fmt.Fprintf(file, "time:%s  结束点===%s:%d耗时%.2f\n", tmp.Format("2006-01-02 15:04:05.000"), fileName, line, d)
	file.Close()
}

func Info() {
	d := time.Since(tmp).Seconds()
	tmp = time.Now()
	_, fileName, line, _ := runtime.Caller(1)
	fmt.Fprintf(file, "time:%s  %s:%d耗时%.2f\n", tmp.Format("2006-01-02 15:04:05.000"), fileName, line, d)
}

func Infof(format string, args ...interface{}) {
	_, fileName, line, _ := runtime.Caller(1)
	fmt.Fprintf(file, "time:%s  %s:%d 打印信息：\n", tmp.Format("2006-01-02 15:04:05.000"), fileName, line)
	fmt.Fprintf(file, format, args...)
}

func CountTime(t time.Time) {
	d := time.Since(t).Seconds()
	_, fileName, line, _ := runtime.Caller(1)
	fmt.Fprintf(file, "%s:%d耗时%.2f\n", fileName, line, d)
}
