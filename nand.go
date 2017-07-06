package main

import "fmt"

func nand(input1 bool, input2 bool) bool {
	return !(input1 && input2)
}

func not(input1 bool) bool {
	return nand(input1, input1)
}

func and(input1 bool, input2 bool) bool {
	return not(nand(input1, input2))
}

func or(input1 bool, input2 bool) bool {
	return nand(not(input1), not(input2))
}

func xor(input1 bool, input2 bool) bool {
	return or(and(input1, not(input2)), and(not(input1), input2))
}

func halfAdder(input1 bool, input2 bool) (bool, bool) {
	return xor(input1, input2), and(input1, input2)
}

func fullAdder(input1 bool, input2 bool, input3 bool) (bool, bool) {
	a, b := halfAdder(input1, input2)
	c, d := halfAdder(a, input3)
	return c, or(b, d)
}

func adder4Bit(value1 uint, value2 uint) uint {
	var result uint
	carry := false

	a, carry := fullAdder(getBool(value1, 0), getBool(value2, 0), carry)
	result = setBool(a, result, 3)
	b, carry := fullAdder(getBool(value1, 1), getBool(value2, 1), carry)
	result = setBool(b, result, 2)
	c, carry := fullAdder(getBool(value1, 2), getBool(value2, 2), carry)
	result = setBool(c, result, 1)
	d, carry := fullAdder(getBool(value1, 3), getBool(value2, 3), carry)
	result = setBool(d, result, 0)

	return result
}

func getBool(value uint, position uint) bool {
	return value>>position&1 == 1
}

func setBool(input1 bool, value uint, position uint) uint {
	var x uint
	if input1 {
		x = 1
	} else {
		x = 0
	}

	return x<<position + value
}

func main() {
	fmt.Println("NAND")
	fmt.Printf("%t\n", nand(false, false))
	fmt.Printf("%t\n", nand(false, true))
	fmt.Printf("%t\n", nand(true, false))
	fmt.Printf("%t\n", nand(true, true))
	fmt.Println("====")

	fmt.Println("XOR")
	fmt.Printf("%t\n", xor(false, false))
	fmt.Printf("%t\n", xor(false, true))
	fmt.Printf("%t\n", xor(true, false))
	fmt.Printf("%t\n", xor(true, true))
	fmt.Println("====")

	fmt.Println("halfAdder")
	a, b := halfAdder(false, false)
	fmt.Printf("%t, %t\n", a, b)
	a, b = halfAdder(false, true)
	fmt.Printf("%t, %t\n", a, b)
	a, b = halfAdder(true, false)
	fmt.Printf("%t, %t\n", a, b)
	a, b = halfAdder(true, true)
	fmt.Printf("%t, %t\n", a, b)
	fmt.Println("====")

	fmt.Println("fullAdder")
	a, b = fullAdder(false, false, false)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(false, false, true)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(false, true, false)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(false, true, true)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(true, false, false)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(true, false, true)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(true, true, false)
	fmt.Printf("%t, %t\n", a, b)
	a, b = fullAdder(true, true, true)
	fmt.Printf("%t, %t\n", a, b)
	fmt.Println("====")

	fmt.Println("getBool")
	fmt.Printf("%t\n", getBool(1, 0))
	fmt.Printf("%t\n", getBool(2, 0))
	fmt.Printf("%t\n", getBool(3, 0))
	fmt.Printf("%t\n", getBool(3, 1))
	fmt.Println("====")

	fmt.Println("adder4Bit")
	fmt.Printf("%d\n", adder4Bit(2, 4))
	fmt.Println("====")
}
