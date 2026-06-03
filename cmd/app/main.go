package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/222erik/todoapp/internal/tools"
)

var tasks = []tools.Task{}

func usageError() {
	fmt.Println("usage: 'todo add <task> <priority>', 'todo list', or 'todo done <task#>'")
	os.Exit(1)
}

func addTask(taskMessage string, priority uint8) {
	var appendIndex = len(tasks)
	for i, v := range tasks {
		if v.Priority > priority {
			appendIndex = i
			break
		}
	}

	var tmp []tools.Task
	// Insert the task in the right place
	tmp = append(tmp, tasks[:appendIndex]...)
	tmp = append(tmp, tools.Task{TaskMessage: taskMessage, Priority: priority})
	tmp = append(tmp, tasks[appendIndex:]...)
	tasks = tmp
}

func main() {
	var message string
	var priority int
	flag.StringVar(&message, "m", "", "task message")
	flag.IntVar(&priority, "p", 1, "priority")
	flag.Parse()

	if flag.NArg() == 0 {
		usageError()
	}

	switch flag.Arg(0) {
	case "add":
		addTask(message, uint8(priority))
		tools.SaveTodo(tasks, "todo.json")
		fmt.Println("Added task!")
	case "list":
		for _, t := range tasks {
			fmt.Printf("(priority %d) %s\n", t.Priority, t.TaskMessage)
		}
	case "done":
		fmt.Println("Done command not implemented yet.")
	default:
		usageError()
	}
}
