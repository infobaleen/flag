package flag

type FlagStr struct{ value *string }

func (f FlagStr) Value() string { return *f.value }

func (f FlagStr) SetValue(value string) (err error) {
	*f.value = value
	return
}

func (f FlagSet) str(flag string, defaultValue string, description string, required bool) FlagStr {
	var handle = FlagStr{value: &defaultValue}
	f[flag] = FlagState{Handle: handle, Type: "string", Description: description, Required: required, Default: `"` + defaultValue + `"`}
	return handle
}

func (f FlagSet) StrR(flag string, description string) FlagStr {
	return f.str(flag, "", description, true)
}

func (f FlagSet) StrD(flag string, defaultValue string, description string) FlagStr {
	return f.str(flag, defaultValue, description, false)
}
