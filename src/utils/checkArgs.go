package utils

import (
	"fmt"
	"strings"
)

type ArgsError struct {
	Command string
	Err     error
}

func (e *ArgsError) Error() string {
	return e.Err.Error()
}

func CheckMainArgs(args []string) error {
	if len(args) < 2 {
		return &ArgsError{"<command>", fmt.Errorf("usage: modlr <command>, try 'modlr help' for more information")}
	}
	return nil
}

func CheckListArgs(args []string) error {
	if len(args) > 2 {
		return &ArgsError{"list", fmt.Errorf("usage: modlr list")}
	}
	return nil
}

func CheckNewIngotArgs(args []string) error {
	// moldr new ingot <name> --mold=<mold_name> (--port=<port>)
	if len(args) < 4 || len(args) > 6 {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
	}
	var name string
	var moldFlag string
	var portFlag string
	if name = args[2]; name == "" {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
	}
	if moldFlag = args[3]; !strings.HasPrefix(moldFlag, "--mold=") {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
	}
	if len(args) == 5 {
		if portFlag = args[4]; !strings.HasPrefix(portFlag, "--port=") {
			return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
		}
	}
	return nil
}

func CheckDelArgs(args []string) error {
	if len(args) < 3 {
		return &ArgsError{"del", fmt.Errorf("usage: modlr del <name> ...<name>")}
	}
	return nil
}

func CheckRunArgs(args []string) error {
	if len(args) < 3 || len(args) > 3 {
		return &ArgsError{"run", fmt.Errorf("usage: modlr run <name>")}
	}
	return nil
}

func CheckStopArgs(args []string) error {
	if len(args) < 3 || len(args) > 3 {
		return &ArgsError{"stop", fmt.Errorf("usage: modlr stop <name>")}
	}
	return nil
}

func CheckLogsArgs(args []string) error {
	if len(args) < 3 || len(args) > 3 {
		return &ArgsError{"log", fmt.Errorf("usage: modlr log <name>")}
	}
	return nil
}
