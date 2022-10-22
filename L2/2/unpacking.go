package unpacking

import (
	"strconv"
)

// UnpackString ...
func UnpackString(s string) string {
	sym := []string{}
	for _, char := range s {
		sym = append(sym, string(char))
	}

	res := ""
	for i := 0; i < len(sym); i++ {

		if checkNumber(sym[i]) && checkNumber(sym[i+1]) {
			break
		}

		if i == len(sym)-1 {
			if !checkNumber(sym[i]) {
				res += sym[i]
				break
			}
		}

		
		if !checkNumber(sym[i]) && !checkNumber(sym[i+1]) {
			res += sym[i]
		}

		if !checkNumber(sym[i]) && checkNumber(sym[i+1]) {
			n, _ := strconv.Atoi(sym[i+1])
			for n > 0 {
				res += sym[i]
				n--
			}
		}

		
	}

	return res

}

func checkNumber(s string) bool {
	n, _ := strconv.Atoi(s)

	if n > 0 && n <= 9 {
		return true
	}

	return false
}
