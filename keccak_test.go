package keccak

import (
	"encoding/hex"
	"fmt"
	"hash"
	"testing"
	"time"
)

type test struct {
	input  []byte
	output []string
}

var tests = []test{
	{
		[]byte("Keccak-256 Test Hash"),
		[]string{"a8d71b07f4af26a4ff21027f62ff60267ff955c963f042c46da52ee3cfaf3d3c"},
	},
}

func TestKeccak(t *testing.T) {
	startTime := time.Now()
	k := 100
	for i := 0; i < k; i++ {
		for i := range tests {
			var h hash.Hash
			h = New256()
			h.Write(tests[i].input)
			d := h.Sum(nil)
			encode := hex.EncodeToString(d)
			fmt.Println(encode)
		}
		elapsedTime := time.Since(startTime) / time.Millisecond
		fmt.Println("Segment finished in", elapsedTime) //Segment finished in xxms
	}
}
