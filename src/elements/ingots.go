package elements

import "moldr/src/utils"

type Ingot struct {
	Name   string
	Mold   string
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

var Ingots map[string]Ingot // name -> Ingot

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

func IsAvailablePort(port int) bool {
	for _, v := range Ingots {
		if v.Port == port {
			return false
		}
	}
	return true
}

func AddIngot(name string, mold string, port int) {
	if port == 0 {
		port = GetAvailablePort()
	}
	utils.DirWrapper(name, func(dir string) {
		new_ingot := Ingot{
			Name:   name,
			Mold:   mold,
			Path:   dir,
			Port:   port,
			Status: "stopped",
			PID:    0,
		}
		Ingots[name] = new_ingot
	})
}

func RemoveIngot(name string) {
	delete(Ingots, name)
}

func UpdateIngot(name string, mold string, port int, status string, path string, pid int) {
	utils.DirWrapper(name, func(dir string) {
		new_ingot := Ingot{
			Name:   name,
			Mold:   mold,
			Path:   dir,
			Port:   port,
			Status: status,
			PID:    pid,
		}
		Ingots[name] = new_ingot
	})
}

func IsIngot(name string) bool {
	_, exists := Ingots[name]
	return exists
}

func GetCommands(name string) []string {
	mold := Ingots[name].Mold
	filename := Molds[mold].Filename
	serve := Molds[mold].Args.Serve
	port := Molds[mold].Args.Port
	return []string{filename, serve, port}
}
