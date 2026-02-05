package api

import (
	"encoding/json"
	"errors"
	"time"
)

type AddTaskDto struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

func (t *AddTaskDto) ValidateForAddTask() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

type ErrorDto struct {
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *ErrorDto) ToString() string {
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

type CompleteTaskDto struct {
	Completed *bool `json:"completed"`
}

func (c *CompleteTaskDto) ValidateForCompleteTask() error {
	if c.Completed == nil {
		return errors.New("completed is required")
	}
	return nil
}