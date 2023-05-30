package main

import (
	"flag"
	"fmt"
	"strings"
)

const Version = "20230529"

var Args struct {
	help    bool
	version bool
}

func main() {
	flag.BoolVar(&Args.version, "version", false, "show the executable version")
	flag.BoolVar(&Args.help, "help", false, "display usage information")

	flag.Usage = func() {
		fmt.Println("Usage: golings [-v] [<command>] [<args>]")
		fmt.Println("\nGolings is a collection of small exercises to get you used to writing and")
		fmt.Println("reading Go Code")
		fmt.Println("\nOptions:")
		fmt.Println(strings.Join(flag.Args(), "\n"))
		flag.PrintDefaults()
	}

	flag.Parse()

	if Args.help {
		flag.Usage()
		return
	}

	if Args.version {
		fmt.Println(Version)
		return
	}

	fmt.Println(
		strings.Join(
			[]string{
				"",
				"Hello and",
				"	welcome to...			  ",
				"             _ _                  ",
				"  __ _  ___ | (_)_ __   __ _ ___  ",
				" / _' |/ _ \\| | | '_ \\ / _` / __| ",
				"| (_| | (_) | | | | | | (_| \\__ \\ ",
				" \\__, |\\___/|_|_|_| |_|\\__, |___/ ",
				" |___/                 |___/      ",
				"",
				"Thanks for installing Golings!",
				"",
				"Is this your first time? Don't worry, Golings was made for beginners! We are",
				"going to teach you a lot of things about Go, but before we can get started,",
				"here's a couple of notes about how Golings operates:",
				"",
				"1. The central concept behind Golings is that you solve exercises. These",
				"	exercises usually have some sort of syntax error in them, which will cause",
				"	them to fail compilation or testing. Sometimes there's a logic error instead",
				"	of a syntax error. No matter what error, it's your job to find it and fix it!",
				"	You'll know when you fixed it because then, the exercise will compile and",
				"	Golings will be able to move on to the next exercise.",
				"",
				"2. If you run Golings in watch mode (which we recommend), it'll automatically",
				"	start with the first exercise. Don't get confused by an error message popping",
				"	up as soon as you run Golings! This is part of the exercise that you're",
				"	supposed to solve, so open the exercise file in an editor and start your",
				"	detective work!",
				"",
				"3. If you're stuck on an exercise, there is a helpful hint you can view by",
				"	running `golings hint exercise_name`, or by using `hint` in watch mode.",
				"",
				"4. If an exercise doesn't make sense to you, feel free to open an issue on",
				"	GitHub! (https://github.com/latenitecode/golings/issues/new). We look at",
				"	every issue, and sometimes, other learners do too so you can help each",
				"	other out!",
				"",
				"Got all that? Great! To get started, run `golings watch` in order to get the",
				"first exercise. Make sure to have your editor open!",
			},
			"\n",
		),
	)
}
