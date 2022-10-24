package sort

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// File ...
type File struct {
	Lines [][]string
}

// ReadByLine ...
func (f *File) ReadByLine(filePath string) error {
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
        return err
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()
  
    for _, line := range fileLines {
        f.Lines = append(f.Lines, strings.Split(line, " "))
    }
     
    return nil
}


// WriteByLine ...
func (f *File) WriteByLine(filePath string) error {
    // 
    return nil
}
