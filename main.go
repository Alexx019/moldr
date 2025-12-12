package main

import (
	"fmt"
	"moldr/src/commands"
	"moldr/src/elements"
	"moldr/src/services"
	"moldr/src/utils"

	"os"
)

func main() {
	if err := utils.CheckMainArgs(); err != nil {
		fmt.Println(err)
		return
	}

	err := loadApp()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer saveApp()

	command := os.Args[1]

	switch command {
	case "help", "--help", "-h":
		help, err := services.ReadHelp()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(help)
		}
	case "version", "--version", "-v":
		fmt.Println("moldr v0.0.2")
	default:
		processCommand(command)
	}
}

func processCommand(command string) {
	var err error
	switch command {
	case "ls", "list":
		if err := utils.CheckListArgs(); err != nil {
			fmt.Println(err)
			return
		}
		commands.ListIngots()
	case "new":
		if err := utils.CheckNewArgs(); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.NewIngot(os.Args[2])
		if err == nil {
			fmt.Printf("Ingot created successfully\n\n")
		}
		commands.ListIngots()
	case "del", "delete":
		if err := utils.CheckDelArgs(); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.DeleteIngot(os.Args[2])
		if err == nil {
			fmt.Printf("Ingot deleted successfully\n\n")
		}
		commands.ListIngots()
	case "run", "start":
		if err := utils.CheckRunArgs(); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.RunIngot(os.Args[2])
		if err == nil {
			fmt.Printf("Ingot started successfully\n\n")
		}
	case "stop":
		if err := utils.CheckStopArgs(); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.StopIngot(os.Args[2])
		if err == nil {
			fmt.Printf("Ingot stopped successfully\n\n")
		}
	case "log":
		if err := utils.CheckLogArgs(); err != nil {
			fmt.Println(err)
			return
		}
		err = commands.TailLog(os.Args[2])
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
	err := services.LoadIngots()
	if err != nil {
		return err
	}
	err = services.ReadPIDS()
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
}
