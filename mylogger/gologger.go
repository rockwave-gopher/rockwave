package mylogger

import (
	"fmt"
	"log"
)

//fatal
func TestDeferfatal() {
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}
//panic
func testDeferpanic() {
	defer func() {
		fmt.Println("--first--")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panicln("test for defer Panic")
	defer func() {
		fmt.Println("--second--")
	}()
}
//func main() {
//	arr := []int{2, 3}
//
//	log.Print("Print array ", arr, "\n")
//	log.Println("Println array", arr)
//	log.Printf("Printf array with item [%d,%d]\n", arr[0], arr[1])
//
//	testDeferpanic()
//
//	TestDeferfatal()
//}