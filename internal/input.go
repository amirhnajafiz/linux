package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type input struct {
	reader *bufio.Reader
}

func newInput() input {
	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		panic(fmt.Errorf("problem in reader"))
	}

	return input{
		reader: reader,
	}
}

func (i input) get() string {
	cmd, _ := i.reader.ReadString('\n')
	cmd = strings.Trim(cmd, "\n")

	return cmd
}

func (i input) decode(cmd string) (map[string]string, error) {
	pack := make(map[string]string)

	parts := strings.Split(cmd, " ")
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	pack["command"] = parts[0]

	for index := 1; index < len(parts); index += 2 {
		if strings.HasPrefix(parts[index], "--") {
			if index+1 >= len(parts) {
				return nil, fmt.Errorf("mismatch in key value pairs")
			}

			if strings.HasPrefix(parts[index+1], "--") {
				return nil, fmt.Errorf("no values for '%s' flag were given", parts[index])
			}

			pack[parts[index]] = parts[index+1]
		} else {
			return nil, fmt.Errorf("given flag '%s' not supported", parts[index])
		}
	}

	return pack, nil
}
