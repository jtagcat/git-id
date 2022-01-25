package ssh_config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	// a := strings.NewReader("test\nkest")
	// err := Decode(a)
}

func TestUnescapeValue(t *testing.T) {
	type out struct {
		values  []string
		comment string
		err     error
	}
	inputs := map[string]out{
		"Inv\"alid":       {nil, "", errInvalidQuoting},
		"Inv'alid":        {nil, "", errInvalidQuoting},
		"\"Valid\"":       {[]string{"Valid"}, "", nil},
		"\"V'alid\"":      {[]string{"V'alid"}, "", nil},
		"String1 String2": {[]string{"String1", "String2"}, "", nil},
		"\"st1\"'s\\t2'":  {[]string{"st1", "s\\t2"}, "", nil},
		"\\":              {[]string{"\\\\"}, "", nil},
		"hello # comment": {[]string{"hello"}, " comment", nil},
		"close#relations": {[]string{"close"}, "relations", nil},
		"#amcomment":      {nil, "amcomment", nil},
		//TODO:
	}
	for input, want := range inputs {
		go func(input string, want out) {
			values, comments, err := unescapeValue(input)
			assert.ErrorIs(t, err, want.err)
			assert.Equal(t, want.values, values)
			assert.Equal(t, want.comment, comments)
		}(input, want)
	}

}
