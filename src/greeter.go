package src

import "log"

func Greet(name string) string {
	log.Print("Executing task greet with input variable: ", name)
	if name == "obi-wan" {
		return "GENERAL KENOBI!"
	} else {
		return "hello, " + name
	}
}
