package ssh_config

import (
	"errors"
)

var (
	ErrInvalidKeyword                 = errors.New("invalid keyword")
	ErrInvalidValue                   = errors.New("invalid value")
	ErrSingleValueOnly                = errors.New("must have exactly 1 value")
	ErrInvalidQuoting                 = errors.New("bad quoting") // TODO: add more info?
	ErrInvalidKeyLocation             = errors.New("keyword must be under Host or Match")
	ErrWarnSingleBackslashTransformed = errors.New("1 or more single backslashes changed to 2 backslashes since OpenSSH ssh_config does this (this always happens: st\\ring → st\\\\ring)")
	ErrImpossible                     = errors.New("situation should be impossible, likely human error, please report")
)

// encoding: indentchar

//func typeCheck()

// func: validate written config at runtime
