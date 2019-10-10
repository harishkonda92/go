package structsType

type Employee struct { // exported struct name should start with Cap
	Name string // same as above
	Lastname string // exported field
	age int // unexported field
}