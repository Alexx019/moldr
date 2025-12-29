package main

import (
	"fmt"
	"moldr/src/commands"
	"moldr/src/elements"
	"moldr/src/services"
	"moldr/src/utils"
	"strconv"
	"strings"

	"os"
)

func main() {
	if err := utils.CheckMainArgs(os.Args); err != nil {
		fmt.Println(err)
		return
	}

	// Load molds and ingots
	err := loadApp()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer saveApp()

	command := os.Args[1]

	// System commands
	switch command {
	case "help", "--help", "-h":
		help, err := services.ReadHelp()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(help)
		}
	case "version", "--version", "-v":
		fmt.Println("moldr v1.0.0")

	// Project commands
	default:
		processCommand(command)
	}
}

func processCommand(command string) {
	var err error
	args := os.Args

	switch command {
	// List Ingots
	case "ls", "list":
		if err := utils.CheckListArgs(args); err != nil {
			fmt.Println(err)
			return
		}
		commands.ListIngots()

	// New Ingot / Mold
	case "new":
		if err := utils.CheckNewIngotArgs(args); err != nil {
			fmt.Println(err)
			return
		}
		var ingotName string
		var moldName string
		var port int

		ingotName = os.Args[2]
		moldName = strings.Split(os.Args[3], "=")[1]
		if len(os.Args) == 5 {
			port, _ = strconv.Atoi(strings.Split(os.Args[4], "=")[1])
		}
		if !elements.IsMold(moldName) {
			fmt.Printf("Mold %s does not exist\n\n", moldName)
			fmt.Println("Available molds:")
			elements.ListMolds()
			return
		}
		if !elements.IsAvailablePort(port) {
			fmt.Printf("Port %d is not available\n\n", port)
			commands.ListIngots()
			return
		}
		err = commands.NewIngot(ingotName, elements.Molds[moldName], port)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Ingot created successfully\n\n")
		commands.ListIngots()

	// Delete Ingot
	case "del", "delete":
		if err := utils.CheckDelArgs(args); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.DeleteIngot(args[2])
		if err == nil {
			fmt.Printf("Ingot deleted successfully\n\n")
		}
		commands.ListIngots()

	// Run Ingot
	case "run", "start":
		if err := utils.CheckRunArgs(args); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.RunIngot(args[2])
		if err == nil {
			fmt.Printf("Ingot started successfully\n\n")
		}

	// Stop Ingot
	case "stop":
		if err := utils.CheckStopArgs(args); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.StopIngot(args[2])
		if err == nil {
			fmt.Printf("Ingot stopped successfully\n\n")
		}

	// Tail Log
	case "logs", "log":
		if err := utils.CheckLogsArgs(args); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.TailLog(args[2])

	default:
		err = fmt.Errorf("invalid command \nrun 'moldr help' for more information")
	}

	if err != nil {
		fmt.Println(err)
	}
}

func loadApp() error {
	services.Pids = make(map[string]int)
	elements.Ingots = make(map[string]elements.Ingot)
	elements.Molds = make(map[string]elements.Mold)
	err := services.LoadIngots()
	if err != nil {
		return err
	}
	err = services.ReadPIDS()
	if err != nil {
		return err
	}
	err = services.LoadMolds()
	if err != nil {
		return err
	}
	elements.GetAvailablePort()
	return nil
}

func saveApp() {
	err := services.SaveIngots()
	if err != nil {
		fmt.Println(err)
	}
	err = services.WritePIDS()
	if err != nil {
		fmt.Println(err)
	}
	err = services.SaveMolds()
	if err != nil {
		fmt.Println(err)
	}
}
