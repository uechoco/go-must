package must_test

import (
	"fmt"
	"math/big"
	"time"

	"github.com/uechoco/go-must"
)

func Example_timeParse() {
	tm := must.Get2(time.Parse(time.RFC3339, "2020-01-01T11:22:33Z"))
	fmt.Println(tm)
	// Output: 2020-01-01 11:22:33 +0000 UTC
}

func Example_parseFloat() {
	v, _ := must.Get3(big.ParseFloat("1.2345678", 10, 1000, big.ToNearestEven))
	fmt.Println(v)
	// Output: 1.2345678
}
