package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	commands    = make(map[string]string)
	commandKeys []string

	tmpDebugfile *os.File
)

func commandsInit() {
	commands = make(map[string]string)

	// Internals
	commands["quit"] = "quit  \n\t close the session and exit\n"

	// Developer
	commands["play"] = "play  \n\t for developer playing\n"

	for key, command := range mainHelp.commands {
		commands[key] = command.description
	}

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

		case "help":
			showHelp(commandFields[1:])
			return true

		case "execute":
			executeScript(commandFields[1:])
			return true

		case "sleep":
			sleepScript(commandFields[1:])
			return true

		case "echo":
			echoScript(commandFields[1:])
			return true

		case "ping":
			//ping(commandFields[1:])
			return true

		case "log":
			//cmdLogging(commandFields[1:])
			return true

		case "quit":
			quitCmdTool(commandFields[1:])
			return true

		case "play":
			play(commandFields[1:])
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

//func sendMessage(arguments []string) {
//
//	for k, v := range chatMembers {
//
//		// Do only send to others, not to yourself
//		if k != conf.Name {
//
//			// create TCP connection to recipient
//			conn, err := net.Dial("tcp", v.Sender)
//			if err != nil {
//				fmt.Printf("could not dial to %v: %v\n", v.Sender, err)
//				return
//			}
//
//			// send message
//			fmt.Fprintf(conn, "%s%s\n", prompt(), strings.Join(arguments, " "))
//
//			// close connection
//			conn.Close()
//		}
//	}
//}

func scriptPrompt(scriptname string) string {
	return fmt.Sprintf("<%s %q> ", time.Now().Format("Jan 2 15:04:05.000"), scriptname)
}

func executeScript(arguments []string) {

	if len(arguments) == 0 {
		fmt.Printf("error: no filename to execute specified\n")
		return
	}

	b, err := ioutil.ReadFile(arguments[0])
	if err != nil {
		fmt.Printf("ioutil.ReadFile: %v\n", err)
		return
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		//fmt.Printf("EXECUTE %d: %q\n", i, line)
		if strings.TrimSpace(line) == "" ||
			strings.Split(strings.TrimSpace(line), "")[0] == "#" {
			continue
		}
		echoScript(strings.Split(scriptPrompt(arguments[0])+line, " "))
		if _, ok := commands[strings.Split(line, " ")[0]]; ok {
			executeCommand(line)
		} else {
			fmt.Printf("error: %q is an unknown command\n", strings.Split(line, " ")[0])
		}
	}

}

func sleepScript(arguments []string) {

	var numSeconds int

	if len(arguments) == 0 {
		numSeconds = 1
	} else {
		numSeconds, err = strconv.Atoi(arguments[0])
	}

	time.Sleep(time.Second * time.Duration(numSeconds))
}

func echoScript(arguments []string) {

	fmt.Printf("%s\n", strings.Join(arguments, " "))
}

//func ping(arguments []string) {
//
//	if len(arguments) > 0 &&
//		arguments[0] != chatSelf.Name &&
//		chatMembers[arguments[0]] != nil {
//
//		pingMember(chatMembers[arguments[0]])
//	}
//}

//func cmdLogging(arguments []string) {
//
//	if len(arguments) == 0 ||
//		(len(arguments) == 1 && arguments[0] != "off") {
//		fmt.Printf("Error: wrong input. Usage: \n\t 'log (on <filename>) | off\n")
//
//		return
//	}
//
//	if arguments[0] == "on" && len(arguments) > 1 {
//		log.Printf("Switch to logging by command to %q\n", arguments[1])
//		tmpDebugfile, err = startLogging(arguments[1])
//		if err != nil {
//			fmt.Printf("Error: startLogging: %v\n", err)
//		} else {
//			log.Printf("Start logging by command to %q\n", arguments[1])
//		}
//
//		return
//	}
//
//	if arguments[0] == "off" {
//		log.Printf("Stop logging by command")
//		_ = tmpDebugfile.Close()
//
//		// Start debugging to file, if switched on or filename specified
//		if *debug || len(*debugfilename) > 0 {
//
//			_, err := startLogging(*debugfilename)
//			if err != nil {
//				fmt.Printf("could not start logging: %v\n", err)
//				return
//			}
//			log.Printf("Switch from logging by command to %q\n", debugfilename)
//		}
//	}
//}
