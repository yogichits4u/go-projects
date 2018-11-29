package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func isOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func isModifiedToday(t time.Time) bool {
	return time.Now().Sub(t) < 24*time.Hour
}

func findFilesOlderThanOneDay(dir string) (files []os.FileInfo, err error) {
	tmpfiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range tmpfiles {
		if file.Mode().IsRegular() {
			if isOlderThanOneDay(file.ModTime()) {
				files = append(files, file)
			}
		}
	}
	return
}

func findFilesModifiedToday(dir string) (files []os.FileInfo, err error) {
	tmpfiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range tmpfiles {
		if file.Mode().IsRegular() {
			if isModifiedToday(file.ModTime()) {
				files = append(files, file)
			}
		}
	}
	return
}

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

func main() {
	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}

	err = os.Chdir("/ziguser2/zig/abp/zigwrk23/data")
	if err != nil {
		fmt.Println("Error in change dir")
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	fname, err := findFilesModifiedToday("/ziguser2/zig/abp/zigwrk23/data")
	if err != nil {
		fmt.Printf("unable to find files in findFilesModifiedToday: %v\n", err)
		return
	}

	for i := range fname {
		if strings.Contains(fname[i].Name(), "wqifh") {
			fmt.Println(fname[i].Name())
		}
	}

}
