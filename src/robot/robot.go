package robot

import (
    "strconv"
    "bytes"
    "enumtypes"
    "command"
    "math"
)

type Robot struct {
   X_max int
   Y_max int
   X_position int
   Y_position int
   Direction  enumtypes.DirectionType
   IsPlaced bool
}

func (r *Robot) Report() string {
    var buffer bytes.Buffer
    buffer.WriteString("Output: ")
    buffer.WriteString(strconv.Itoa(r.X_position))
    buffer.WriteString(",")
    buffer.WriteString(strconv.Itoa(r.Y_position))
    buffer.WriteString(",")
    buffer.WriteString(enumtypes.DirectionToString(r.Direction))
    return buffer.String()
}

func (r *Robot) Move() {
    if r.IsPlaced {
        switch r.Direction {
            case enumtypes.NORTH:
                if r.Y_position < r.Y_max -1 {
                    r.Y_position++
                }
            case enumtypes.EAST:
                if r.X_position < r.X_max -1 {
                    r.X_position++
                }
            case enumtypes.SOUTH:
                if r.Y_position > 0 {
                    r.Y_position--
                }
            case enumtypes.WEST:
                if r.X_position > 0 {
                    r.X_position--
                }
        }
    }
}

func (r *Robot) Right() {
    if r.IsPlaced {
        newDirection := math.Mod(float64(r.Direction) + 1.0, 4.0)
        r.Direction = enumtypes.DirectionType(newDirection)
    }
}

func (r *Robot) Left() {
    if r.IsPlaced {
        newDirection := math.Mod(float64(r.Direction) - 1.0 + 4.0, 4.0)
        r.Direction = enumtypes.DirectionType(newDirection)
    }
}

func (r *Robot) Place(x int, y int, d enumtypes.DirectionType) {
    if(x < 0 || x >= r.X_max || y < 0 || y >= r.Y_max){
        return;
    }
    r.IsPlaced = true
    r. X_position = x
    r.Y_position = y
    r.Direction = d
}

func (r *Robot) Execute(command command.Command) {
    if command.IsValid() {
        switch command.GetInstruction() {
            case enumtypes.MOVE:
                r.Move();
            case enumtypes.LEFT:
                r.Left();
            case enumtypes.RIGHT:
                r.Right();
            case enumtypes.PLACE:
                r.Place(command.GetX(), command.GetY(), command.GetDirection())
        }
    }
}
