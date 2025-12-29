package commands

import (
	"bufio"
	"fmt"
	"io"
	"moldr/src/elements"
	"moldr/src/services"
	"moldr/src/utils"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"
)

const (
	CREATE_NO_WINDOW = 0x08000000
)

func ListIngots() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(w, "NAME\tMOLD\tPATH\tPORT\tSTATUS\tPID\n")
	for k, v := range elements.Ingots {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", strings.ToUpper(k), v.Mold, v.Path, fmt.Sprint(v.Port), strings.ToUpper(v.Status), fmt.Sprint(v.PID))
	}
	w.Flush()
	if len(elements.Ingots) == 0 {
		fmt.Println("This list is empty")
	}
	fmt.Println()
}

func NewIngot(name string, mold elements.Mold, port int) error {
	if _, ok := elements.Ingots[name]; ok {
		return fmt.Errorf("ingot with name %s already exists", name)
	}

	elements.AddIngot(name, mold.Name, port)
	err := services.NewIngotFolder(name, mold)
	if err != nil {
		return err
	}
	return nil
}

func DeleteIngot(name string) error {
	if _, ok := elements.Ingots[name]; !ok {
		return fmt.Errorf("ingot with name %s does not exist", name)
	}
	elements.RemoveIngot(name)
	services.RemovePID(name)
	services.RemoveIngotFolder(name)
	elements.GetAvailablePort()
	return nil
}

func RunIngot(name string) error {
	if _, ok := elements.Ingots[name]; !ok {
		return fmt.Errorf("ingot with name %s does not exist", name)
	}

	// Check if is running
	if pid, running := services.Pids[name]; running {
		if services.IsProcessRunning(pid) {
			return fmt.Errorf("ingot with name %s is already running", name)
		}
		// If it is in the map but not running, clean the map
		services.RemovePID(name)
	}
	err := services.RunProcess(name)
	if err != nil {
		return err
	}
	return nil
}

func StopIngot(name string) error {
	if !services.ExistsPID(name) {
		return fmt.Errorf("ingot with name %s is not running", name)
	}
	if elements.Ingots[name].Status == "running" {
		services.StopProcess(name)
		elements.UpdateIngot(name, elements.Ingots[name].Mold, elements.Ingots[name].Port, "stopped", elements.Ingots[name].Path, 0)
	}
	return nil
}

func TailLog(name string) error {
	return utils.DirWrapperWithError(name, func(dir string) error {
		logPath := filepath.Join(dir, "logs", "out.log")
		file, err := os.Open(logPath)
		if err != nil {
			return fmt.Errorf("no se encontró archivo de log (¿El ingot ha corrido alguna vez?)")
		}
		defer file.Close()

		// Read all written and wait for the process to write more
		stat, _ := file.Stat()
		if stat.Size() > 2048 {
			file.Seek(-2048, io.SeekEnd)
		}
		file.Seek(0, io.SeekStart)

		reader := bufio.NewReader(file)
		fmt.Printf("Viendo logs de '%s' (Ctrl+C para salir)...\n", name)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// If we reach the end, wait a little while the process writes more
					time.Sleep(500 * time.Millisecond)
					continue
				}
				break
			}
			fmt.Print(line)
		}
		return nil
	})
}
