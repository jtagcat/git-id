package ssh_config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeValue(t *testing.T) {
	type out struct {
		values  []RawValue
		comment string
		err     error
	}
	inputs := map[string]out{
		"Inv\"alid":       {nil, "", ErrInvalidQuoting},
		"Inv'alid":        {nil, "", ErrInvalidQuoting},
		"\"Valid\"":       {[]RawValue{{"Valid", 2}}, "", nil},
		"\"V'alid\"":      {[]RawValue{{"V'alid", 2}}, "", nil},
		"String1 String2": {[]RawValue{{"String1", 0}, {"String2", 0}}, "", nil},
		"\"st1\"'s\\t2'":  {[]RawValue{{"st1", 2}, {"s\\t2", 1}}, "", nil},
		"\\":              {[]RawValue{{"\\\\", 0}}, "", nil},
		"hello # comment": {[]RawValue{{"hello", 0}}, " comment", nil},
		"close#relations": {[]RawValue{{"close", 0}}, "relations", nil},
		"#amcomment":      {nil, "amcomment", nil},
		//TODO:
	}
	for input, want := range inputs {
		go func(input string, want out) {
			values, comments, err := DecodeValue(input)
			assert.ErrorIs(t, err, want.err)
			assert.Equal(t, want.values, values)
			assert.Equal(t, want.comment, comments)
		}(input, want)
	}

}
