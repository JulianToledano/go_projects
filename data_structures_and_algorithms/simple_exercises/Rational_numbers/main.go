package main

import "fmt"

type rationalNumber struct {
	numerator   int
	denominator int
}

func main() {
	a := newRationalNumber(-3, -4)
	b := newRationalNumber(2, -5)

	fmt.Println(sumRationals(a, b))
	fmt.Println(multiplyRationals(a, b))

	c := newRationalNumber(2366, 273)
	c.simplify()
	fmt.Println(*c)
}

func (rn *rationalNumber) Init(numerator int, denominator int) {
	rn.numerator = numerator
	rn.denominator = denominator
}

func newRationalNumber(numerator int, denominator int) *rationalNumber {
	if denominator < 0 {
		return &rationalNumber{-numerator, -denominator}
	}
	return &rationalNumber{numerator, denominator}
}

func sumRationals(a *rationalNumber, b *rationalNumber) rationalNumber {
	if a.denominator == b.denominator {
		return *newRationalNumber((a.numerator + b.numerator), a.denominator)
	}
	numerator := ((a.numerator * b.denominator) + (b.numerator * a.denominator))
	return *newRationalNumber(numerator, (a.denominator * b.denominator))
}

func multiplyRationals(a *rationalNumber, b *rationalNumber) rationalNumber {
	return *newRationalNumber(a.numerator*b.numerator, a.denominator*b.denominator)
}

func calculateGCD(a int, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}

func (rn *rationalNumber) simplify() {
	gcd := calculateGCD(rn.numerator, rn.denominator)
	rn.numerator = rn.numerator / gcd
	rn.denominator = rn.denominator / gcd
}
