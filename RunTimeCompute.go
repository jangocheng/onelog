package onelog

import (
	"runtime"
	"strconv"
	"time"
)

//RunTimeCompute 运行时计算的接口，实现此接口的可放入运行时值
type RunTimeCompute interface {
	//GetName 在写入日志时需要的名称
	GetName() string
	//Values 在运行时计算得到的值
	Values() []byte
}

//RunTimeCompute 运行时计算的接口，实现此接口的可放入运行时值
type RunTimeComputes struct {
	curr RunTimeCompute
	next *RunTimeComputes
}

type CoroutineID struct {
}

func (cid *CoroutineID) GetName() string {
	return CoroutineIDName
}

func (cid *CoroutineID) Values() []byte {
	result := make([]byte, 0)

	result = strconv.AppendInt(result, runtime.Goid(), 10)
	return result
}

//TimeValue 得到当前时间的值
type TimeValue struct {
}

func (t *TimeValue) GetName() string {
	return TimeName
}

func (t *TimeValue) Values() []byte {
	buf := make([]byte, 0)
	var now = time.Now()

	if TimeFormat == "" {
		return strconv.AppendInt(buf, now.Unix(), 10)
	}

	return now.AppendFormat(buf, TimeFormat)
}

//Caller 得到当前的调用者信息，可根据跳过值增加
type Caller struct {
}

func (*Caller) GetName() string {
	return CallerName
}

func (*Caller) Values() []byte {
	_, file, line, ok := runtime.Caller(CallerSkipFrameCount + 2)
	var buf = make([]byte, len(file)+7)

	if ok {
		buf = append(buf[:0], file...)
		buf = append(buf, ' ')
		buf = strconv.AppendInt(buf, int64(line), 10)
	} else {
		buf = append(buf[:0], "no found."...)
	}

	return buf
}