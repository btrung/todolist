package repository

import (
	"context"
	"todolist/entities"
	"log"
	"cloud.google.com/go/firestore"

)

type TaskRepository interface {
	Read() ([]*entities.Task, error)
	Write(task *entities.Task) (*entities.Task, error)
}

type taskRepo struct{

}

func NewtaskRepo() TaskRepository{
	return &taskRepo{}
}

const (
	collectionName string = "task"
)



func (*taskRepo) Read() ([]*entities.Task, error) {

	return nil, nil
}


func (*taskRepo) Write(task *entities.Task) (*entities.Task, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "Todolist")
	if err != nil {
		log.Fatalf("fail to connect firestore: %v",err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": task.ID,
		"Task": task.Task,
	})
	if err != nil {
		log.Fatalf("fail adding a new task: %v", err)
		return nil, err
	}
	return task, nil
}