package utils

import (
	"fmt"
	"os"
)

type ArgsError struct {
	Command string
	Err     error
}

func (e *ArgsError) Error() string {
	return fmt.Sprintf("usage: modlr %s, try 'modlr help' for more information", e.Command)
}

func CheckMainArgs() error {
	if len(os.Args) < 2 {
		return &ArgsError{"<command>", fmt.Errorf("usage: modlr <command>, try 'modlr help' for more information")}
	}
	return nil
}

func CheckListArgs() error {
	if len(os.Args) > 2 {
		return &ArgsError{"list", fmt.Errorf("usage: modlr list")}
	}
	return nil
}

func CheckNewArgs() error {
	if len(os.Args) < 3 {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name>")}
	}
	return nil
}

func CheckDelArgs() error {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		return &ArgsError{"del", fmt.Errorf("usage: modlr del <name>")}
	}
	return nil
}

func CheckRunArgs() error {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		return &ArgsError{"run", fmt.Errorf("usage: modlr run <name>")}
	}
	return nil
}

func CheckStopArgs() error {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		return &ArgsError{"stop", fmt.Errorf("usage: modlr stop <name>")}
	}
	return nil
}

func CheckLogArgs() error {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		return &ArgsError{"log", fmt.Errorf("usage: modlr log <name>")}
	}
	return nil
}
