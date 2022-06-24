package main

import "sort"

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	bestScore             int
	direction             string
	currentFloor          int
	requestedDirection    string
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

func NewElevator(_ID int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	return &Elevator{_ID, _status, _amountOfFloors, 1, "on", _currentFloor, "up", *NewDoor(1, "closed"), []int{}, []int{}}
}

func (e *Elevator) move() {
	i := 0
	for i < len(e.floorRequestsList) {
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			e.sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()

		e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Slice(e.floorRequestsList, func(i, j int) bool {
			return e.floorRequestsList[i] < e.floorRequestsList[j]
		})
	} else {
		sort.Slice(e.floorRequestsList, func(i, j int) bool {
			return e.floorRequestsList[i] > e.floorRequestsList[j]
		})
	}
}

func (e *Elevator) operateDoors() {
	e.door.status = "opened"
}

func (e *Elevator) addNewRequest(requestedFloor int) {
	if contains(e.floorRequestsList, requestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}

	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}

	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}
