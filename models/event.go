package models

import (
	"time"

	"rest.com/main/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (event Event) Save() error {
	// todo: add to db
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.Id = id

	if err != nil {
		return err
	}

	events = append(events, event)
	return nil
}

func GetAllEvents() []Event {
	return events
}
