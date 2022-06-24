package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	callButtonID int
	status       string
	floor        int
	direction    string
}

func NewCallButton(_ID int, _status string, _floor int, _direction string) *CallButton {
	return &CallButton{_ID, _status, _floor, _direction}
}
