package main

import "math"

type Battery struct {
	ID                        int
	amountOfColumns           int
	amountOfBasements         int
	amountOfFloors            int
	amountOfElevatorPerColumn int
	status                    string
	columnsList               []Column
	floorRequestsButtonsList  []FloorRequestButton
}

var columnID int = 1
var floorRequestButtonID = 1

func NewBattery(_ID, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	return &Battery{_ID, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn, "online", *createColumns(_amountOfColumns, _amountOfBasements, _amountOfFloors, _amountOfElevatorPerColumn), *createFloorRequestButtons(_amountOfFloors)}
}

//---------------------------------Methods--------------------------------------------//

func createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn int) {

	var servedFloors []int
	floor := -1

	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
}

func createColumns(_amountOfColumns, _amountOfBasements, _amountOfFloors, _amountOfElevatorPerColumn int) *[]Column {
	var columnsList []Column

	if _amountOfBasements > 0 {
		createBasementFloorRequestButtons(_amountOfBasements)
		createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}

	amountOfFloorsPerColumn := int(math.Ceil((float64(_amountOfFloors) / float64(_amountOfColumns))))
	floor := 1

	for i := 0; i < _amountOfColumns; i++ {
		servedFloors := []int{}
		for j := 0; j < amountOfFloorsPerColumn; j++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}

		column := *NewColumn(columnID, _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
		columnsList = append(columnsList, column)
		columnID++
	}
	return &columnsList
}

func createFloorRequestButtons(_amountOfFloors int) *[]FloorRequestButton {
	var floorRequestsButtonsList []FloorRequestButton
	buttonFloor := 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := *NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "up")
		floorRequestsButtonsList = append(floorRequestsButtonsList, floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
	return &floorRequestsButtonsList
}

func createBasementFloorRequestButtons(_amountOfBasements int) *[]FloorRequestButton {
	var floorRequestsButtonsList []FloorRequestButton
	buttonFloor := -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := *NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "down")
		floorRequestsButtonsList = append(floorRequestsButtonsList, floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
	return &floorRequestsButtonsList
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	// var chosenColumn *Column
	chosenColumn := b.columnsList[0]
	for _, column := range b.columnsList {
		for _, x := range column.servedFloors {
			if x == _requestedFloor {
				// chosenColumn := *column
				// break
				return &column
			}
		}
	}
	// for i := 0; i < len(b.columnsList); i++ {
	// 	if contains(b.columnsList[i].servedFloors, _requestedFloor) {
	// 		chosenColumn = &b.columnsList[i]
	// 		break
	// 	}
	// }
	return &chosenColumn
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

	column := *b.findBestColumn(_requestedFloor)
	elevator := *column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()

	elevator.addNewRequest(_requestedFloor)
	elevator.move()

	return &column, &elevator
}
