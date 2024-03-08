// 5.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Create new FlagSet
	command := flag.NewFlagSet("calc", flag.ContinueOnError)

	// Getting numbers and operator written after subcommand using flags
	// ex.: go run main.go calc -num1=2 -num2=3 -op=+
	num1 := command.Float64("num1", 0.0, "First Number")
	num2 := command.Float64("num2", 0.0, "Second Number")
	operator := command.String("op", "", "Operator defines operation")

	// Checking that passing subcommand or not
	if len(os.Args) < 2 {
		fmt.Println("Please provide sub-command")
		return
	}

	// Check that "calc" is passed as sub-command or not
	switch os.Args[1] {
	case "calc":
		// Parsing the arguments in terms of flags
		command.Parse(os.Args[2:])
		// Checking for the operator and perform operation according to that
		switch *operator {
		case "+":
			fmt.Println("Ans :", *num1+*num2)
		case "-":
			fmt.Println("Ans :", *num1-*num2)
		case "*":
			fmt.Println("Ans :", *num1**num2)
		case "//":
			fmt.Println("Ans :", int(*num1 / *num2))
		case "/":
			if *num2 == 0 {
				fmt.Println("hmmm! Don't try Zero.")
				break
			}
			fmt.Println("Ans :", *num1 / *num2)
		// Handling if not pass enough flags
		default:
			fmt.Println("Invalid sub-command, Seem like forgot to add flags!")
		}
	// Handling if not passed subcommand
	default:
		fmt.Println("Invalid sub-command, Try 'calc' Bro!")
	}
}

// 4.

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// )

// // GreetCommand struct
// type GreetCommand struct {
// 	flagSet *flag.FlagSet
// 	name    string
// }

// // Runner interface to bind methods to GreetCommand
// type Runner interface {
// 	Init([]string) error
// 	Run() error
// 	Name() string
// }

// // Initialize new instance of GreetCommand and return a pointer of it
// func NewGreetCommand() *GreetCommand {
// 	// Create new flagset for subcommand "greet"
// 	greetCommand := &GreetCommand{
// 		flagSet: flag.NewFlagSet("greet", flag.ContinueOnError),
// 	}

// 	// through name flag store it's value in instance's name
// 	greetCommand.flagSet.StringVar(&greetCommand.name, "name", "Default One! Put Your Name and Try", "Name of the Person to be Greeted")
// 	return greetCommand
// }

// func (gc *GreetCommand) Name() string {
// 	return gc.flagSet.Name()
// }

// func (gc *GreetCommand) Run() error {
// 	fmt.Println("Hello", gc.name, "\b!")
// 	return nil
// }

// func (gc *GreetCommand) Init(args []string) error {
// 	return gc.flagSet.Parse(args)
// }

// func root(args []string) error {
// 	// Checking that length of arguments should not 0
// 	if len(args) < 1 {
// 		return fmt.Errorf("You must pass a sub command")
// 	}
// 	commands := []Runner{
// 		NewGreetCommand(),
// 	}
// 	// take subcommand as first argument
// 	subCommand := os.Args[1]

// 	// Ranging over commands to match it with subcommand
// 	// if it match, parse it and return result
// 	for _, command := range commands {
// 		if command.Name() == subCommand {
// 			command.Init(os.Args[2:])
// 			return command.Run()
// 		}
// 	}

// 	return fmt.Errorf("Unknown subcommand: %s", subCommand)
// }

// func main() {
// 	// passing the subcommand to root function
// 	err := root(os.Args[1:])
// 	if err != nil {
// 		fmt.Println("ERROR :", err.Error())
// 		return
// 	}
// }

// 3.

// package main

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	var count int
// 	flag.IntVar(&count, "n", 3, "Number of Lines have to Read per File")
// 	// flag.Int("count", 0, "Number of Lines have to Read per File")

// 	// takes flag -n with int value and store it in count variable
// 	flag.Parse()

// 	// If filename is passed then read from it, otherwise read from console
// 	var in io.Reader
// 	fileName := flag.Arg(0)
// 	if fileName != "" {
// 		file, err := os.Open(fileName)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return
// 		}
// 		defer file.Close()
// 		in = file
// 	} else {
// 		in = os.Stdin
// 	}

// 	buf := bufio.NewScanner(in)
// 	for i := 0; i < count; i++ {
// 		fmt.Print("Write Here: ")
// 		if !buf.Scan() {
// 			break
// 		}
// 		fmt.Println("You Wrote: ", buf.Text())
// 	}

// 	err := buf.Err()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "error reading: err:", err)
// 	}
// }

// 2.

// package main

// import (
// 	"flag"
// 	"fmt"
// )

// type Color string

// const (
// 	ColorBlack  Color = "\u001b[30m"
// 	ColorRed          = "\u001b[31m"
// 	ColorGreen        = "\u001b[32m"
// 	ColorYellow       = "\u001b[33m"
// 	ColorBlue         = "\u001b[34m"
// 	ColorReset        = "\u001b[0m"
// )

// func colorize(color Color, message string) {
// 	// Set the color and reset the color after printing message in console, using escape sequence characters
// 	fmt.Println(string(color), message, string(ColorReset))
// }

// // Show colored output using escape sequence character if -color flag is passed
// func main() {
// 	useColor := flag.Bool("color", false, "display colorized output")
// 	flag.Parse()

// 	if *useColor {
// 		// pass the color you want to show for output
// 		colorize(ColorBlue, "Hello, World from Golang!")
// 		return
// 	}
// 	fmt.Println("Hello, World from Golang!")
// }

// 1.

// package main

// import (
// 	"flag"
// 	"fmt"
// )

// // Basic program that checks that -color flags is passed or not
// func main() {
// 	useColor := flag.Bool("color", false, "Display Color String")
// 	flag.Parse()

// 	if *useColor {
// 		fmt.Println("With Color Life is Good!")
// 		return
// 	}

// 	fmt.Println("Without Color Life is Bad")
// }
