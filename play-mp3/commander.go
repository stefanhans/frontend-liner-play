package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	commands    = make(map[string]string)
	commandKeys []string
)

func commandsInit() {
	commands = make(map[string]string)

	// Developer
	commands["play"] = "play  \n\t for developer playing\n"

	// Internals
	commands["quit"] = "quit  \n\t close the session and exit\n"
	commands["usage"] = "usage  \n\t shows this usage with all command usages\n"

	// Add your own command to call a function interactively
	// commands["mycmd"] = "mycmd  \n\t does something\n"

	//
	commands["record"] = "record  \n\t records 5 sec\n"

	//
	commands["speak"] = "speak <english text>  \n\t text to speech and speak out\n"

	//
	commands["translate"] = "translate <german text>  \n\t translate to English and speak\n"

	// To store the keys in sorted order
	for commandKey := range commands {
		commandKeys = append(commandKeys, commandKey)
	}
	sort.Strings(commandKeys)
}

// Execute a command specified by the argument string
func executeCommand(commandline string) bool {

	// Trim prefix and split string by white spaces
	commandFields := strings.Fields(commandline)

	// Check for empty string without prefix
	if len(commandFields) > 0 {

		// Switch according to the first word and call appropriate function with the rest as arguments
		switch commandFields[0] {

		case "quit":
			quitCmdTool(commandFields[1:])
			return true

		case "play":
			play(commandFields[1:])
			return true

		// Link command and function
		//case "mycmd":
		//	myCmd(commandFields[1:])
		//	return true

		case "record":
			record(commandFields[1:])
			return true

		case "speak":
			speak(commandFields[1:])
			return true

		case "translate":
			translate(commandFields[1:])
			return true

		default:
			usage()
			return false
		}
	}
	return false
}

// Display the usage of all available commands
func usage() {
	for _, key := range commandKeys {
		fmt.Printf("%v\n", commands[key])
	}

}

func quitCmdTool(arguments []string) {

	// Get rid of warnings
	_ = arguments

	os.Exit(0)
}

// func myCmd(arguments []string) {}
