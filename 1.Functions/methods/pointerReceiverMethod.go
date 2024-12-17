package main

type Rectangle struct {
	Length  float64
	Breadth float64
}

//Modifying the original value
func (r *Rectangle) Resize(newLength, newBreadth float64) (float64, float64) {
	r.Length = newLength
	r.Breadth = newBreadth

	return r.Length, r.Breadth
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func PointerReceiverMethodExample() (float64,float64,float64,float64,float64){
	rectangle := Rectangle{
		Length:  10,
		Breadth: 5,
	}

	ogLength := rectangle.Length
	ogBreadth := rectangle.Breadth

	modLength, modBreadth := rectangle.Resize(20, 10)

	return ogLength, ogBreadth, modLength, modBreadth, rectangle.Area()

}

//This allows the method to modify the original object
