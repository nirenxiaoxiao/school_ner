package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	n "school_ner"
	"strings"
)

var mode = flag.String("mode", "suggest", "suggest/recognize")
var aliasPath = flag.String("alias_path", "school_alias.dict", "path to school alias")
var schoolPath = flag.String("school_path", "school_post.dict", "path to schools")

func main() {
	flag.Parse()
	if err := n.Initialize(*schoolPath, *aliasPath); err != nil {
		log.Fatalf("Error in initialize, %v", err)
	}

	log.Printf("mode = %s", *mode)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name := scanner.Text()
		fmt.Printf("======%s=======\n", name)
		var res []string
		if *mode == "suggest" {
			res = n.SuggestSchool(name)
		} else if *mode == "recognize" {
			res = n.RecognizeSchool(name)
		}

		fmt.Printf("%s\n", strings.Join(res, "\n"))
		fmt.Printf("===============\n")
	}
}
