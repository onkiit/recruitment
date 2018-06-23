package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"crypto/md5"
	s "strings"
	"path/filepath"
)

func scanDir() (source []string, target []string) {
	var files []string
	root := "."
	
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		converted_path := filepath.ToSlash(path) //change \ to /
        files = append(files, converted_path)
        return nil
    })
    if err != nil {
        log.Println("an error occured")
	}

	// separate source and target from given array
	for _, file := range files {
		if s.Contains(file, "source/") {
			trim_source := s.TrimPrefix(file, "source/")
			source = append(source, trim_source)

		} else if s.Contains(file, "target/"){
			trim_target := s.TrimPrefix(file, "target/")
			target = append(target, trim_target)
		}
	}
	return
}

func hashFile(path string) (hashed []byte){
	file, err := os.Open(path)
	if err != nil{
		return
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}
	hashed = hash.Sum(nil)

	return
}

func compareFile(c string, ns []string) bool{
	for _, n := range ns {
		if c == n {
			return false
		}
	}
	return true
}

func compareContent(s string, t string){
	src_root := "source/"
	tgt_root := "target/"
	if s == t {
		hs := hashFile(src_root + s)
		ts := hashFile(tgt_root + t)
		if string(hs) != string(ts) {
			fmt.Printf("%s => MODIFIED\n", s)
		}
	}
}

func main() {
	sources, targets := scanDir()

	for _, src := range sources {
		res := compareFile(src, targets)
		if(res){
			fmt.Printf("%s => NEW\n", src)
		}
	}

	for _, tgt := range targets {
		res := compareFile(tgt, sources)
		if(res){
			fmt.Printf("%s => DELETED\n", tgt)
		}
	}
	
	for _, src := range sources {
		for _, tgt := range targets {
			compareContent(src, tgt)
		}	
	} 
}