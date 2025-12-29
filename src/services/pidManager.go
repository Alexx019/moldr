package services

import (
	"fmt"
	"moldr/src/elements"
	"moldr/src/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

var Pids map[string]int

func AddPID(name string, pid int) {
	Pids[name] = pid
}

func RemovePID(name string) {
	delete(Pids, name)
}

func UpdatePID(name string, pid int) {
	Pids[name] = pid
}

func ExistsPID(name string) bool {
	_, exists := Pids[name]
	return exists
}

func RunProcess(name string) error {
	return utils.DirWrapperWithError(name, func(dir string) error {
		logPath := filepath.Join(dir, "logs", "out.log")
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer logFile.Close()

		// Get mold especific commands
		commands := elements.GetCommands(name)
		filename := commands[0]
		serve := commands[1]
		port := commands[2]

		filename = filepath.Join(dir, "data", filename)                                     // Adjust path
		port = strings.ReplaceAll(port, "{{PORT}}", fmt.Sprint(elements.Ingots[name].Port)) // Replace port

		cmd := exec.Command(filename, serve, port)
		cmd.Stdout = logFile
		cmd.Stderr = logFile

		cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: CREATE_NO_WINDOW,
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		// save pid
		AddPID(name, cmd.Process.Pid)
		elements.UpdateIngot(name, elements.Ingots[name].Mold, elements.Ingots[name].Port, "running", elements.Ingots[name].Path, cmd.Process.Pid)
		return nil
	})
}

func StopProcess(name string) error {
	proc, err := os.FindProcess(Pids[name])
	if err == nil {
		err := proc.Kill()
		if err != nil {
			return err
		}
		RemovePID(name)
	}
	return nil
}

func IsProcessRunning(pid int) bool {
	// Trick for Windows: Search in tasklist
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/NH")
	out, _ := cmd.CombinedOutput()
	return len(out) > 0 && string(out) != "" && strings.Contains(string(out), fmt.Sprintf("%d", pid))
}
