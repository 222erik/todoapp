package main

import (
	"flag"
	"fmt"
	"os"
)

type task struct {
	taskMessage string
	priority    uint8
}

var tasks = []task{}

func usageError() {
	fmt.Println("usage: 'todo add <task> <priority>', 'todo list', or 'todo done <task#>'")
	os.Exit(1)
}

func addTask(taskMessage string, priority uint8) {
	tasks = append(tasks, task{taskMessage, priority})

	var appendIndex int // Index where the task should be added in the list to make everything sorted by priority
	for i, v := range tasks {
		if v.priority > priority {
			appendIndex = i
		}
	}

	var tmp []task = make([]task, len(tasks)+1)

	// Insert the task in the right place
	tmp = append(tmp, tasks[:appendIndex]...)
	tmp = append(tmp, task{taskMessage, priority})
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
		fmt.Println("Added task!")
	case "list":
		for _, t := range tasks {
			fmt.Printf("(priority %d) %s\n", t.priority, t.taskMessage)
		}
	case "done":
		fmt.Println("Done command not implemented yet.")
	default:
		usageError()
	}
}
