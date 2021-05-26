package calc

func Add(x, y float64) float64 {
	return x + y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Divide(x, y float64) float64 {
	if y == 0 {
		panic("no number is divisible by zero")
	}
	return x / y
}
