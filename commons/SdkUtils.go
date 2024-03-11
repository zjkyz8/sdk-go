package commons

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strconv"
	"time"
)

func GetTimeStamp() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	str := strconv.FormatInt(timestamp, 10)
	return str
}

func GetUUID() string {
	// 生成UUID
	uuidObj := uuid.New()
	// 将UUID转换成字符串
	uuidStr := uuidObj.String()
	// 打印UUID字符串
	return uuidStr
}

func GetMD5(str string) string {
	md5 := md5.New()
	md5.Write([]byte(str))
	result := fmt.Sprintf("%x", md5.Sum(nil))
	return result
}

func IsValidJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// isPropertyExists 判断是否包含指定属性
func IsPropertyExists(obj interface{}, propertyName string) bool {
	if obj == nil || obj == "" {
		return false
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == propertyName {
			return true
		}
	}
	return false
}

func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Interface:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}
