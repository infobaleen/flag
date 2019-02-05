package flag

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type FlagSet map[string]FlagState

type FlagState struct {
	Handle      FlagHandle
	Set         bool
	Required    bool
	Type        string
	Default     string
	Description string
}

type FlagHandle interface{}

type FlagHandleSetValue interface {
	FlagHandle
	SetValue(string) error
}

type FlagHandleSet interface {
	FlagHandle
	Set()
}

func NewFlagSet() FlagSet {
	return make(map[string]FlagState)
}

func (f FlagSet) Parse(flags []string) {
	for i := 0; i != len(flags); i++ {
		var flag = flags[i]
		if !strings.HasPrefix(flag, "-") {
			fmt.Fprintf(os.Stderr, "Expected flag, found \"%s\".\n\n", flag)
			f.Usage()
		}
		flag = flag[1:]

		var state = f[flag]
		if state.Handle == nil {
			fmt.Fprintf(os.Stderr, "Unknown flag \"%s\".\n\n", flag)
			f.Usage()
		}
		if state.Set {
			fmt.Fprintf(os.Stderr, "Duplicate flag \"%s\".\n\n", flag)
			f.Usage()
		}

		if handle, ok := state.Handle.(FlagHandleSetValue); ok {
			if len(os.Args) < i+2 {
				fmt.Fprintf(os.Stderr, "Missing value for flag \"%s\".\n\n", flag)
				f.Usage()
			}
			var value = flags[i+1]
			if err := handle.SetValue(value); err != nil {
				fmt.Fprintf(os.Stderr, "Could not parse value \"%s\", as %s.\n\n", value, state.Type)
				f.Usage()
			}
			i++
		} else if handle, ok := state.Handle.(FlagHandleSet); ok {
			handle.Set()
		} else {
			fmt.Fprintf(os.Stderr, "Flag handle does not implement any of the required interfaces.\n\n")
		}
		state.Set = true
		f[flag] = state
	}

	for flag, state := range f {
		if state.Required && !state.Set {
			fmt.Fprintf(os.Stderr, "Required flag \"%s\" not specified.\n\n", flag)
			f.Usage()
		}
	}
}

func (f FlagSet) Usage() {
	fmt.Fprintln(os.Stderr, "Flags:")
	var flags []string
	var maxLen int
	for flag, state := range f {
		flags = append(flags, flag)
		if l := len(flag) + len(state.Type); maxLen < l {
			maxLen = l
		}
	}
	sort.Strings(flags)
	for _, flag := range flags {
		var state = f[flag]
		var def = "required"
		if !f[flag].Required {
			def = f[flag].Default
		}
		var space = strings.Repeat(" ", maxLen-len(flag)-len(state.Type))
		fmt.Fprintf(os.Stderr, "  -%s  %s%s: %s (%s)\n", flag, space, state.Type, state.Description, def)
	}
	os.Exit(1)
}
