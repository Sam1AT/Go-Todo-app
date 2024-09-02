package qtodo

import "errors"

type Database interface {
	GetTaskList() []Task
	GetTask(string) (Task, error)
	SaveTask(Task) error
	DelTask(string) error
}

type DataBaseInMemory struct {
	Tasks []Task
}

func NewDatabase() Database {

	return &DataBaseInMemory{}
}

func (db *DataBaseInMemory) GetTaskList() []Task {
	return db.Tasks
}

func (db *DataBaseInMemory) GetTask(taskName string) (Task, error) {
	tasks := db.GetTaskList()
	for _, task := range tasks {
		if task.GetName() == taskName {
			return task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (db *DataBaseInMemory) SaveTask(task Task) error {
	db.Tasks = append(db.Tasks, task)
	return nil
}

func (db *DataBaseInMemory) DelTask(taskName string) error {
	tasks := db.GetTaskList()
	for i, task := range tasks {
		if task.GetName() == taskName {
			db.Tasks = append(db.Tasks[0:i], db.Tasks[i:len(tasks)-1]...)
			return nil
		}
	}
	return errors.New("Task not found")
}
