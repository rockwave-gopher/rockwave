package mytest

func Add(x int, y int) (z int) {
	z = x + y
	return z
}

type ForTest struct {
	num int
}

func (t *ForTest) Loops() {
	for i := 0; i < 10000; i++ {
		t.num++
	}
}
