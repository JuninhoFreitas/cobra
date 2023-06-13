package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Task struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func ReadTasksFromFile(filename string) ([]Task, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		// Se o arquivo não existir, criar um novo arquivo com valor padrão
		if os.IsNotExist(err) {
			tasks := []Task{{Name: "Default Task", Status: false}}
			err := SaveTasksToFile(filename, tasks)
			if err != nil {
				return nil, err
			}
			return tasks, nil
		}
		return nil, err
	}

	var taskList TaskList
	err = json.Unmarshal(file, &taskList)
	if err != nil {
		return nil, err
	}

	return taskList.Tasks, nil
}

func SaveTasksToFile(filename string, tasks []Task) error {
	taskList := TaskList{Tasks: tasks}

	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddOrUpdateTask(filename string, taskName string, taskStatus bool) error {
	tasks, err := ReadTasksFromFile(filename)
	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.Name == taskName {
			tasks[i].Status = taskStatus
			found = true
			break
		}
	}

	if !found {
		tasks = append(tasks, Task{Name: taskName, Status: taskStatus})
	}

	err = SaveTasksToFile(filename, tasks)
	if err != nil {
		return err
	}

	return nil
}
