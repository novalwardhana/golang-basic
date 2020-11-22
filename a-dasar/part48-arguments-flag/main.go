package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 48 Arguments Flag")
	fmt.Println("")

	fmt.Println("Part 48.1 Penggunaan arguments")
	// go run part48-arguments-flag/*.go sate bakso "soto ayam"
	var argsRaw = os.Args
	fmt.Printf("ArgsRaw: %#v\n", argsRaw)
	fmt.Printf("argasRaw[0]: %#v\n", argsRaw[0])
	fmt.Printf("argasRaw[1]: %#v\n", argsRaw[1])
	fmt.Printf("argsRaw[2:]: %#v\n", argsRaw[2:])
	fmt.Println("")

	fmt.Println("Part 48.2 Penggunaan flag")
	var flagName = flag.String("name", "Novalita Kusuma Wardhana", "Your full name")
	var flagAge = flag.Int64("age", 25, "Insert your age")
	var flagBool = flag.Bool("active", true, "Set active")
	// go run part48-arguments-flag/*.go -name="wardhana"
	flag.Parse()
	fmt.Printf("flagName: %#v\n", *flagName)
	fmt.Printf("flagAge: %#v\n", *flagAge)
	fmt.Printf("flagBool: %#v\n", *flagBool)
	fmt.Println("")

	fmt.Println("Part 48.3 Deklarasi flag dengan cara passing reference")
	var flagNameReference string
	flag.StringVar(&flagNameReference, "name", "Novalita Kusuma Wardhana", "Your full name")
	var flagAgeReference int
	flag.IntVar(&flagAgeReference, "age", 25, "Insert your age")
	var flagBoolReference bool
	flag.BoolVar(&flagBoolReference, "active", true, "Set active")
	flag.Parse()
	fmt.Printf("flagNameReference: %#v\n", flagNameReference)
	fmt.Printf("flagAgeReference: %#v\n", flagAgeReference)
	fmt.Printf("flagBoolReference: %#v\n", flagBoolReference)
}
