package template

import (
	"fmt"
	"strings"

	e "github.com/sorinlg/stencil/err"
)

func Panic(s interface{}) (string, error) {
	if s != nil {
		return s.(string), nil
	}

	return "", fmt.Errorf("nil value encountered")
}

func Default(args ...interface{}) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("default called with no values!")
	}

	if len(args) > 0 {
		if args[0] != nil {
			return args[0].(string), nil
		}
	}

	if len(args) > 1 {
		return args[1].(string), nil
	}

	return "", fmt.Errorf("default called with no default value")
}

func GenListFromString(s string, sep string) []string {
	return strings.Split(s, sep)
}

func RemoveItemFromList(i string, l []string) []string {
	pos, err := getItemIndexFromList(l, i)
	e.Check(err)

	l = append(l[:pos], l[pos+1:]...)
	return l
}

func getItemIndexFromList(l []string, s string) (int, error) {
	for i, v := range l {
		if v == s {
			return i, nil
		}
	}
	return -1, fmt.Errorf("item %s cannot be found in list %s", s, l)

}
