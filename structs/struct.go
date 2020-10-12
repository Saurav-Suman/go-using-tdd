package structs

type StructData struct {
	a int
	b int
}

func StructSum(data StructData) int {
	return data.a + data.b
}

func (s StructData) StructSumR() int {
	return s.a + s.b
}
