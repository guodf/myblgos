package util

import "time"

// DateTimeUtil 时间包装类
type DateTimeUtil struct {
	time.Time
}

// Today 今天
func (t DateTimeUtil) Today() DateTimeUtil {
	year, month, day := t.Date()
	time := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return DateTimeUtil{time}
}

// Seconds 获取总描述
func (t DateTimeUtil) Seconds() int64 {
	return t.Unix()
}

// AddDays 添加指定的天数
func (t DateTimeUtil) AddDays(days int) DateTimeUtil {
	time := t.AddDate(0, 0, days)
	return DateTimeUtil{time}
}
