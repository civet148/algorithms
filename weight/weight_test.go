package weight

import (
	"fmt"
	"testing"
)

const (
	ip5 = "192.168.1.5"
	ip3 = "192.168.1.3"
	ip1 = "192.168.1.1"
)

func TestAlgorithmWeight(t *testing.T) {
	alg := NewAlgorithmWeight()
	alg.Add(ip5, ip5, 5)
	alg.Add(ip3, ip3, 3)
	alg.Add(ip1, ip1, 1)

	for i := 0; i < 9; i++ {
		v := alg.Get()
		if v == nil {
			fmt.Printf("nil value\n")
			return
		}
		fmt.Printf("[%d] %s\n", i, v.(string))
	}
	/*  ----------------------
	[0] 192.168.1.5
	[1] 192.168.1.3
	[2] 192.168.1.5
	[3] 192.168.1.1
	[4] 192.168.1.5
	[5] 192.168.1.3
	[6] 192.168.1.5
	[7] 192.168.1.3
	[8] 192.168.1.5
	*/
}
