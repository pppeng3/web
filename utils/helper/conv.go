package helper

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func MustString(val interface{}) string {
	switch v := val.(type) {
	case float32:
		return fmt.Sprintf("%f", val)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", val)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case string:
		return val.(string)
	case time.Time:
		return v.Format("2006-01-02 15:04:05.000")
	case decimal.Decimal:
		return v.StringFixed(16)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func ParseShortFloat(f float64, n int) string {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	v := strings.TrimRight(floatStr, "0")
	if v[len(v)-1] == '.' {
		v = v[:len(v)-1]
	}
	return v
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		} else if fraction <= -0.5 {
			whole--
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func MustInt8(val interface{}) int8 {
	switch val.(type) {
	case float32, float64:
		return int8(reflect.ValueOf(val).Float())
	case int, int8, int16, int32, int64:
		return int8(reflect.ValueOf(val).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int8(reflect.ValueOf(val).Uint())
	case string:
		v, _ := strconv.ParseInt(reflect.ValueOf(val).String(), 0, 0)
		return int8(v)
	default:
		panic(fmt.Errorf("invalid value type", val))
	}
}

func MustInt64(val interface{}) int64 {
	switch val.(type) {
	case float32, float64:
		return int64(reflect.ValueOf(val).Float())
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(val).Int()
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(val).Uint())
	case string:
		v, _ := strconv.ParseInt(reflect.ValueOf(val).String(), 0, 0)
		return v
	default:
		panic(fmt.Errorf("invalid value type", val))
	}
}

func MustInt(val interface{}) int {
	switch val.(type) {
	case float32, float64:
		return int(reflect.ValueOf(val).Float())
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(val).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(val).Uint())
	case string:
		v, _ := strconv.ParseInt(reflect.ValueOf(val).String(), 0, 0)
		return int(v)
	default:
		panic(fmt.Errorf("invalid value type", val))
	}
}

func MustFloat(val interface{}) float64 {
	switch vv := val.(type) {
	case float32, float64:
		return reflect.ValueOf(val).Float()
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(val).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(val).Uint())
	case string:
		v, _ := strconv.ParseFloat(reflect.ValueOf(val).String(), 64)
		return v
	case decimal.Decimal:
		v, _ := vv.Float64()
		return v
	default:
		panic(fmt.Errorf("invalid value type", val))
	}
}

func SliceToStrings(list []interface{}) (ret []string) {
	for _, val := range list {
		ret = append(ret, val.(string))
	}
	return ret
}

//use tag to define encode name, or use name origin
func Struct2Map(obj interface{}, fields ...[]string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//如果是指针变量的话，就取指针指向的值
	if t.Kind() == reflect.Ptr {
		t, v = t.Elem(), v.Elem()
	}
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		filedName := t.Field(i).Name
		if tag := t.Field(i).Tag.Get("s2m"); tag != "" {
			filedName = tag
		}
		data[filedName] = v.Field(i).Interface()
	}
	if len(fields) == 0 || len(fields[0]) == 0 || fields[0][0] == "" {
		return data
	}
	var result = make(map[string]interface{}, len(fields[0]))
	for _, field := range fields[0] {
		if field == "*" {
			return data
		}
		result[field] = data[field]
	}
	return result
}
