package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MasterDimmy/go-cls"
	"github.com/fatih/color"
)

type Task struct {
	name    string
	toggled bool
}

func clearScreen() {
	cls.CLS()
}

func toggleTask(tasks []Task, index int) []Task {
	if index < 0 || index > len(tasks) {
		fmt.Println("Invalide Task number")
		return tasks
	}
	tasks[index].toggled = !tasks[index].toggled
	status := "incomplete" // <- status is an unuser variable
	if tasks[index].toggled {
		status = "complete"
	}
	fmt.Print("Task %s is now %s\n", tasks[index].name, status)
	return tasks
}

func main() {
	command := map[string]int8{
		"list":   1,
		"toggle": 2,
		"exit":   3,
		"clear":  4,
	}
	_ = command
	scanner := bufio.NewScanner(os.Stdin)
	var tasks []Task

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			break
		} else if input == "clear" {
			clearScreen()
		} else if strings.HasPrefix(input, "add ") {
			desc := strings.TrimSpace(input[4:])
			if desc == "" {
				fmt.Println("Usage: add <task description>")
			} else {
				tasks = append(tasks, Task{name: desc, toggled: false})
				fmt.Println("Task added:", desc)
			}
		} else if input == "list" || input == "ls" {
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for i, t := range tasks {
					status := color.YellowString("[ ]")
					if t.toggled {
						status = color.YellowString("[X]")
					}
					fmt.Printf("%d. %s %s\n", i+1, t.name, status)
				}
			}
		} else if strings.HasPrefix(input, "toggle") {
			args := strings.TrimSpace(input[6:]) // take everything after "toggle"
			if args == "" {
				fmt.Println("Usage: toggle <task number>")
				continue
			}

			index, err := strconv.Atoi(args)
			if err != nil {
				fmt.Println("Invalid number")
				continue
			}

			tasks = toggleTask(tasks, index-1)
		} else if input == "help" {
			fmt.Println("add <task description>   to add a task")
			fmt.Println("toggle <task index>      to mark the task")
			fmt.Println("exit                     to quite the app")
		} else {
			color.Green("use the help command use 'help'")
		}
	}
}
