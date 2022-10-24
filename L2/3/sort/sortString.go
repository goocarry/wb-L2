package sort

import (
	"fmt"
	"sort"
)

type sortableStrings struct {
    strings []string
    idxs []int
}

func (s sortableStrings) Len() int           { return len(s.strings) }
func (s sortableStrings) Less(i, j int) bool { return s.strings[i] < s.strings[j] }
func (s sortableStrings) Swap(i, j int) {
    s.strings[i], s.strings[j] = s.strings[j], s.strings[i]
    s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}


// SortStringColumn ...
func (f *File) SortStringColumn(colNumber int,  reversed bool) error {
    strings := []string{}
    idxs := []int{}
    for i, l := range f.Lines {
        if colNumber-1 > len(l) {
            return fmt.Errorf("no column %d found at line %d", colNumber, i)
        } 
        strings = append(strings, l[colNumber]) 
        idxs = append(idxs, i)
    }

    sortable := &sortableStrings{
        strings: strings,
        idxs: idxs,
    }
    
    sort.Sort(sortable)

    tmp := [][]string{}
	if !reversed {
		fmt.Println("reversed:", false)
		for i:=0; i <sortable.Len(); i++ {
			tmp = append(tmp, f.Lines[sortable.idxs[i]])
			fmt.Println(f.Lines[sortable.idxs[i]])
		}
	} else {
		fmt.Println("reversed:", true)
		for i:=sortable.Len()-1; i >= 0; i-- {
			tmp = append(tmp, f.Lines[sortable.idxs[i]])
			fmt.Println(f.Lines[sortable.idxs[i]])
		}
	}

    f.Lines = tmp

    return nil
}
