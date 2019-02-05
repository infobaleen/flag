package flag

import (
	"fmt"
	"strconv"
)

type FlagF64 struct{ value *float64 }

func (f FlagF64) Value() float64 { return *f.value }

func (f FlagF64) SetValue(value string) (err error) {
	*f.value, err = strconv.ParseFloat(value, 64)
	return
}

func (f FlagSet) f64(flag string, defaultValue float64, description string, required bool) FlagF64 {
	var handle = FlagF64{value: &defaultValue}
	f[flag] = FlagState{Handle: handle, Type: "float64", Description: description, Required: required, Default: fmt.Sprint(defaultValue)}
	return handle
}

func (f FlagSet) F64R(flag string, description string) FlagF64 {
	return f.f64(flag, 0, description, true)
}

func (f FlagSet) F64D(flag string, defaultValue float64, description string) FlagF64 {
	return f.f64(flag, defaultValue, description, false)
}
