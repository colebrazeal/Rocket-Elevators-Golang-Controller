package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(_ID int, _status string, _floor int, _direction string) *FloorRequestButton {
	return &FloorRequestButton{_ID, _status, _floor, _direction}
}
