package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/222erik/todoapp/internal/tools"
)

var todoFileName = "todo.json"

var tasks = []tools.Task{}

func usageError() {
	fmt.Println("usage: 'todo -m <task> -p <priority> add', 'todo list', or 'todo done <priority of task>'")
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

	var err error
	tasks, err = tools.RestoreTodo(todoFileName)
	if err != nil {
		panic(err)
	}

	switch flag.Arg(0) {
	case "add":
		addTask(message, uint8(priority))
		err := tools.SaveTodo(tasks, todoFileName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Task added!")
	case "list":
		for _, v := range tasks {
			switch {
			case v.Priority < 10:
				fmt.Printf("  ")
			case v.Priority < 100:
				fmt.Printf(" ")
			}
			fmt.Printf("(%v) %v\n", v.Priority, v.TaskMessage)
		}
	case "done":
		fmt.Println("Done command doesn't work yet")
	default:
		usageError()
	}
}
