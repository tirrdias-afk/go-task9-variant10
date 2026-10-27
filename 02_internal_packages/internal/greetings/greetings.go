package greetings

import "strings"

// Hello атаудың айналасындағы бос орындарды алып тастайды.
func Hello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "студент"
	}
	return "Сәлем, " + name + "!"
}
