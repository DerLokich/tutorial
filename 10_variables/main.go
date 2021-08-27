package main

func main() {
	var i int = 10
	var autoint = -10
	var bigint int64 = 1<<32 - 1
	var unsignedint uint = 100500
	var unsignedbigint uint64 = 1<<64 - 1
	println("integers", i, autoint, bigint, unsignedbigint, unsignedint)

	// var un = 23
	var p float32 = 3.14
	println("float: ", p)

	var b = true
	println("bool bariable", b)

	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)

	meaningOfLife := 42
	println("Meaning of life is ", meaningOfLife)

	println("float to int covers ", int(p))

	z := 2 + 3i
	println("complex number: ", z)

	s1 := "Vasiliy"
	s2 := "Romanov"
	fullname := s1 + " " + s2
	println(fullname)

	escaping := `Hello\r\n
	World`
	println("as-is escaping: ", escaping)

	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("def values: ", defaultInt, defaultFloat, defaultString, defaultBool)

	var v1, v2 string = "v1", "v2"

	println(v1, v2)

	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)
	println(m0, m2, m3)

}
