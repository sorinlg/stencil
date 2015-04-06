package template

import (
	"strings"
)

func List(list string, delim string) string {
	slice := strings.Split(list, " ")
	result := ""
	for i, v := range slice {
		if i != 0 {
			result += delim
		}
		result += v
	}

	return result
}

func Pos(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1

}

func ListExclude(list string, item string, delim string) string {
	slice := strings.Split(list, " ")
	pos := Pos(slice, item)
	shortSlice := append(slice[:pos], slice[pos+1:]...)
	shortList := strings.Join(shortSlice, " ")

	result := List(shortList, delim)
	return result
}
