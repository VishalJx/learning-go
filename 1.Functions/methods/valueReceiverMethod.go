package main

type Circle struct {
	Radius float64
}

//Method with a value receiver (does not modify the original object)
func (c Circle) Area() float64 {
	A := 3.14 * c.Radius * c.Radius

	return A
}

func ValueReceiverMethodExample() (float64, float64){
	circle := Circle{
		Radius: 5.0,
	}
	return circle.Radius, circle.Area()
}

//The receiver is passed by value, 
//meaning it operates on a copy of the object.
