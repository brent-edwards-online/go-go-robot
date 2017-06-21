package command

import (
	"testing"
	"enumtypes"
)

func TestParseCommand(t *testing.T){
	var tests = []struct {
		inputString string
		want string
		isvalid bool
		instruction enumtypes.InstructionType
	}{
		{"", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"RUBBISH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"MOVE", "MOVE", true, enumtypes.MOVE},
		{"LEFT", "LEFT", true, enumtypes.LEFT},
		{"RIGHT", "RIGHT", true, enumtypes.RIGHT},
		{"REPORT", "REPORT", true, enumtypes.REPORT},
		{"move", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"left", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"right", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"report", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"MOVE ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"LEFT ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"RIGHT ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"REPORT ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"MOVE RUBBISH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"LEFT RUBBISH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"RIGHT RUBBISH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"REPORT RUBBISH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"MOVE 10,20,NORTH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"LEFT 10,20,NORTH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"RIGHT 10,20,NORTH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 10", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 10,10", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 20,WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 10,20,NORTH", "PLACE 10,20,NORTH", true, enumtypes.PLACE},
		{"PLACE 10,20,NORTH ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE -10,20,SOUTH", "PLACE -10,20,SOUTH", true, enumtypes.PLACE},
		{"PLACE 10,-20,EAST", "PLACE 10,-20,EAST", true, enumtypes.PLACE},
		{"PLACE -10,-20,WEST", "PLACE -10,-20,WEST", true, enumtypes.PLACE},
		{"PLACE 10,20,NORTH,RUBBISH", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 0.5,1,WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 1,0.5,WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 1 ,1,WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 1, 1,WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 1,1 ,WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 1,1, WEST", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE 1,1,WEST ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
		{"PLACE1,1,WEST ", "INVALIDINSTRUCTION", false, enumtypes.INVALIDINSTRUCTION},
	}
	
	for _, testcase := range tests {
		var command Command
		command.Parse(testcase.inputString)
		
		isValid := command.IsValid()
		if isValid != testcase.isvalid {
			t.Errorf("Want: %v, Got: %v", testcase.isvalid, isValid)
		}
		got := command.ToString()
		if got != testcase.want {
			t.Errorf("Want: %v, Got: %v", testcase.want, got)
		}		
		instruction := command.GetInstruction()
		if instruction != testcase.instruction {
			t.Errorf("Want: %v, Got: %v", testcase.instruction, instruction)
		}
	}
}
