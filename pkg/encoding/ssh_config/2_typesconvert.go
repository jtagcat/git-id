package ssh_config

// Handles conversion between raw and nonraw root structures (this is a wormbox)

type TopLevel struct {
	Key string // enum(4): "" (comment / empty line), Host, Match,
	//                     Include: not recursed, nothing is done (no Children)
	Values []RawValue
	// "# foobar" → " foobar", note the leading space
	Comment                     string
	EncodingKVSeperatorIsEquals bool // "Key=Value" instead of "Key Value"

	Children Keywords
}

// ErrNotImplemented

func DecodeType()

//
//
//

// func validateKeywords(kws Keywords) error {
// 	tkw := reflect.TypeOf(kws)
// 	vkw := reflect.ValueOf(kws)
// 	// .Name()
// 	// .String()
// 	for i := 0; i < tkw.NumField(); i++ {
// 		field := tkw.Field(i)
// 		str := field.Name
// 		def := field.Tag.Get("definition")

// 		err := validateEnumKeyword()

// 		switch def {
// 		default:
// 			return fmt.Errorf("%w: field %s has definition %q", errImpossible, str, def)
// 		case "Flag":

// 		}
// 	}
// 	return nil // TODO:
// }

// .FieldByName

// func parseKV(kv Keyword) (Keyword, error) {
// 	kIndexType, inIndex := keywordIndex[kv.Key]
// 	if !inIndex {
// 		return nil, errInvalidKeyword
// 	}
// 	typ := kIndexType.definition
// 	kv.ValueType = typ

// 	res, err := parseEnum(typ, values)
// 	if err != errInvalidKeyword {
// 		return res, err
// 	}

// 	switch typ {
// 	default:
// 		return nil, errImpossible

// 	}
// 	//
// 	//TODO:
// 	return nil, nil
// }

// func parseEnum(typ string, kv Keyword) (Keyword, error) {
// 	valuePairs, isEnum := enumIndex[typ]
// 	if !isEnum {
// 		return nil, errInvalidKeyword
// 	}

// 	if len(values) != 1 {
// 		return nil, fmt.Errorf("%s: %w: got %d values", typ, errSingleValueOnly, len(values))
// 	}

// 	for _, v := range kv.Values {
// 		for _, vp := range valuePairs {
// 			if typ == v.Value
// 		}
// 	}

// 	for _, v := range valuePairs {
// 		if typ == v.stringName {
// 			kv.ValueType = v.stringName
// 			kv.
// 			return , nil
// 		}
// 	}

// 	var validValues []string
// 	for _, v := range valuePairs {
// 		validValues = append(validValues, v.stringName)
// 	}
// 	return nil, fmt.Errorf("%s, %w: must be one of %q", typ, errInvalidKeyword, validValues)
// }

// TODO: have an overwrite detector for duplicate(ish) key-values

// TODO: transform case, strip 2+ spaces

// func importRawKeyword
