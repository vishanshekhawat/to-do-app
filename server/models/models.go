package models

// ToDoList struct
type ToDoList struct {
	ID     int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string `json:"task,omitempty"`
	Status int    `json:"status,omitempty"`
}
