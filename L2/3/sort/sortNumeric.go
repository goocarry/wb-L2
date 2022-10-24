package sort

import (
	"fmt"
	"sort"
	"strconv"
)

// TODO: refactor with one sortable using interface
type sortableInts struct {
    ints, idxs []int 
}


func (s sortableInts) Len() int           { return len(s.ints) }
func (s sortableInts) Less(i, j int) bool { return s.ints[i] < s.ints[j] }
func (s sortableInts) Swap(i, j int) {
    s.ints[i], s.ints[j] = s.ints[j], s.ints[i]
    s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}

// SortNumericColumn ...
func (f *File) SortNumericColumn(colNumber int, reversed bool) error {
    ints := []int{}
    idxs := []int{}
    for i, l := range f.Lines {
        if colNumber-1 > len(l) {
            return fmt.Errorf("no column %d found at line %d", colNumber, i)
        } 
        num, err := strconv.Atoi(l[colNumber])
        if err != nil {
            return fmt.Errorf("can't sort column %d by numeric value", colNumber)
        }
        ints = append(ints,  num) 
        idxs = append(idxs, i)
    }

    sortable := &sortableInts{
        ints: ints,
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

