package main

type BestElevatorInformations struct {
	bestScore    int
	referenceGap int
	bestElevator Elevator
}

func BestElevatorInfo(bestElevator Elevator) *BestElevatorInformations {
	return &BestElevatorInformations{6, 1000000, bestElevator}
}
