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


func ReadTasks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(tasks)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling the tasks array"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func WriteTasks(resp http.ResponseWriter, req *http.Request) {
	var task Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling the request"`))
		return
	}
	task.ID = len(tasks) +1
	tasks = append(tasks, task)
	resp.WriteHeader(http.StatusOK)

	result, err := json.Marshal(task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling the task array"`))
		return
	}
	resp.Write(result)
}

