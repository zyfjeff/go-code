package main

import (
	"fmt"
	"launchpad.net/gnuflag"
)

var name = gnuflag.String("name", "World", "A name to say hello to")
var spanish bool

func init() {
	gnuflag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	gnuflag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	gnuflag.Parse(true)
	/*
		gnuflag.PrintDefaults()

		gnuflag.VisitAll(func(flag *gnuflag.Flag) {
			format := "\t-%s: %s (Default: '%s')\n"
			fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
		})
	*/
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
