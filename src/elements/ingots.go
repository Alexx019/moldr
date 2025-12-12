package elements

import "moldr/src/utils"

type Ingot struct {
	Name   string
	Path   string
	Port   int
	Status string
	PID    int
}

const (
	initial_port = 8090
)

var (
	port int
)

var Ingots map[string]Ingot

func GetAvailablePort() int {
	port = initial_port
	for _, v := range Ingots {
		if v.Port > port {
			port = v.Port
		}
	}
	port++
	return port
}

func AddIngot(name string) {
	utils.DirWrapper(name, func(dir string) {
		new_ingot := Ingot{
			Name:   name,
			Path:   dir,
			Port:   GetAvailablePort(),
			Status: "stopped",
			PID:    0,
		}
		Ingots[name] = new_ingot
	})
}

func RemoveIngot(name string) {
	delete(Ingots, name)
}

func UpdateIngot(name string, port int, status string, path string, pid int) {
	utils.DirWrapper(name, func(dir string) {
		new_ingot := Ingot{
			Name:   name,
			Path:   dir,
			Port:   port,
			Status: status,
			PID:    pid,
		}
		Ingots[name] = new_ingot
	})
}
