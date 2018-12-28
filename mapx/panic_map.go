package mapx

import (
	"fmt"
)

// PanicMap will panic if key doesn't exit with no default value.
type PanicMap struct {
	Data map[string]string
}

// Get a value from a map
func (pc *PanicMap) Get(key string, defaultValue ...string) string {
	value, ok := pc.Data[key]
	if !ok {
		if defaultValue == nil {
			panic(fmt.Sprintf("map key %s not exist", key))
		}
		return defaultValue[0]
	}
	return value
}
