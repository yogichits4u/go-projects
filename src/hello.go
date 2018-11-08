package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	absPath, _ := filepath.Abs("data/sample.dat")
	dat, err := ioutil.ReadFile(absPath)
	//dat, err := ioutil.ReadFile("/home/yogeshc/go/src/github.com/go-projects/data/sample.dat")
	check(err)
	fmt.Print(string(dat))

	file, err := os.Open(absPath)
	check(err)
	defer file.Close()

	fmt.Println(file.Name())

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(scanner.Text()) > 15 {
			//fmt.Println(scanner.Text())
			result := strings.Split(scanner.Text(), ",")
			for i := range result {
				str := []string{result[i]}
				sort.Strings(str)
				fmt.Printf("%-10s \t", str)
			}
			fmt.Println()
		}

	}
}
