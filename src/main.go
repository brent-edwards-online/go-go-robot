package main

import (
    "fmt"
		"robot/robot"
)

func main() {
    robert := Robot {
        x_position: -1,
				y_position: -1,
        direction: NORTH,
		}

		fmt.Print(robert.Report())
}