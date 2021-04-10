package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID int `json:"id"`
	Task string `json:"task"`
}

var (
	tasks []Task
)

func init () {
	tasks = []Task{Task{ID: 1, Task: "read book"}}
}


func readTasks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(tasks)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling the tasks array"`))
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)


}