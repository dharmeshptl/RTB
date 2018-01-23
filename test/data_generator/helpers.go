package data_generator

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"

	"github.com/satori/go.uuid"
)

var uniqueStringMap map[string]map[string]bool
var uniqueIntMap map[string]map[int]bool

func init() {
	uniqueStringMap = make(map[string]map[string]bool)
	uniqueIntMap = make(map[string]map[int]bool)
}

func uniqueString(key string, fn func() string, tries int) string {
	result := fn()
	for i := 0; i < tries; i++ {
		if _, ok := uniqueStringMap[key][result]; !ok {
			if _, ok := uniqueStringMap[key]; !ok {
				uniqueStringMap[key] = make(map[string]bool)
			}
			uniqueStringMap[key][result] = true
			return result
		}
		result = fn()
	}
	panic(fmt.Sprintf("Can't find unique string for key: %v, tries: %v", key, tries))
}

func uniqueInt(key string, fn func() int, tries int) int {
	result := fn()
	for i := 0; i < tries; i++ {
		if _, ok := uniqueIntMap[key][result]; !ok {
			if _, ok := uniqueIntMap[key]; !ok {
				uniqueIntMap[key] = make(map[int]bool)
			}
			uniqueIntMap[key][result] = true
			return result
		}
		result = fn()
	}
	panic(fmt.Sprintf("Can't find unique string for key: %v, tries: %v", key, tries))
}

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func setIdIfNotExists(model interface{}) {
	ps := reflect.ValueOf(model)
	if ps.Kind() != reflect.Ptr {
		return
	}
	e := ps.Elem()
	if e.Kind() != reflect.Struct {
		return
	}
	f := e.FieldByName("Id")
	if !f.IsValid() {
		return
	}
	if f.String() == "" && f.CanSet() {
		f.SetString(uuid.NewV4().String())
	}
}

func randomUUID() string {
	return uuid.NewV4().String()
}

func randomStrArrFromList(data []string) []string {
	//TODO: update this
	count := len(data)
	result := make([]string, count, count)
	for i := 0; i < count; i++ {
		result[i] = data[i]
	}

	return result
}

func randomStringFromSlice(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func randomFloat(min, max, precision int) float64 {
	return toFixed(float64(min)+float64(rand.Intn(max-min))*rand.Float64(), precision)
}

func randomUInt() uint {
	return uint(rand.Uint32())
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func timeBeetween(start, end time.Time) time.Time {
	delta := end.Sub(start)
	randomDelta := time.Duration(rand.Int63n(int64(delta.Seconds()))) * time.Second
	return end.Add(-randomDelta)
}

type Data map[string]interface{}

func (d Data) getString(key string, def string) string {
	if d == nil {
		return def
	}
	val, ok := d[key]
	if !ok {
		return def
	}
	return val.(string)
}
