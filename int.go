package flag

import (
	"fmt"
	"strconv"
)

type FlagI64 struct{ value *int64 }

func (f FlagI64) Value() int64 { return *f.value }

func (f FlagI64) SetValue(value string) (err error) {
	*f.value, err = strconv.ParseInt(value, 10, 64)
	return
}

func (f FlagSet) i64(flag string, defaultValue int64, description string, required bool) FlagI64 {
	var handle = FlagI64{value: &defaultValue}
	f[flag] = FlagState{Handle: handle, Type: "int64", Description: description, Required: required, Default: fmt.Sprint(defaultValue)}
	return handle
}

func (f FlagSet) I64R(flag string, description string) FlagI64 {
	return f.i64(flag, 0, description, true)
}

func (f FlagSet) I64D(flag string, defaultValue int64, description string) FlagI64 {
	return f.i64(flag, defaultValue, description, false)
}
