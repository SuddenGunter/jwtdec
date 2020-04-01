package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/tidwall/pretty"
	"log"
	"strings"
)

func main() {
	noColor := flag.Bool("nc", false, "disable colorized output")
	flag.Parse()
	jwtToken := flag.Args()[0]
	payload := strings.Split(jwtToken, ".")[1]
	json, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		log.Fatal(err)
	}
	prettyJson := pretty.Pretty(json)
	if !*noColor {
		prettyJson = pretty.Color(prettyJson, nil)
	}
	fmt.Print(string(prettyJson))
}
