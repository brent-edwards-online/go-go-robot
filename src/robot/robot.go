package robot

type DirectionType uint8

const (
    NORTH DirectionType = iota + 1
    EAST
    SOUTH
    WEST
)

type Robot struct {
   x_position int
   y_position int
   direction  DirectionType
}

func (r *Robot) Report() string{
    return "Hi, my name is Robt"
}
