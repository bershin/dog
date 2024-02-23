package dog

import "strings"

func Bark(bark string) string {
	return "Dog bark:" + strings.ToUpper(bark)
}

func Barks(barks string) string {
	return "Dogs bark:" + strings.ToUpper(barks)
}
