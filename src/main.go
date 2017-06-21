package main

import (
    "fmt"
	"robot"
	"os"
	"bufio"
	"strings"
	"command"
	"enumtypes"
)

const XMAX int= 5
const YMAX int= 5

func main() {
    robot := robot.Robot {
		X_max: XMAX,
		Y_max: YMAX,
        X_position: -1,
		Y_position: -1,
        Direction: enumtypes.INVALIDDIRECTION,
	}
	keepGoing := true

	fmt.Println("Enter commands at the prompt:")
	fmt.Println("Enter 'x' to exit:")

	for keepGoing {
		reader := bufio.NewReader(os.Stdin)
    	inputCommand, _ := reader.ReadString('\n')
		inputCommand = strings.TrimRight(inputCommand, "\r\n")
		
		if inputCommand == "x" {
			keepGoing = false
		} else {
			command := command.Command {}
			command.Parse(inputCommand)
			if command.IsValid() {
				if command.GetInstruction() == enumtypes.REPORT && robot.IsPlaced {
					fmt.Println(robot.Report())
				} else {
					robot.Execute(command)
				}
			} 
		}
	}
}

