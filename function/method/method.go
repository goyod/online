package method

type Rectangle struct {
	width  float64
	length float64
}

func NewRectangle(w, l float64) Rectangle {
	return Rectangle{
		width:  w,
		length: l,
	}
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (r *Rectangle) SetWidth(w float64) {
	r.width = w
}

func (r Rectangle) SetLength(l float64) {
	r.length = l
}
