// VCScheck 2018
// Authored by salmansamie

package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"path"
	"os/exec"
)


func main(){
	for _,elem := range rootFiles() {
		fmt.Printf("Repository: %s Directory: %s\n\n", isGitRepo(elem), elem)
	}
}


func isGitRepo(runExecPath string) string {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = runExecPath
	out, _ := cmd.CombinedOutput()
	return string(out)
}


func rootFiles() []string {
	itemsMeta, err := ioutil.ReadDir(getRoot())

	if err != nil {
		log.Fatalf("ioutil.ReadDir(getRoot()) failed with %s\n", err)
	}

	var storedDir []string
	for _, fName := range itemsMeta{
		setPath := path.Join(getRoot(), fName.Name())
		boolverd, _ := isDir(setPath)
		if  boolverd == true {
			storedDir = append(storedDir, setPath)
		}
	}
	return storedDir
}


func isDir(pathStr string) (bool, error) {
	fname, err := os.Stat(pathStr)

	if err != nil{
		log.Fatalf("os.Stat(pathStr) failed with %s\n", err)
	}
	return fname.Mode().IsDir(), nil
}


func getRoot() string {
	srcRoot, retErr := os.Getwd()

	if retErr != nil {
		log.Fatalf("os.Getwd() failed with %s\n", retErr)
	}
	return srcRoot
}
