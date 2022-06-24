package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_ID int, _status string) *Door {
	return &Door{_ID, _status}
}
