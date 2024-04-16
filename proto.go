package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET = "SET"
	CommandKEY = "KEY"
	CommandFOO = "FOO"
)

type Command interface {
	//
}

type SetCommand struct {
	key, val string
}

func ParseCommand(raw string) (Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(raw))
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read %s\n", v.Type())
		if v.Type() == resp.Array {
			for _, value := range v.Array() {
				switch value.String() {
				case CommandSET:
					// fmt.Println(len(v.Array()), "hey")
					if len(v.Array()) != 3 {
						return nil, fmt.Errorf("invalid number of variables for SET Command")
					}
					cmd := SetCommand{
						key: v.Array()[1].String(),
						val: v.Array()[2].String(),
					}

					fmt.Printf("%+v\n", cmd)
					return cmd, nil
				}
			}
		}
	}

	return "foo", fmt.Errorf("invalid or unkown cmd: %s", raw)
}
