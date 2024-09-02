package qtodo

import (
	"errors"
	"time"
)

type Task interface {
	DoAction()
	GetAlarmTime() time.Time
	GetAction() func()
	GetName() string
	GetDescription() string
}

type T struct {
	Name        string
	Description string
	Action      func()
	AlarmTime   time.Time
}

func NewTask(action func(), Alarm time.Time, name string, description string) (Task, error) {
	if name == "" || description == "" {
		return nil, errors.New("task name or description is empty")
	}
	if time.Now().After(Alarm) {
		return nil, errors.New("alarm time is out of date")
	}
	addr := &T{Name: name, Description: description, Action: action, AlarmTime: Alarm}

	return addr, nil
}

func (t *T) DoAction() {
	t.Action()
}

func (t *T) GetName() string {
	return t.Name
}

func (t *T) GetDescription() string {
	return t.Description
}
func (t *T) GetAlarmTime() time.Time {
	return t.AlarmTime
}
func (t *T) GetAction() func() {
	return t.Action
}
