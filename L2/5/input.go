package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Input ...
type Input struct {
	Text []string
	FindBy string
	Params
}

// Params ...
type Params struct {
	// print + N lines after match
	After int 
	// print + N lines before match
	Before int
	// print +-N lines around match
	Context int
	// show lines count
	Count bool
	// ignore letters case
	IgnoreCase bool
	// exclude match instead
	Invert bool
	// find whole line match
	Fixed bool
	// print match line number
	LineNum bool
}

// NewInput ...
func NewInput(filePath string) (*Input, error) {
	input, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	text := []string{}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		text = append(text, scanner.Text())
	}

	input.Close()
	return &Input{
		Text: text,
	}, nil
}


// Grep ...
func (input *Input) Grep() {

	res := []string{}
	
	var found []int
	if input.Fixed {
		found = findByWholeString(input.Text, input.FindBy)
	} else {
		found = findByPattern(input.Text, input.FindBy)
	}

	for _, v  := range found {
		//TODO: use slicing for contexts and before/after
		res = append(res, input.Text[v])
	}


	fmt.Println(res)
}

func findByPattern(text []string, pattern string) []int {
	res := []int{}
	re, _ := regexp.Compile(pattern)
	for i := range text {
		if re.FindString(text[i]) !=  "" {
			res = append(res, i)
		}
	}

	return res
}


func findByWholeString(text []string, pattern string) []int {
	res := []int{}

	for i := range text {
		if text[i] == pattern {
			res = append(res, i)
		}
	}

	return res
}

