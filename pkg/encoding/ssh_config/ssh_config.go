package ssh_config

import (
	"errors"
)

var (
	errInvalidKeyword                 = errors.New("invalid keyword")
	errInvalidValue                   = errors.New("invalid value")
	errSingleValueOnly                = errors.New("must have exactly 1 value")
	errInvalidQuoting                 = errors.New("bad quoting") // TODO: add more info?
	errInvalidKeyLocation             = errors.New("keyword must be under Host or Match")
	errWarnSingleBackslashTransformed = errors.New("1 or more single backslashes changed to 2 backslashes since OpenSSH ssh_config does this (this always happens: st\\ring → st\\\\ring)")
	errImpossible                     = errors.New("situation should be impossible, likely human error, please report")
)

// encoding: indentchar

//func typeCheck()

// func: validate written config at runtime
