package models

import "time"

type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (event Event) Save() {
	// todo: add to db
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
