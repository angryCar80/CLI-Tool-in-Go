package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/MasterDimmy/go-cls"
	"github.com/fatih/color"
)

type Task struct {
	Name    string `json:"name"`
	Toggled bool   `json:"toggled"`
}

func saveTask(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error marshaling the file")
		return
	}
	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("Error Writing file: ", err)
	}
}

func loadTasks() []Task {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{} // if no file yet, start empty
		}
		fmt.Println("Error reading file:", err)
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error unmarshaling tasks:", err)
		return []Task{}
	}

	return tasks
}

func clearScreen() {
	cls.CLS()
}

func toggleTask(tasks []Task, index int) []Task {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalide Task number")
		return tasks
	}
	tasks[index].Toggled = !tasks[index].Toggled
	status := "incomplete" // <- status is an unuser variable
	if tasks[index].Toggled {
		status = "complete"
	}
	fmt.Printf("Task %s is now %s\n", tasks[index].Name, status)
	return tasks
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := loadTasks()

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			break
		} else if input == "clear" {
			clearScreen()
		} else if input == "save" {
			saveTask(tasks)
			fmt.Println("You saved the tasks")
		} else if strings.HasPrefix(input, "add ") {
			desc := strings.TrimSpace(input[4:])
			if desc == "" {
				fmt.Println("Usage: add <task description>")
			} else {
				tasks = append(tasks, Task{Name: desc, Toggled: false})
				fmt.Println("Task added:", desc)
				saveTask(tasks)
			}
		} else if input == "list" || input == "ls" {
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for i, t := range tasks {
					status := color.YellowString("[ ]")
					if t.Toggled {
						status = color.YellowString("[X]")
					}
					fmt.Printf("%d. %s %s\n", i+1, t.Name, status)
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
