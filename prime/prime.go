package prime

// import "fmt"

// const num int = 100000
const Channum int = 10000

// var Slice []int
// var G_chan chan string = make(chan string)

func prime(c chan int, base int, slice *[]int, strchan chan string) {

	flag := 0
	ch := make(chan int, Channum)

	flag2 := 0

	*slice = append(*slice, <-c)
	// fmt.Println("prime len", len(*slice), *slice)
	// v := <-c
	for {
		if value, ok := <-c; ok {
			flag2++
			if value%base != 0 {
				if flag == 0 {
					go prime(ch, value, slice, strchan)
					flag = 1
				}

				ch <- value
			}
		} else {
			break
		}
	}

	if flag2 == 0 {
		// fmt.Println("last is", v)
		strchan <- "send count over"
	}

	close(ch)
}

func CalcPrime(num int, slice *[]int) {

	var strchan chan string = make(chan string)
	flag := 0

	c := make(chan int, Channum)

	*slice = append(*slice, 2)
	// fmt.Println("main len", len(*slice), *slice)

	for i := 2; i <= num; i += 1 {
		if i%2 != 0 {
			if flag == 0 {
				go prime(c, i, slice, strchan)
				flag = 1
			}

			// fmt.Println("main send", i)
			c <- i
		}
	}

	close(c)

	if _, ok := <-strchan; ok {
		// fmt.Println("slice len is", len(*slice))
		// fmt.Println(slice)
	}
}

// func main() {
// 	CalcPrime()
// }
