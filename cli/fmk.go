package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func PrintError(message, msgType string) {
	fmt.Printf("\n %s  %s: %s\n\n", "fmk", msgType, message)
	fmt.Printf("     %4s %-10s %-10s\n", "fmk", "{command}", "{subcommand}")
	fmt.Printf("     %4s %-10s %-10s\n", "fmk", "generate", "crud")
	fmt.Println()
}

type CrudData struct {
	ExportedResource string
	PrivateResource  string
	PackageName      string
}

func CreateCrud() error {
	data := CrudData{
		ExportedResource: "Restaurant",
		PrivateResource:  "restaurant",
		PackageName:      "restaurants",
	}

	log.Println(os.Getwd())

	t, err := template.ParseFiles("./cli/templates/endpoint.tmpl")
	if err != nil {
		panic(err)
	}

	// Create the file
	f, err := os.Create("restaurant.go")
	if err != nil {
		panic(err)
	}

	// Execute the template to the file.
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

	// Close the file when done.
	f.Close()

	return nil
}

func main() {

	filename := os.Args[0] // get command line first parameter

	filedirectory := filepath.Dir(filename)

	thepath, err := filepath.Abs(filedirectory)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(thepath)

	//if len(os.Args) != 3 { // fmk command subcommand
	//	PrintError("missing 2 arguments!", "HELP")
	//	return
	//}
	//
	//cmd := os.Args[1]
	//scmd := os.Args[2]
	//
	//switch {
	//case cmd == "generate" && scmd == "crud":
	//	CreateCrud()
	//default:
	//	PrintError("invalid arguments!", "ERROR")
	//}
}
