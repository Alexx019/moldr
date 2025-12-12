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
	fmt.Fprintf(w, "NAME\tPATH\tPORT\tSTATUS\tPID\n")
	for k, v := range elements.Ingots {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", strings.ToUpper(k), v.Path, fmt.Sprint(v.Port), strings.ToUpper(v.Status), fmt.Sprint(v.PID))
	}
	w.Flush()
	if len(elements.Ingots) == 0 {
		fmt.Println("This list is empty")
	}
	fmt.Println()
}

func NewIngot(name string) error {
	if _, ok := elements.Ingots[name]; ok {
		return fmt.Errorf("ingot with name %s already exists", name)
	}

	elements.AddIngot(name)
	err := services.NewIngotFolder(name)
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

	// Verificar si ya está corriendo
	if pid, running := services.Pids[name]; running {
		if services.IsProcessRunning(pid) {
			return fmt.Errorf("ingot with name %s is already running", name)
		}
		// Si está en el mapa pero no corriendo, limpiamos el mapa
		services.RemovePID(name)
	}
	services.RunProcess(name)
	return nil
}

func StopIngot(name string) error {
	if !services.ExistsPID(name) {
		return fmt.Errorf("ingot with name %s is not running", name)
	}
	if elements.Ingots[name].Status == "running" {
		services.StopProcess(name)
		elements.UpdateIngot(name, elements.Ingots[name].Port, "stopped", elements.Ingots[name].Path, 0)
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

		// Ir al final del archivo para ver solo lo nuevo (Opcional: quitar Seek para ver todo)
		file.Seek(0, 2)

		reader := bufio.NewReader(file)
		fmt.Printf("Viendo logs de '%s' (Ctrl+C para salir)...\n", name)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// Si llegamos al final, esperamos un poco a que el proceso escriba más
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
