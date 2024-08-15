package must

import (
	"fmt"
	"math/big"
	"time"
)

func Example_timeParse() {
	tm := Get2(time.Parse(time.RFC3339, "2020-01-01T11:22:33Z"))
	fmt.Println(tm)
	// Output: 2020-01-01 11:22:33 +0000 UTC
}

func Example_parseFloat() {
	v, _ := Get3(big.ParseFloat("1.2345678", 10, 1000, big.ToNearestEven))
	fmt.Println(v)
	// Output: 1.2345678
}
