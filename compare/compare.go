package main

import (
    "fmt"
    "os"
	"path/filepath"
	s "strings"
)

func scanDir() (source []string, target []string) {
	var files []string
    root := "."
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        fmt.Println("an error occured")
	}

	// separate source and target from given array
	for _, file := range files {
		if s.Contains(file, "source\\") {
			trim_source := s.TrimPrefix(file, "source\\")
			source = append(source, trim_source)

		} else if s.Contains(file, "target\\"){
			trim_target := s.TrimPrefix(file, "target\\")
			target = append(target, trim_target)
		}
	}
	return
}

func compare(source []string, target []string) {
	for i := 0; i<len(source);i++ {
		for j:=i; j<len(target);i++ {
			if source[i] != target[j]{
				if s.Contains(source[i], target[j]){
					
				}
			}
		}
	}
}

func main() {
	source, target := scanDir()
	compare(source, target)
	fmt.Println("comparing this shit")
}