package iniflags

import (
	"flag"
	"strings"
	"time"
)

type Subcommand struct {
	Name    string
	flagSet *flag.FlagSet
}

func NewSubcommand(name string) *Subcommand {
	if _, ok := subcommands[name]; ok {
		panic("iniflags: cannot add two subcommands with the same commandname: " + name)
	}
	flagSet := flag.NewFlagSet(name, flag.ExitOnError)
	subcommand := Subcommand{Name: name, flagSet: flagSet}
	subcommands[name] = &subcommand

	// import iniflags flags
	subcommand.ImportFlag("allowUnknownFlags")
	subcommand.ImportFlag("config")
	subcommand.ImportFlag("configUpdateInterval")
	subcommand.ImportFlag("dumpflags")

	return &subcommand
}

func (s *Subcommand) ImportFlag(name string) {
	if f := flag.Lookup(name); f != nil {
		s.flagSet.Var(f.Value, f.Name, f.Usage)
	}
}

func (s *Subcommand) ImportFlags() {
	flag.VisitAll(func(f *flag.Flag) {
		s.flagSet.Var(f.Value, f.Name, f.Usage)
	})
}

func (s *Subcommand) ImportFlagsWithPrefix(prefix string) {
	flag.VisitAll(func(f *flag.Flag) {
		if strings.HasPrefix(f.Name, prefix) {
			s.flagSet.Var(f.Value, f.Name, f.Usage)
		}
	})
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (s *Subcommand) BoolVar(p *bool, name string, value bool, usage string) {
	s.flagSet.BoolVar(p, name, value, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (s *Subcommand) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	s.flagSet.BoolVar(p, name, value, usage)
	return p
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (s *Subcommand) IntVar(p *int, name string, value int, usage string) {
	s.flagSet.IntVar(p, name, value, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (s *Subcommand) Int(name string, value int, usage string) *int {
	p := new(int)
	s.flagSet.IntVar(p, name, value, usage)
	return p
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (s *Subcommand) Int64Var(p *int64, name string, value int64, usage string) {
	s.flagSet.Int64Var(p, name, value, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (s *Subcommand) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	s.flagSet.Int64Var(p, name, value, usage)
	return p
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (s *Subcommand) UintVar(p *uint, name string, value uint, usage string) {
	s.flagSet.UintVar(p, name, value, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (s *Subcommand) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	s.flagSet.UintVar(p, name, value, usage)
	return p
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (s *Subcommand) Uint64Var(p *uint64, name string, value uint64, usage string) {
	s.flagSet.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (s *Subcommand) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	s.flagSet.Uint64Var(p, name, value, usage)
	return p
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (s *Subcommand) StringVar(p *string, name string, value string, usage string) {
	s.flagSet.StringVar(p, name, value, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (s *Subcommand) String(name string, value string, usage string) *string {
	p := new(string)
	s.flagSet.StringVar(p, name, value, usage)
	return p
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (s *Subcommand) Float64Var(p *float64, name string, value float64, usage string) {
	s.flagSet.Float64Var(p, name, value, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (s *Subcommand) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	s.flagSet.Float64Var(p, name, value, usage)
	return p
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (s *Subcommand) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	s.flagSet.DurationVar(p, name, value, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (s *Subcommand) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	s.flagSet.DurationVar(p, name, value, usage)
	return p
}
