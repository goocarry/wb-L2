package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestFindAnagram(t *testing.T) {


	t.Run("find anagram", func(t *testing.T) {
		source := &[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
		res := &map[string][]string{
			"пятак":  {"пятка", "тяпка"},
			"листок": {"слиток", "столик"},
		}

		anagrams := FindAnagrams(source)

		fmt.Println(*anagrams)

		for k, v := range *anagrams {
			fmt.Println(k)
			fmt.Println(v)
			fmt.Println((*res)[k])
			if !reflect.DeepEqual(*v, (*res)[k]) {
				log.Fatalf("error")
			}
		}
		
	})
}