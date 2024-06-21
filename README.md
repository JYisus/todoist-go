# Todoist SDK for Go

> **WARNING:** This SDK is still in alpha phase. This is a part of a side project
of developing a Todoist CLI, so I may be doing breaking changes while I use it.

## Getting Started

### Installing

To install the SDK:

```bash
go get github.com/jyisus/go-todoist
```

To update the SDK:

```bash
go get -u github.com/jyisus/todoist-go
```

### Examples

```go
package main

import (
	"context"
	"log"

	"github.com/jyisus/todoist-go"
)

func main() {
	ctx := context.Background()
	c := todoist.NewClient("<your-todoist-token>")

	tasks, err := c.GetTasks(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting tasks: %s", err)
	}

	t, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Error marshalling tasks: %s", err)
	}

	task, err := c.GetTask(ctx, "<task-id>")
	if err != nil {
		log.Fatalf("Error getting task: %s", err)
	}

	t, err := json.Marshal(task)
	if err != nil {
		log.Fatalf("Error marshalling task: %s", err)
	}


	task, err := c.AddTask(ctx, &gotodoist.AddTaskOpts{
		Content:   "Decide if opts should be pointers or values",
		ProjectID: "<project-id>",
	})
	if err != nil {
		log.Fatalf("Error creating task: %s", err)
	}

	t, err := json.Marshal(task)
	if err != nil {
		log.Fatalf("Error marshalling task: %s", err)
	}

	if err := c.DeleteTask(ctx, "<task-id>"); err != nil {
		log.Fatalf("Error closing task: %s", err)
	}
}
```

