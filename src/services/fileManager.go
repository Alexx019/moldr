package services

import (
	"encoding/gob"
	"io"
	"moldr/src/elements"
	"moldr/src/utils"
	"os"
	"path/filepath"
)

const (
	data_path   = ".data"
	pids_file   = "pids.bin"
	ingots_file = "ingots.bin"
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

func NewIngotFolder(name string) error {
	err := utils.DirWrapperWithError(name, func(dir string) error {
		if err := os.Mkdir(dir, 0755); err != nil {
			return err
		}
		if err := os.Mkdir(filepath.Join(dir, "logs"), 0755); err != nil {
			return err
		}
		if err := os.Mkdir(filepath.Join(dir, "data"), 0755); err != nil {
			return err
		}
		if _, err := os.Create(filepath.Join(dir, "logs", "log.txt")); err != nil {
			return err
		}
		// Copy the pocketbase executable into the ingot's data directory
		srcPath := filepath.Join("./pocketbase", "pocketbase.exe")
		dstPath := filepath.Join(dir, "data", "pocketbase.exe")

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
