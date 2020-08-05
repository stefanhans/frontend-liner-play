package main

import (
	"fmt"
	"log"
	"sort"
)

type help struct {
	name        string
	group       string
	usage       string
	description string
	commands    map[string]*help
}

var (
	helpCommandKeys  []string
	groupKeys        []string
	groupCommandKeys = make(map[string][]string)

	mainHelp = &help{
		name:  "help",
		usage: "/help [ groups [<group>] | <command> ]",
		description: "help shows: \n" +
			"- this help text (/help)\n" +
			"- all groups of commands (/help groups)\n" +
			"- all commands of a group (/help groups <group>)\n" +
			"- help text of a command (/help <command>)",
		commands: map[string]*help{
			"shell": {
				name:        "shell",
				group:       "script",
				usage:       "/shell <script>",
				description: "shell executes the shell script",
			},
			"log": {
				name:        "log",
				group:       "cli",
				usage:       "/log (on <filename> | off)",
				description: "log starts or stops writing logging output in the specified file",
			},
			"exit": {
				name:        "exit",
				group:       "cli",
				usage:       "/exit",
				description: "exit exits the application directly",
			},
			"play": {
				name:        "play",
				group:       "development",
				usage:       "/play [...]",
				description: "play is for developer to play",
			},
			"init": {
				name:        "init",
				group:       "development",
				usage:       "/init",
				description: "init simulates the init process of the application",
			},
			"quit": {
				name:        "quit",
				group:       "development",
				usage:       "/quit",
				description: "quit simulates the exit process of the application",
			},
		},
	}
)

func helpInit() {

	// To store the commands in sorted order
	for _, command := range mainHelp.commands {
		helpCommandKeys = append(groupKeys, command.name)
	}
	sort.Strings(helpCommandKeys)

	// To store the groups in sorted order
	var exists bool
	for _, command := range mainHelp.commands {

		exists = false
		for _, group := range groupKeys {
			if group == command.group {
				exists = true
			}
		}
		if !exists {
			groupKeys = append(groupKeys, command.group)
		}
	}
	sort.Strings(groupKeys)

	for _, helpGroup := range groupKeys {
		var helpGroupCommands []string
		for _, helpGroupCommand := range mainHelp.commands {
			if helpGroupCommand.group == helpGroup {
				helpGroupCommands = append(helpGroupCommands, helpGroupCommand.name)
			}
		}
		sort.Strings(helpGroupCommands)

		groupCommandKeys[helpGroup] = helpGroupCommands
	}
}

//
func showHelp(arguments []string) {

	switch len(arguments) {
	case 0:
		displayHelp(mainHelp)
		break
	case 1:
		displayHelp(mainHelp, arguments[0])
		break
	case 2:
		displayHelp(mainHelp, arguments[0], arguments[1])
		break
	}
}

func displayHelp(help *help, cmd ...string) {

	log.Printf("displayHelp: %v\n", cmd)

	// Shows main help
	if len(cmd) == 0 {
		fmt.Printf("Usage: %s\nDescription: %s\n%s\n", help.usage, help.description,
			prompt)
		return
	}

	// commands related?
	if cmd[0] == "commands" {
		log.Printf("in commands\n")

		// Shows available commands
		var helpCommands string
		for _, k := range helpCommandKeys {
			helpCommands += fmt.Sprintf("%v ", k)
		}
		fmt.Printf("%s\n%s", helpCommands, prompt)
		return
	}

	// groups related?
	if cmd[0] == "groups" {
		log.Printf("in groups\n")

		// Shows available groups
		if len(cmd) == 1 {
			var helpGroups string
			for _, k := range groupKeys {
				helpGroups += fmt.Sprintf("%v ", k)
			}
			fmt.Printf("%s\n%s", helpGroups, prompt)
			return
		}

		// Shows commands for a group
		helpGroup := cmd[1]
		var groupCommands string
		for _, command := range groupCommandKeys[helpGroup] {

			groupCommands += fmt.Sprintf("%v ", command)
		}
		if groupCommands == "" {
			fmt.Printf("ERROR: %q is not a known group", cmd[0])
			return
		}
		fmt.Printf("%s\n%s", groupCommands, prompt)
		return
	}

	// Shows help for command
	if commandHelp, ok := help.commands[cmd[0]]; ok {
		fmt.Printf("Usage: %s\nDescription: %s\n%s", commandHelp.usage, commandHelp.description,
			prompt)
		return
	}
	fmt.Printf("ERROR: %q is not a known command", cmd[0])
	return

}
