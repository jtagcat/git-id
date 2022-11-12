package ssh_config

// var testingValidConfigString = "ssh: Could not resolve hostname invalid: Temporary failure in name resolution\r\n"

// func TestKeywordArgLength(t *testing.T) {
// 	tc := path.Join(t.TempDir(), "test.conf")
// 	for keyword, indexDef := range keywordIndex {
// 		if keyword != "LocalForward" { //TODO: can not get it to a valid config
// 			validArg, specialSecond, err := exampleValidValueForTesting(indexDef.definition)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			// min..max args, should succeed
// 			for ac := indexDef.minArgs; ac <= indexDef.maxArg; ac++ {
// 				// construct arguments
// 				var testArgs string
// 				for i := 1; i <= ac; i++ {
// 					if specialSecond != "" && i == 2 {
// 						testArgs += specialSecond + " "
// 					} else {
// 						testArgs += validArg + " "
// 					}

// 				}
// 				os.WriteFile(tc, []byte("Host *\n"+keyword+"="+testArgs), 0644)
// 				out, _ := exec.Command("ssh", "-T", "-F", tc, "invalid").CombinedOutput()
// 				if testingValidConfigString != string(out) {
// 					t.Fatalf("%q: min..max at %d failed: %v", keyword, ac, string(out))
// 				}
// 			}
// 			// min out of bounds (should fail)
// 			if indexDef.minArgs > 0 {
// 				var testArgs string
// 				for i := indexDef.minArgs - 1; i > 0 && i > indexDef.minArgs-5; i-- {
// 					if i != indexDef.minArgs-1 {
// 						testArgs += " "
// 					}
// 					testArgs += validArg
// 				}
// 				os.WriteFile(tc, []byte("Host *\n"+keyword+"="+testArgs), 0644)
// 				out, _ := exec.Command("ssh", "-T", "-F", tc, "invalid").CombinedOutput()
// 				if testingValidConfigString == string(out) {
// 					t.Fatalf("%q: min oob %d succeeded: %v", keyword, indexDef.minArgs, string(out))
// 				}
// 			}
// 			// max oob (should fail)
// 			if indexDef.maxArg >= 0 {
// 				var testArgs string
// 				for i := indexDef.maxArg + 1; i < indexDef.maxArg+5; i++ {
// 					testArgs += validArg + " "
// 				}
// 				os.WriteFile(tc, []byte("Host *\n"+keyword+"="+testArgs), 0644)
// 				out, _ := exec.Command("ssh", "-T", "-F", tc, "invalid").CombinedOutput()
// 				if testingValidConfigString == string(out) {
// 					t.Fatalf("%q: max oob %q succeeded: %v", keyword, indexDef.maxArg, string(out))
// 				}
// 			}
// 		}
// 	}
// 	//TODO: test indexDef.repeatable: if it overwrites or adds
// 	//TODO: is 'Key= ' treated differently than 'Key='?

// }

// // secondmustbe is a valid value for the second placement in array, rarely returned non-empty
// func exampleValidValueForTesting(definition string) (value, secondmustbe string, err error) {
// 	switch definition {
// 	case "YesNoAsk", "multiDefineStringSlice", "RequestTTY", "StrictHostkey", "Tunnel", "Compression", "boolString", "CanonicalizeHostname", "ControlMaster", "indifferentString", "YesNoAskConfirm", "string", "stringSlice", "tunnelDevice", "Flag", "controlPersist", "unsupported", "deprecated", "deprecatedHidden", "csvStringSlice", "cipher":
// 		return "yes", "", nil
// 	case "AddressFamily", "permitRemoteOpen":
// 		return "any", "", nil
// 	case "SessionType":
// 		return "none", "", nil
// 	case "LogFacility":
// 		return "user", "", nil
// 	case "canonicalizeCNAMEs":
// 		return "foo:bar", "", nil
// 	case "nzint32", "uint16", "duration", "dynamicForward", "permoctal":
// 		return "0777", "", nil
// 	case "Hash":
// 		return "sha256", "", nil
// 	case "LogLevel":
// 		return "error", "", nil
// 	case "rekeyLimit":
// 		return "16", "", nil
// 	case "ipqos":
// 		return "cs1", "", nil
// 	case "forward":
// 		return "1337", "host:1337", nil
// 	default:
// 		return "", "", fmt.Errorf("undefined definition in test: %q", definition)
// 	}
// }
// func TestKeywordArgRepetition(t *testing.T) {

// }
