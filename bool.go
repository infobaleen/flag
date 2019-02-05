package flag

type FlagBool struct{ value *bool }

func (f FlagBool) Value() bool { return *f.value }

func (f FlagBool) Set() {
	*f.value = true
}

func (f FlagSet) Bool(flag string, description string) FlagBool {
	var handle = FlagBool{value: new(bool)}
	f[flag] = FlagState{Handle: handle, Type: "bool", Description: description, Default: "false"}
	return handle
}
