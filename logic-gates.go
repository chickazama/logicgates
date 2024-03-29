package logicgates

type T interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func AND[n T](a, b n) n {
	return a & b
}

func OR[n T](a, b n) n {
	return a | b
}

func NAND[n T](a, b n) n {
	return ^(a & b)
}

func NOR[n T](a, b n) n {
	return ^(a | b)
}

func XOR[n T](a, b n) n {
	return (a & ^b) | (^a & b)
}

func XNOR[n T](a, b n) n {
	return (a & b) | (^a & ^b)
}
