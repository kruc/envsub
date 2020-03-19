package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func prepareFilepaths(args []string) []string {

	allPaths := []string{}
	for _, arg := range args {

		matches, err := filepath.Glob(arg)

		if err != nil {
			fmt.Println(err)
		}

		allPaths = append(allPaths, matches...)
	}

	return allPaths
}

func parseFile(path string, values values) {
	fmt.Printf("Start parsing %v...", path)
	if fileExists(path) {

		tmpl, err := template.ParseFiles(path)

		if err != nil {
			panic(err)
		}

		file, err := os.Create(strings.Replace(path, ".tmpl", "", -1))
		defer file.Close()
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(file, values)
		if err != nil {
			panic(fmt.Sprintf("Parse error: %v", err))
		}
		fmt.Print("Success!\n")
	} else {
		fmt.Printf("File %s doesn't exist\n", path)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
