package entities

import (
	"encoding/json"
	"strconv"
)

type M map[string]interface{}

type MValue struct {
	Value interface{}
}

// Bool implements IMValue.
func (mv *MValue) Bool() bool {
	str := mv.String()

	val, _ := strconv.ParseBool(str)

	return val
}

// Float32 implements IMValue.
func (mv *MValue) Float32() float32 {
	str := mv.String()

	val, _ := strconv.ParseFloat(str, 32)

	return float32(val)
}

// Float64 implements IMValue.
func (mv *MValue) Float64() float64 {
	str := mv.String()

	val, _ := strconv.ParseFloat(str, 32)

	return val
}

// Int implements IMValue.
func (mv *MValue) Int() int {
	str := mv.String()

	val, _ := strconv.ParseInt(str, 10, 64)

	return int(val)
}

// Interface implements IMValue.
func (mv *MValue) Interface() interface{} {
	return mv.Value
}

// String implements IMValue.
func (mv *MValue) String() string {
	if mv.Value == nil {
		return ""
	}

	if val, ok := mv.Value.(string); ok {
		return val
	}

	return ""
}

type IMValue interface {
	String() string
	Bool() bool
	Int() int
	Float64() float64
	Float32() float32
	Interface() interface{}
}

func newMValue(value interface{}) IMValue {
	return &MValue{
		Value: value,
	}
}

func (m M) Get(key string) IMValue {
	return newMValue(m[key])
}

func (m M) Set(key string, val interface{}) M {
	m[key] = val

	return m
}

func (m M) IsEmpty() bool {
	return len(m) == 0
}

func (m M) From(target interface{}) M {
	b, _ := json.Marshal(target)
	json.Unmarshal(b, &m)

	return m
}
