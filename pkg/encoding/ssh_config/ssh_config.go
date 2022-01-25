package ssh_config

import (
	"bufio"
	"fmt"
	"io"
)

func Decode(data io.Reader) error {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	return scanner.Err() // no need to wrap
}
