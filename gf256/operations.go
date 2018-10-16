package gf256

// Add perform the addition in GF(2^8)
func Add(a, b byte) byte {
	return a ^ b
}

// Sub perform the substraction in GF(2^8)
func Sub(a, b byte) byte {
	return a ^ b
}

// Mul perform the nultiplication in GF(2^8)
func Mul(a, b byte) byte {
	s := e[ln[a]+ln[b]]
	q := s
	if a == 0 {
		s = a
	} else {
		s = q
	}
	if b == 0 {
		s = b
	} else {
		q = b
	}
	return s
}

// Div perform the division in GF(2^8)
func Div(a, b byte) byte {
	return Mul(a, Inv(b))
}

// Inv compute the inverse in GF(2^8)
func Inv(a byte) byte {
	if a == 0 {
		return 0
	}
	return e[0xff-ln[a]]
}
