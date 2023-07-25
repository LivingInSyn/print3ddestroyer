package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getMatchingFolders(basepath string) []string {
	matching := []string{}
	entries, err := os.ReadDir(basepath)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if strings.HasPrefix(e.Name(), "Microsoft.Print3D") {
			matching = append(matching, e.Name())
		}
	}
	return matching
}

func deletePrint3d(path string) {
	delpath := fmt.Sprintf("%s\\%s", path, "Print3D.exe")
	if _, err := os.Stat(delpath); err == nil {
		e := os.Remove(delpath)
		if e == nil {
			log.Printf("deleted Print3d at path: %s", path)
		} else {
			log.Printf("Error deleting path: %s. Error: %s", path, e.Error())
		}
	} else {
		log.Printf("No Print3D at path: %s\n", path)
	}
}

func main() {
	basepath := "c:\\Program Files\\WindowsApps"

	dirs := getMatchingFolders(basepath)

	for _, dir := range dirs {
		path := fmt.Sprintf("%s\\%s", basepath, dir)
		deletePrint3d(path)
	}
}
