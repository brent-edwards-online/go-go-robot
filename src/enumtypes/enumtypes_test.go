package enumtypes

import "testing"

func TestDirectionToString(t *testing.T){
	var tests = []struct {
		direction DirectionType
		want string
	}{
		{INVALIDDIRECTION, "INVALIDDIRECTION"},
		{NORTH, "NORTH"},
		{EAST, "EAST"},
		{SOUTH, "SOUTH"},
		{WEST, "WEST"},
	}
	
	for _, testcase := range tests {
		got := DirectionToString(testcase.direction)
		if got != testcase.want {
			t.Errorf("Want: %v, Got: %v", testcase.want, got)
		}		
	}
}

func TestStringToDirection(t *testing.T){
	var tests = []struct {
		inputString string		
		want DirectionType
	}{
		{"RANDOM", INVALIDDIRECTION },
		{"NORTH", NORTH},
		{"EAST", EAST},
		{"SOUTH", SOUTH},
		{"WEST", WEST},
	}
	
	for _, testcase := range tests {
		got := StringToDirection(testcase.inputString)
		if got != testcase.want {
			t.Errorf("Want: %v, Got: %v", testcase.want, got)
		}		
	}
}

func TestInstructionToString(t *testing.T){
	var tests = []struct {
		instruction InstructionType
		want string
	}{
		{INVALIDINSTRUCTION, "INVALIDINSTRUCTION"},
		{MOVE, "MOVE"},
		{LEFT, "LEFT"},
		{RIGHT, "RIGHT"},
		{PLACE, "PLACE"},
		{REPORT, "REPORT"},
	}
	
	for _, testcase := range tests {
		got := InstructionToString(testcase.instruction)
		if got != testcase.want {
			t.Errorf("Want: %v, Got: %v", testcase.want, got)
		}		
	}
}

