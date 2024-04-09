package main

import (
	"fmt"

	"generate-client/client"
	"generate-client/client/operations"
)

func main() {
	params := &operations.GetTasksParams{}

	resp, err := client.Default.Operations.GetTasks(params)
	if err != nil {
		fmt.Printf("Что-то пошло не так: %s\n", err)
		return
	}

	taskTitles := make([]string, len(resp.Payload))
	for i, t := range resp.Payload {
		taskTitles[i] = *t.Title
	}
	fmt.Printf("Активные задачи: %v\n", taskTitles)
}
