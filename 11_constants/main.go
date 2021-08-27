package main

const someInt = 1
const typedInt int32 = 17
const fullName = "Vasiliy"

const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	one = iota + 1
	two
	_
	four
)

const (
	_         = iota
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	pi := 3.14
	println(pi + someInt)
	println(pi + float64(typedInt))
	println(one)
	println(KB, MB, GB, TB, PB, EB)
}
