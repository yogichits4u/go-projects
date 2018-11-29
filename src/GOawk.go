package main

import (
	"fmt"
	"os"

	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()
	s.Begin = func(s *awk.Script) {
		s.SetFS(",")
	}
	s.AppendStmt(nil, func(s *awk.Script) {
		fmt.Printf("%10v \t + %10v \t = %10v\n", s.F(7).Int(), s.F(8).Int(), s.F(7).Int()+s.F(8).Int())
	})
	s.Run(os.Stdin)
}
