package prime_test

// import "fmt"
import "testing"
import "prime"

func Test10(t *testing.T) {
	slice := make([]int, 0)
	prime.CalcPrime(10, &slice)
	// fmt.Println("test len is ", len(slice))

	if len(slice) != 4 {
		t.Error("Expected 4, got ", len(slice))
	}
}

func Test100(t *testing.T) {
	slice := make([]int, 0)
	prime.CalcPrime(100, &slice)
	// fmt.Println("test len is ", len(slice))

	if len(slice) != 25 {
		t.Error("Expected 25, got ", len(slice))
	}
}

func Test1000(t *testing.T) {
	slice := make([]int, 0)
	prime.CalcPrime(1000, &slice)

	if len(slice) != 168 {
		t.Error("Expected 168, got ", len(slice))
	}
}

func Test10000(t *testing.T) {
	slice := make([]int, 0)
	prime.CalcPrime(10000, &slice)

	if len(slice) != 1229 {
		t.Error("Expected 1229, got ", len(slice))
	}
}

func Test20000(t *testing.T) {
	slice := make([]int, 0)
	prime.CalcPrime(20000, &slice)

	if len(slice) != 2262 {
		t.Error("Expected 1229, got ", len(slice))
	}
}
