package services

import (
	"encoding/gob"
	"io"
	"moldr/src/elements"
	"moldr/src/utils"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	data_path   = ".data"
	pids_file   = "pids.bin"
	ingots_file = "ingots.bin"
	molds_dir   = "molds"
)

func LoadIngots() error {
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		file, err := os.Open(filepath.Join(dir, ingots_file))
		if err != nil {
			// If file doesn't exist, return
			return err
		}
		defer file.Close()

		decoder := gob.NewDecoder(file)
		if err := decoder.Decode(&elements.Ingots); err != nil {
			return err
		}
		return nil
	})
	return err
}

func SaveIngots() error {
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		file, err := os.Create(filepath.Join(dir, ingots_file))
		if err != nil {
			return err
		}
		defer file.Close()

		encoder := gob.NewEncoder(file)
		if err := encoder.Encode(elements.Ingots); err != nil {
			return err
		}
		return nil
	})
	return err
}

func NewIngotFolder(name string, mold elements.Mold) error {
	err := utils.DirWrapperWithError(name, func(dir string) error {
		if err := os.Mkdir(dir, 0755); err != nil { // Check if dir exists
			return err
		}
		if err := os.Mkdir(filepath.Join(dir, "logs"), 0755); err != nil { // Check if logs dir exists
			return err
		}
		if err := os.Mkdir(filepath.Join(dir, "data"), 0755); err != nil { // Check if data dir exists
			return err
		}
		if _, err := os.Create(filepath.Join(dir, "logs", "log.txt")); err != nil { // Check if log.txt exists
			return err
		}
		// Copy the pocketbase executable into the ingot's data directory
		srcPath := MoldPath(mold.Name)
		dstPath := filepath.Join(dir, "data", mold.Filename)

		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}
		return nil
	})
	return err
}

func RemoveIngotFolder(name string) error {
	err := utils.DirWrapperWithError(name, func(dir string) error {
		err := os.RemoveAll(dir)
		return err
	})
	return err
}

func LoadMolds() error {
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		// Open the mold dir and get a list of the dirs inside it
		moldDir := filepath.Join(dir, molds_dir)
		moldDirs, err := os.ReadDir(moldDir)
		if err != nil {
			return err
		}
		for _, mDir := range moldDirs {
			if mDir.IsDir() {
				// Open the mold dir and get a list of the files inside it
				moldFile, err := os.Open(filepath.Join(dir, molds_dir, mDir.Name(), "mold.yaml"))
				if err != nil {
					return err
				}
				defer moldFile.Close()

				/* AN EXAMPLE OF MOLD.YAML
				name: pocketbase
				filename: pocketbase.exe
				args:
					serve: serve
					port: --http=127.0.0.1:{{PORT}}
				*/

				yamlContent, err := io.ReadAll(moldFile)
				if err != nil {
					return err
				}
				var mold elements.Mold
				if err := yaml.Unmarshal(yamlContent, &mold); err != nil {
					return err
				}
				elements.Molds[mDir.Name()] = mold
			}
		}
		return nil
	})
	return err
}

// TODO: Implement a way to use other molds and then save them
func SaveMolds() error {
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		/*
			file, err := os.Create(filepath.Join(dir, molds_file))
			if err != nil {
				return err
			}
			defer file.Close()

			encoder := gob.NewEncoder(file)
			if err := encoder.Encode(elements.Molds); err != nil {
				return err
			}
			return nil
		*/
		return nil
	})
	return err
}

func MoldPath(name string) string {
	var path string
	utils.DirWrapper(data_path, func(dir string) {
		path = filepath.Join(dir, "molds", name, elements.Molds[name].Filename)
	})
	return path
}

// TODO: Implement a way to use other molds and then save them
func NewMoldFromFile(path string) (elements.Mold, error) {
	/* path -> yaml file with:
	name: string
	filename: string
	args:
		serve: string
		port: string
	*/
	var mold elements.Mold = elements.Mold{}
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(content, &mold)
		if err != nil {
			return err
		}
		return nil
	})
	return mold, err
}

func ReadPIDS() error {
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		file, err := os.Open(filepath.Join(dir, pids_file))
		if err != nil {
			// If file doesn't exist, just return, it will be created on save
			return nil
		}
		defer file.Close()

		decoder := gob.NewDecoder(file)
		if err := decoder.Decode(&Pids); err != nil {
			return err
		}
		return nil
	})
	return err
}

func WritePIDS() error {
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		dirPath := filepath.Join(dir, pids_file)
		file, err := os.Create(dirPath)
		if err != nil {
			return err
		}
		defer file.Close()

		encoder := gob.NewEncoder(file)
		if err := encoder.Encode(Pids); err != nil {
			return err
		}
		return nil
	})
	return err
}

func ReadHelp() (string, error) {
	var help string
	err := utils.DirWrapperWithError(data_path, func(dir string) error {
		content, err := os.ReadFile(filepath.Join(dir, "commands.txt"))
		if err != nil {
			return err
		}
		help = string(content)
		return nil
	})
	return help, err
}
