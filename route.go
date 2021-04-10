package main

import (
	"math/rand"
	"encoding/json"
	"net/http"
	"todolist/repository"
	"todolist/entities"
)

type Task struct {
	ID int `json:"id"`
	Task string `json:"task"`
}

var (
	repo repository.TaskRepository = repository.NewtaskRepo()
)


// func ReadTasks(resp http.ResponseWriter, req *http.Request) {
// 	resp.Header().Set("Content-type", "application/json")
// 	result, err := json.Marshal(tasks)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		resp.Write([]byte(`"error": "Error marshalling the tasks array"`))
// 		return
// 	}
// 	resp.WriteHeader(http.StatusOK)
// 	resp.Write(result)
// }

func WriteTasks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var task entities.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling the request"`))
		return
	}
	task.ID = rand.Int63()
	repo.Write(&task)
	resp.WriteHeader(http.StatusOK)

	json.NewEncoder(resp).Encode(task)
}

