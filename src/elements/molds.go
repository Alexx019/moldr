package elements

import "fmt"

type Mold struct {
	Name     string `yaml:"name"`
	Filename string `yaml:"filename"`
	Args     Args   `yaml:"args"`
}

type Args struct {
	Serve string `yaml:"serve"`
	Port  string `yaml:"port"`
}

var Molds map[string]Mold // name -> Mold

func AddMold(name string, filename string, serve string, port string) {
	Molds[name] = Mold{
		Name:     name,
		Filename: filename,
		Args:     Args{Serve: serve, Port: port},
	}
}

func RemoveMold(name string) {
	delete(Molds, name)
}

func UpdateMold(name string, filename string, serve string, port string) {
	Molds[name] = Mold{
		Name:     name,
		Filename: filename,
		Args:     Args{Serve: serve, Port: port},
	}
}

func IsMold(name string) bool {
	_, ok := Molds[name]
	return ok
}

func ListMolds() {
	for _, v := range Molds {
		fmt.Println(v)
	}
	if len(Molds) == 0 {
		fmt.Println("This list is empty")
	}
	fmt.Println()
}
