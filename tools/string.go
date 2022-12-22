package tools

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func StringToInt64(str string) int64 {
	parsedInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return parsedInt
}

func IsOneOf(obj string, list ...string) bool {
	for _, v := range list {
		if obj == v {
			return true
		}
	}
	return false
}

// Shuffle 切片随机排序，传入的参数类型为 []any
func Shuffle(list interface{}) []interface{} {
	// interface{} 代表 []any
	rv := reflect.ValueOf(list)
	if rv.Kind() != reflect.Slice {
		// 不是切片就是来捣乱的
		return nil
	}
	// interface{} 转 []interface{}
	realList := make([]interface{}, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		realList[i] = rv.Index(i).Interface()
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(realList), func(i, j int) {
		realList[i], realList[j] = realList[j], realList[i]
	})
	return realList
}
