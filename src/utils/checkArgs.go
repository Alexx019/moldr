package utils

import (
	"fmt"
	"os"
	"strings"
)

type ArgsError struct {
	Command string
	Err     error
}

func (e *ArgsError) Error() string {
	return e.Err.Error()
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

func CheckNewIngotArgs() error {
	// moldr new ingot <name> --mold=<mold_name> (--port=<port>)
	if len(os.Args) < 4 || len(os.Args) > 6 {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
	}
	var name string
	var moldFlag string
	var portFlag string
	if name = os.Args[2]; name == "" {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
	}
	if moldFlag = os.Args[3]; !strings.HasPrefix(moldFlag, "--mold=") {
		return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
	}
	if len(os.Args) == 5 {
		if portFlag = os.Args[4]; !strings.HasPrefix(portFlag, "--port=") {
			return &ArgsError{"new", fmt.Errorf("usage: modlr new <name> --mold=<mold_name> (--port=<port>)")}
		}
	}
	return nil
}

func CheckDelArgs() error {
	if len(os.Args) < 3 {
		return &ArgsError{"del", fmt.Errorf("usage: modlr del <name> ...<name>")}
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

func CheckLogsArgs() error {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		return &ArgsError{"log", fmt.Errorf("usage: modlr log <name>")}
	}
	return nil
}
