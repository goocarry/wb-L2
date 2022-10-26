package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalln("usage: [flags] [pattern or string] [file]")
	}

	input, err := NewInput(args[len(args)-1])
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	findBy := args[len(args)-2:len(args)-1]
	input.FindBy = strings.Join(findBy, " ")

	flag.IntVar(&input.After, "A", 0, "print + N lines after match")
	flag.IntVar(&input.Before, "B", 0, "print + N lines before match")
	flag.IntVar(&input.Context, "C", 0, "print +-N lines around match")
	flag.BoolVar(&input.Count, "c", false, "show lines count")
	flag.BoolVar(&input.IgnoreCase, "i", false, "ignore letters case")
	flag.BoolVar(&input.Invert, "v", false, "exclude match instead")
	flag.BoolVar(&input.Fixed, "F", false, "find whole line match")
	flag.BoolVar(&input.LineNum, "n", false, "print match line number")
	flag.Parse()

	input.Grep()
	
}
