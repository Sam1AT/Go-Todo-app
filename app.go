package qtodo

import (
	"time"
)

type App interface {
	StartTask(string) error
	StopTask(string)
	AddTask(string, string, time.Time, func(), bool) error
	DelTask(string) error
	GetTaskList() []Task
	GetActiveTaskList() []Task
	GetTask(string) (Task, error)
}

type App1 struct {
	Database       Database
	AlertChannels  map[string]chan string
	DeleteAfterRun map[string]bool
}

func NewApp(db Database) App {
	return &App1{Database: db, AlertChannels: make(map[string]chan string), DeleteAfterRun: make(map[string]bool)}
}

func (app *App1) StartTask(name string) error {
	t, err := app.Database.GetTask(name)
	app.AlertChannels[name] = make(chan string)
	if err != nil {
		return err
	}

	go func() {
		for time.Now().Before(t.GetAlarmTime()) {

		}
		select {
		case <-app.AlertChannels[name]:
			return
		default:
			t.GetAction()()
			if app.DeleteAfterRun[name] {
				app.DelTask(name)
			}
		}
	}()
	return nil
}

func (app *App1) StopTask(name string) {
	if stopChan, exists := app.AlertChannels[name]; exists {
		close(stopChan)
		delete(app.AlertChannels, name)
	}
}

func (app *App1) AddTask(name string, description string, alarmTime time.Time, callback func(), enabled bool) error {
	err := app.Database.SaveTask(&T{Name: name, Description: description, Action: callback, AlarmTime: alarmTime})
	if err != nil {
		return err
	}
	app.DeleteAfterRun[name] = enabled
	return nil

}

func (app *App1) DelTask(name string) error {

	return app.Database.DelTask(name)
}

func (app *App1) GetTaskList() []Task {
	return app.Database.GetTaskList()
}

func (app *App1) GetActiveTaskList() []Task {
	list := []Task{}
	for _, task := range app.GetTaskList() {
		if _, exists := app.AlertChannels[task.GetName()]; exists {
			list = append(list, task)
		}
	}
	return list
}

func (app *App1) GetTask(name string) (Task, error) {
	task, err := app.Database.GetTask(name)
	if err != nil {
		return nil, err
	}
	return task, nil

}
