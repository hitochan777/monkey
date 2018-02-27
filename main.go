package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/hitochan777/monkey/runner"
	// "github.com/nsf/termbox-go"
)

type Options struct {
	Filename string
}

func startRepl() {
	// err := termbox.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer termbox.Close()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	runner.Start(os.Stdin, os.Stdout, true)
}

func runScript(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open " + filename)
		return
	}
	runner.Start(bufio.NewReader(file), os.Stdout, false)
}

func cmdline() *Options {
	options := &Options{}
	flag.StringVar(&options.Filename, "filename", "", "filename of the script you want to run")
	flag.Parse()

	return options
}

func main() {
	options := cmdline()
	if len(options.Filename) > 0 {
		runScript(options.Filename)
	} else {
		startRepl()
	}
}
