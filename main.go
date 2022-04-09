package main

import (
	"encoding/json"
	"fmt"
	"github.com/tufanbarisyildirim/gonginx/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need nginx file as argument")
		return
	}
	nginxFile := os.Args[1]
	p, err := parser.NewParser(nginxFile)
	if err != nil {
		fmt.Println("{\"error\": true}")
		return
	}

	conf := p.Parse()
	s, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		fmt.Println("{\"error\": true}")
		return
	}
	fmt.Println(string(s))

}
