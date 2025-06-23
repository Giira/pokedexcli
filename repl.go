package main

import "strings"

func cleanInput(text string) []string {
	var output []string
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	tmp := strings.Split(text, " ")
	for _, word := range tmp {
		if word != " " && word != "" {
			output = append(output, word)
		}
	}
	return output
}
