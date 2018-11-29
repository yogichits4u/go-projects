package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	type Key struct {
		ES, SRC string
	}
	op := make(map[Key]int)

	absPath, _ := filepath.Abs(os.Args[1])
	// absPath, _ := filepath.Abs("TC_EVENT_LOG_RB_ES_RB1253_52042_IVER_180212662_20181108_010552_0.log")
	// dat, err := ioutil.ReadFile(absPath)
	// //dat, err := ioutil.ReadFile("/home/yogeshc/go/src/github.com/go-projects/data/sample.dat")
	// check(err)
	// fmt.Print(string(dat))

	file, err := os.Open(absPath)
	check(err)
	defer file.Close()

	fmt.Println(file.Name())

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(scanner.Text()) > 100 {
			//fmt.Println(scanner.Text())
			result := strings.Split(scanner.Text(), ",")
			// for i := range result {
			// 	fmt.Println(result[i])
			// }
			// fmt.Println(result[1])
			if result[1] == "RB" {
				op[Key{result[1], result[15]}]++

			}
		}

	}
	fmt.Println(scanner.Text())
	fmt.Println(op)
}
