package gf256

// Polynom represents a polynom
type Polynom struct {
	Coefficients []byte
}

// Evaluate computes p at x in GF(2^8)
func (p Polynom) Evaluate(x byte) byte {
	var result byte
	for _, c := range p.Coefficients {
		result = Add(Mul(result, x), c)
	}
	return result
}

// Interpolate computes the polynom represented by points at x
// It uses Lagrange interpolation
func Interpolate(points [][2]byte, x byte) byte {
	var result, elt byte
	for i := range points {
		elt = 1
		for j := range points {
			if i == j {
				continue
			}
			term := Div(Sub(x, points[j][0]), Sub(points[i][0], points[j][0]))
			elt = Mul(elt, term)
		}
		result = Add(result, Mul(points[i][1], elt))
	}
	return result
}
