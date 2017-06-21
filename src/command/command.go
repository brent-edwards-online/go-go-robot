package command

import (
    "enumtypes"
		"strconv"
		"strings"
)

type Command struct {
   isValid bool
   instruction enumtypes.InstructionType
	 placeX int
	 placeY int
	 placeDirection enumtypes.DirectionType
}

func ValidatePlaceInstruction(instruction string) (enumtypes.InstructionType, int, int, enumtypes.DirectionType, bool) {

	params := strings.Split(instruction, " ")
	if len(params) != 2 || params[0] != "PLACE" {
		return enumtypes.INVALIDINSTRUCTION, -1, -1, enumtypes.INVALIDDIRECTION, false
	}

	placeArguments := strings.Split(params[1], ",")
	if len(placeArguments) != 3 {
		return enumtypes.INVALIDINSTRUCTION, -1, -1, enumtypes.INVALIDDIRECTION, false
	}

	xpos, xerr := strconv.ParseInt(placeArguments[0],10,0)
	if xerr != nil {
		return enumtypes.INVALIDINSTRUCTION, -1, -1, enumtypes.INVALIDDIRECTION, false
	}
	ypos, yerr := strconv.ParseInt(placeArguments[1],10,0)
	if yerr != nil {
		return enumtypes.INVALIDINSTRUCTION, -1, -1, enumtypes.INVALIDDIRECTION, false
	}
	
	direction := enumtypes.StringToDirection(placeArguments[2])
	if direction == enumtypes.INVALIDDIRECTION {
		return enumtypes.INVALIDINSTRUCTION, -1, -1, enumtypes.INVALIDDIRECTION, false
	}

	return enumtypes.PLACE, int(xpos), int(ypos), direction, true
} 

func (c *Command) Parse(input string){
		c.isValid = true
    switch input {
        case "MOVE":
						c.instruction = enumtypes.MOVE
				case "LEFT":
						c.instruction = enumtypes.LEFT
				case "RIGHT":
						c.instruction = enumtypes.RIGHT
				case "REPORT":
						c.instruction = enumtypes.REPORT
        default:
						c.instruction, c.placeX, c.placeY, c.placeDirection, c.isValid  = ValidatePlaceInstruction(input)
	}
}

func (c *Command) IsValid() bool{
    return c.isValid
}

func (c *Command) GetInstruction() enumtypes.InstructionType{
    return c.instruction
}

func (c *Command) GetX() int{
	return c.placeX
}

func (c *Command) GetY() int{
	return c.placeY
}

func (c *Command) GetDirection() enumtypes.DirectionType{
	return c.placeDirection
}

func (c *Command) ToString() string{

		instructionString := enumtypes.InstructionToString(c.instruction)

    switch c.instruction {
        case enumtypes.PLACE:
						return instructionString + " " + strconv.Itoa(c.placeX) + "," + strconv.Itoa(c.placeY) + "," + enumtypes.DirectionToString(c.placeDirection)
        default:
						return instructionString
	}
}



