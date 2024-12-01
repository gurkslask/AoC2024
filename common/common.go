package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SumInts(ints []int) int {
	var result int
	for _, val := range ints {
		result += val
	}
	return result
}

func MinInt(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
func MaxInt(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func MinInts(ints []int) (int, int) {
	result := ints[0]
	rKey := 0
	for key, val := range ints {
		if val < result {
			result = val
			rKey = key
		}
	}
	return result, rKey
}
func MaxInts(ints []int) int {
	result := ints[0]
	for _, val := range ints {
		if val > result {
			result = val
		}
	}
	return result
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFileSlice(day string, test bool) []string {
	fmt.Println(os.Getwd())
	path := ""
	if test {
		path = fmt.Sprintf("../%v/test_input.txt", day)
	} else {
		path = fmt.Sprintf("../%v/input.txt", day)
	}

	file, err := os.ReadFile(path)
	CheckErr(err)

	s := strings.Split(string(file), "\n")

	return s[:len(s)-1]
}
func ContainsString(slice []string, i string) bool {
	res := false
	for _, num := range slice {
		if num == i {
			res = true
			break
		}
	}
	return res

}

func ContainsInt(slice []int, i int) bool {
	res := false
	for _, num := range slice {
		if num == i {
			res = true
			break
		}
	}
	return res

}

func StrToInt(s string) int {
	s = strings.Replace(s, " ", "", -1)
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return i
}
func StringsToInt(s []string) []int {
	res := []int{}
	for _, v := range s {
		if v == " " || v == "" {
			continue
		}
		res = append(res, StrToInt(v))
	}
	return res
}
func IntAbs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}
func StrToInt64(s string) int64 {
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return int64(i)
}

func ConvertSliceStringToInt(slice []string) []int {
	si := []int{}
	fmt.Printf("slice = %+v\n", slice)
	for _, num := range slice {
		nn, err := strconv.Atoi(num)
		CheckErr(err)
		si = append(si, nn)
	}
	return si
}
func ConvertSliceStringToInt64(slice []string) []int64 {
	si := []int64{}
	fmt.Printf("slice = %+v\n", slice)
	for _, num := range slice {
		nn, err := strconv.Atoi(num)
		CheckErr(err)
		si = append(si, int64(nn))
	}
	return si
}

func mapNumtextAndNums() map[string]int {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	return numbers

}
func mapAlphaWithNum() map[string]int {
	var A = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}
	return A
}

func CheckIfInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	} else {
		return true
	}
}

func MapContainsKey(indata map[int]int, value int) bool {
	if indata[value] == 0 {
		return true
	} else {
		return false
	}
}

func RemoveFromSliceInt(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
