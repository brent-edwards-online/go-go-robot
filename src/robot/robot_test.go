package robot

import (
	"testing"
	"enumtypes"
	"command"
)

func TestTurnLeft(t *testing.T){
	robot := Robot {
				X_max: 5,
				Y_max: 5,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = enumtypes.NORTH
	var got = robot.Direction
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Should still be NORTH since no PLACE command received yet
	robot.Left()
	got = robot.Direction
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// LEFT command should now function since valid place has been received
	robot.Place(0,0,enumtypes.NORTH)
	robot.Left()
	got = robot.Direction
	want = enumtypes.WEST
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Left()
	got = robot.Direction
	want = enumtypes.SOUTH
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Left()
	got = robot.Direction
	want = enumtypes.EAST
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Left()
	got = robot.Direction
	want = enumtypes.NORTH
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Left()
	got = robot.Direction
	want = enumtypes.WEST
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

}

func TestTurnRight(t *testing.T){
	robot := Robot {
				X_max: 5,
				Y_max: 5,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = enumtypes.NORTH
	var got = robot.Direction
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Should still be NORTH since no PLACE command received yet
	robot.Right()
	got = robot.Direction
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// LEFT command should now function since valid place has been received
	robot.Place(0,0,enumtypes.NORTH)
	robot.Right()
	got = robot.Direction
	want = enumtypes.EAST
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Right()
	got = robot.Direction
	want = enumtypes.SOUTH
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Right()
	got = robot.Direction
	want = enumtypes.WEST
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Right()
	got = robot.Direction
	want = enumtypes.NORTH
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Right()
	got = robot.Direction
	want = enumtypes.EAST
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

}

func TestMove(t *testing.T){
	robot := Robot {
				X_max: 3,
				Y_max: 3,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = -1
	var got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should not move if no PLACE is received yet
	robot.Move();
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Place(0,0,enumtypes.NORTH)
	want = 0
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should now respond to MOVE command
	robot.Move();
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should not move outside boundary of 3x3 table
	robot.Move();
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestMoveSouth(t *testing.T){
	robot := Robot {
				X_max: 3,
				Y_max: 3,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = -1
	var got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should not move if no PLACE is received yet
	robot.Move();
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Place(0,2,enumtypes.SOUTH)
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should now respond to MOVE command
	robot.Move();
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = 0
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should not move outside boundary of 3x3 table
	robot.Move();
	want = 0
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestMoveEast(t *testing.T){
	robot := Robot {
				X_max: 3,
				Y_max: 3,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = -1
	var got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = -1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Place(0,2,enumtypes.EAST)
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = 1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should not move outside boundary of 3x3 table
	robot.Move();
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestMoveWest(t *testing.T){
	robot := Robot {
				X_max: 3,
				Y_max: 3,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = -1
	var got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = -1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Place(2,1,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = 1
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	robot.Move();
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	// Robot should not move outside boundary of 3x3 table
	robot.Move();
	want = 0
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestPlace(t *testing.T){
	robot := Robot {
				X_max: 3,
				Y_max: 3,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	var want = -1
	var got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = -1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	var wantDirection = enumtypes.NORTH
	var gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(2,1,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.WEST
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(10,1,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.WEST
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(0,10,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 1
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.WEST
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(2,2,enumtypes.SOUTH)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.SOUTH
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(-2,1,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.SOUTH
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(1,-2,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.SOUTH
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}

	robot.Place(-1,-2,enumtypes.WEST)
	want = 2
	got = robot.X_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	want = 2
	got = robot.Y_position
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
	wantDirection = enumtypes.SOUTH
	gotDirection = robot.Direction
	if gotDirection != wantDirection {
		t.Errorf("Want: %v, Got: %v", wantDirection, gotDirection)
	}
}


func TestExecute(t *testing.T){
	robot := Robot {
				X_max: 3,
				Y_max: 3,
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	command := command.Command {}

	// Robot should ignore commands until PLACE command is received 
	command.Parse("MOVE")
	robot.Execute(command)
	got := robot.Report()
	want := "Output: -1,-1,NORTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("LEFT")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: -1,-1,NORTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("RIGHT")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: -1,-1,NORTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("RIGHT")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: -1,-1,NORTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("RUBBISH")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: -1,-1,NORTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	// Test commands are executed correctly
	command.Parse("PLACE 2,2,SOUTH")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: 2,2,SOUTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("RIGHT")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: 2,2,WEST"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("LEFT")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: 2,2,SOUTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("MOVE")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: 2,1,SOUTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("INVALID")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: 2,1,SOUTH"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }

	command.Parse("PLACE 0,0,EAST")
	robot.Execute(command)
	got = robot.Report()
	want = "Output: 0,0,EAST"
	if got != want { t.Errorf("Want: %v, Got: %v", want, got) }
}

func TestReport(t *testing.T){
	robot := Robot {
        X_position: -1,
				Y_position: -1,
        Direction: enumtypes.NORTH,
	}

	want := "Output: -1,-1,NORTH"

	got := robot.Report()
	if got != want {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
