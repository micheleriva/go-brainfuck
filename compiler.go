package main

var maxVal int = 255
var a []int
var n []int
var p int = 0
var x int = 0
var output string = ""

func compile(program string) string {

	for i := 0; i < len(program); i++ {

		switch program[i] {
			case '>':
				incrementPointer()
			case '<':
				decrementPointer()
			case '+':
				incrementByte()
			case '-':
				decrementByte()
			case '.':
				outputByte()
			case '[':
				beginLoop(i)
			case ']':
				i = endLoop(i)
		}
	}

	return output

}

func incrementPointer() {
	if(x == 0) {
		p++
	}
}

func decrementPointer() {
	if(x == 0){
		p--
	}
}

func incrementByte() {
	if(x == 0){
		if(a[p] < maxVal) {
			a[p] = a[p] + 1
		} else {
			a[p] = 0
		}
	}
}

func decrementByte() {
	if(x == 0) {
		if(a[p] > 0) {
			a[p] = a[p] - 1
		} else {
			a[p] = maxVal
		}
	}
}

func outputByte() {
	if(x == 0) {
		output = output + string(byte(a[p]))
	}
}

func beginLoop(i int) {

	if(x != 0) {
		x++
		return
	}

	if(a[p] > 0) {
		n = append(n, i)
	} else {
		x++
	}
}

func endLoop(i int) int {

	if(x != 0) {
		x--
		return i
	}

	if(a[p] > 0) {
		origin := n[len(n) - 1]
		i = origin - 1
	}

	return i
}