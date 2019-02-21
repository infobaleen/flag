package flag

import (
	"fmt"
	"strconv"
)

type FlagU64 struct{ value *uint64 }

func (f FlagU64) Value() uint64 { return *f.value }

func (f FlagU64) SetValue(value string) (err error) {
	*f.value, err = strconv.ParseUint(value, 10, 64)
	return
}

func (f FlagSet) u64(flag string, defaultValue uint64, description string, required bool) FlagU64 {
	var handle = FlagU64{value: &defaultValue}
	f[flag] = FlagState{Handle: handle, Type: "uint64", Description: description, Required: required, Default: fmt.Sprint(defaultValue)}
	return handle
}

func (f FlagSet) U64R(flag string, description string) FlagU64 {
	return f.u64(flag, 0, description, true)
}

func (f FlagSet) U64D(flag string, defaultValue uint64, description string) FlagU64 {
	return f.u64(flag, defaultValue, description, false)
}

type FlagU32 struct{ value *uint32 }

func (f FlagU32) Value() uint32 { return *f.value }

func (f FlagU32) SetValue(value string) (err error) {
	tmp, err := strconv.ParseUint(value, 10, 32)
	*f.value = uint32(tmp)
	return
}

func (f FlagSet) u32(flag string, defaultValue uint32, description string, required bool) FlagU32 {
	var handle = FlagU32{value: &defaultValue}
	f[flag] = FlagState{Handle: handle, Type: "uint32", Description: description, Required: required, Default: fmt.Sprint(defaultValue)}
	return handle
}

func (f FlagSet) U32R(flag string, description string) FlagU32 {
	return f.u32(flag, 0, description, true)
}

func (f FlagSet) U32D(flag string, defaultValue uint32, description string) FlagU32 {
	return f.u32(flag, defaultValue, description, false)
}

type FlagU8 struct{ value *uint8 }

func (f FlagU8) Value() uint8 { return *f.value }

func (f FlagU8) SetValue(value string) (err error) {
	tmp, err := strconv.ParseUint(value, 10, 8)
	*f.value = uint8(tmp)
	return
}


func (f FlagSet) u8(flag string, defaultValue uint8, description string, required bool) FlagU8 {
	var handle = FlagU8{value: &defaultValue}
	f[flag] = FlagState{Handle: handle, Type: "uint8", Description: description, Required: required, Default: fmt.Sprint(defaultValue)}
	return handle
}

func (f FlagSet) U8R(flag string, description string) FlagU8 {
	return f.u8(flag, 0, description, true)
}

func (f FlagSet) U8D(flag string, defaultValue uint8, description string) FlagU8 {
	return f.u8(flag, defaultValue, description, false)
}
