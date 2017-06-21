package enumtypes

type DirectionType uint8

const (
    NORTH DirectionType = iota
    EAST
    SOUTH
    WEST
    INVALIDDIRECTION
)

func DirectionToString(d DirectionType)string{
    switch d {
        case NORTH:
            return "NORTH"
        case EAST:
            return "EAST"
        case SOUTH:
            return "SOUTH"
        case WEST:
            return "WEST"
        default:
            return "INVALIDDIRECTION"
	}
}

func StringToDirection(s string) DirectionType {
    switch s {
        case "NORTH":
            return NORTH
        case "EAST":
            return EAST
        case "SOUTH":
            return SOUTH
        case "WEST":
            return WEST
        default:
            return INVALIDDIRECTION
	}
}

type InstructionType uint8

const (
    PLACE InstructionType = iota
    MOVE
    LEFT
	RIGHT
	REPORT
	INVALIDINSTRUCTION
)

func InstructionToString(d InstructionType)string{
    switch d {
        case PLACE:
            return "PLACE"
        case MOVE:
            return "MOVE"
        case LEFT:
            return "LEFT"
        case RIGHT:
            return "RIGHT"
        case REPORT:
            return "REPORT"
        default:
            return "INVALIDINSTRUCTION"
	}
}

