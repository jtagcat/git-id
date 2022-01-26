package ssh_config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	// a := strings.NewReader("test\nkest")
	// err := Decode(a)
}

func TestDecodeValue(t *testing.T) {
	type out struct {
		values  []Value
		comment string
		err     error
	}
	inputs := map[string]out{
		"Inv\"alid":       {nil, "", errInvalidQuoting},
		"Inv'alid":        {nil, "", errInvalidQuoting},
		"\"Valid\"":       {[]Value{{"Valid", 2, ""}}, "", nil},
		"\"V'alid\"":      {[]Value{{"V'alid", 2, ""}}, "", nil},
		"String1 String2": {[]Value{{"String1", 0, ""}, {"String2", 0, ""}}, "", nil},
		"\"st1\"'s\\t2'":  {[]Value{{"st1", 2, ""}, {"s\\t2", 1, ""}}, "", nil},
		"\\":              {[]Value{{"\\\\", 0, ""}}, "", nil},
		"hello # comment": {[]Value{{"hello", 0, ""}}, " comment", nil},
		"close#relations": {[]Value{{"close", 0, ""}}, "relations", nil},
		"#amcomment":      {nil, "amcomment", nil},
		//TODO:
	}
	for input, want := range inputs {
		go func(input string, want out) {
			values, comments, err := decodeValue(input)
			assert.ErrorIs(t, err, want.err)
			assert.Equal(t, want.values, values)
			assert.Equal(t, want.comment, comments)
		}(input, want)
	}

}
