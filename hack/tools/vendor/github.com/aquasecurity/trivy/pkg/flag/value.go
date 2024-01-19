package flag

import (
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"golang.org/x/xerrors"
)

type ValueNormalizeFunc func(string) string

// -- string Value
type customStringValue struct {
	value     *string
	allowed   []string
	normalize ValueNormalizeFunc
}

func newCustomStringValue(val string, allowed []string, fn ValueNormalizeFunc) *customStringValue {
	return &customStringValue{
		value:     &val,
		allowed:   allowed,
		normalize: fn,
	}
}

func (s *customStringValue) Set(val string) error {
	if s.normalize != nil {
		val = s.normalize(val)
	}
	if len(s.allowed) > 0 && !slices.Contains(s.allowed, val) {
		return xerrors.Errorf("must be one of %q", s.allowed)
	}
	s.value = &val
	return nil
}
func (s *customStringValue) Type() string {
	return "string"
}

func (s *customStringValue) String() string { return *s.value }

// -- stringSlice Value
type customStringSliceValue struct {
	value     *[]string
	allowed   []string
	normalize ValueNormalizeFunc
	changed   bool
}

func newCustomStringSliceValue(val, allowed []string, fn ValueNormalizeFunc) *customStringSliceValue {
	return &customStringSliceValue{
		value:     &val,
		allowed:   allowed,
		normalize: fn,
	}
}

func (s *customStringSliceValue) Set(val string) error {
	values := strings.Split(val, ",")
	if s.normalize != nil {
		values = lo.Map(values, func(item string, _ int) string { return s.normalize(item) })
	}
	for _, v := range values {
		if len(s.allowed) > 0 && !slices.Contains(s.allowed, v) {
			return xerrors.Errorf("must be one of %q", s.allowed)
		}
	}
	if !s.changed {
		*s.value = values
	} else {
		*s.value = append(*s.value, values...)
	}
	s.changed = true
	return nil
}

func (s *customStringSliceValue) Type() string {
	return "stringSlice"
}

func (s *customStringSliceValue) String() string {
	if len(*s.value) == 0 {
		// "[]" is not recognized as a zero value
		// cf. https://github.com/spf13/pflag/blob/d5e0c0615acee7028e1e2740a11102313be88de1/flag.go#L553-L565
		return ""
	}
	return "[" + strings.Join(*s.value, ",") + "]"
}

func (s *customStringSliceValue) Append(val string) error {
	s.changed = true
	return s.Set(val)
}

func (s *customStringSliceValue) Replace(val []string) error {
	*s.value = val
	return nil
}

func (s *customStringSliceValue) GetSlice() []string {
	return *s.value
}
