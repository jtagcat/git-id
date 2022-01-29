package ssh_config

// func Decode(data io.Reader) ([]TopLevel, error) {
// 	var deep bool // under a host or match
// 	var tl []TopLevel
// 	var cl TopLevel // current level
// 	var i int       // can't integrate it in to for?
// 	scanner := bufio.NewScanner(data)
// 	for scanner.Scan() {
// 		i++
// 		rkw, err := parseLine(strings.ToValidUTF8(scanner.Text(), ""))
// 		if err == errInvalidQuoting { // crash and burn
// 			err = fmt.Errorf("while parsing line %d: %w", i, err)
// 		}
// 		if err != nil {
// 			return tl, err
// 		}

// 		switch rkw.Key {
// 		case "Host", "Match", "Include":
// 			if cl.Key != "" { // flush
// 				tl = append(tl, cl)
// 			}
// 			cl.Key = rkw.Key
// 			cl.Values = rkw.Values
// 			cl.Comment = rkw.Comment
// 			cl.EncodingKVSeperatorIsEquals = rkw.EncodingKVSeperatorIsEquals
// 			cl.Children = Keywords{}
// 			if rkw.Key == "Include" {
// 				deep = false
// 			} else {
// 				deep = true
// 			}
// 		default:
// 			if !deep {
// 				return tl, fmt.Errorf("while parsing line %d: %w", i, errInvalidKeyLocation)
// 			}

// 			// kw, err :=
// 			// cl.Children = append(cl.Children, kw)
// 			//

// 			//

// 		}
// 	}
// 	return tl, scanner.Err() // no need to wrap
// }

// func rawToKeyword(rkw rawKeyword, kws *Keywords) error {
// 	tkw := reflect.TypeOf(kws)
// 	tkw.FieldByName(rkw.Key)
// }
