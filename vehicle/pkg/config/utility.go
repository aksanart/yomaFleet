package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Getter interface {
	Get() interface{}
	GetString() string
	GetBool() bool
	GetFloat() float64
	GetInt() int64
	GetUint() uint64
	GetDuration() time.Duration
	GetList() []string
}

func (v data) Get() interface{} {
	return v.value
}

func (v data) GetString() string {
	return fmt.Sprintf("%s", v.value)
}

func (v data) GetBool() bool {
	if val, err := strconv.ParseBool(fmt.Sprintf("%s", v.value)); err == nil {
		return val
	}
	return false
}

func (v data) GetFloat() float64 {
	if val, err := strconv.ParseFloat(fmt.Sprintf("%s", v.value), 64); err == nil {
		return val
	}
	return 0.0
}

func (v data) GetInt() int64 {
	if val, err := strconv.ParseInt(fmt.Sprintf("%s", v.value), 10, 64); err == nil {
		return val
	}
	return 0
}

func (v data) GetUint() uint64 {
	if val, err := strconv.ParseUint(fmt.Sprintf("%s", v.value), 10, 64); err == nil {
		return val
	}
	return 0
}

func (v data) GetList() []string {
	slice := strings.Split(v.value.(string), ",")
	for i, val := range slice {
		slice[i] = strings.TrimSpace(val)
	}
	return slice
}

func (v data) GetDuration() time.Duration {
	if parsed, err := time.ParseDuration(fmt.Sprintf("%s", v.value)); err == nil {
		return parsed
	}
	return 0
}

// GetMapOfBoolean returns config value for `key` as map of boolean. If no value is
// found, returns a map with `defaultKey` as the key, `defaultVal` as the value
func (v data) GetMapOfBoolean(value bool) map[string]bool {
	_map := map[string]bool{}

	arr := v.GetList()
	if len(arr) == 0 {
		return _map
	}

	for _, key := range arr {
		_map[key] = value
	}

	return _map
}

func Newdata(val interface{}) data {
	return data{value: val}
}

func (v *data) Setdata(val interface{}) {
	v.value = val
}
