package mytest

import "testing"

type AddArray struct {
	result int
	add1   int
	add2   int
}

//run go test
func TestAdd(t *testing.T) {
	var test_data = [3]AddArray{
		{2, 1, 1},
		{5, 2, 3},
		{4, 2, 2},
	}
	for _, v := range test_data {
		if v.result != Add(v.add1, v.add2) {
			t.Errorf("Add( %d , %d ) != %d \n", v.add1, v.add2, v.result)
		}
	}
}

//go test -test.bench=".*"
func BenchmarkAdd(b *testing.B) {
	var test_data = [3]AddArray{
		{2, 1, 1},
		{5, 2, 3},
		{4, 2, 2},
	}
	for _, v := range test_data {
		if v.result != Add(v.add1, v.add2) {
			b.Errorf("Add( %d , %d ) != %d \n", v.add1, v.add2, v.result)
		}
	}

}

// go test -test.bench=".*" -count=5
//go test -run=文件名字 -bench=bench名字 -cpuprofile=生产的cprofile文件名称 文件夹
func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	var test_data = AddArray{2, 1, 1}

	for i := 0; i < b.N; i++ {
		Add(test_data.add1, test_data.add2)
	}
}
