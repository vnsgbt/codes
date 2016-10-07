package exercise

import (
	"fmt"
	"testing"
)

type InOut struct {
	In     int
	Expect int
}

func InOutTest(test func(int) int, i InOut) error {
	if actual := test(i.In); actual != i.Expect {
		return fmt.Errorf("Input %v, Expected: %v, Actual: %v \n", i.In, i.Expect, actual)
	}
	return nil
}

func TestSum(t *testing.T) {
	sum := 0
	input := 100
	for i := 1; i <= input; i++ {
		sum += i
	}
	if err := InOutTest(Sum, InOut{100, sum}); err != nil {
		t.Fatal(err.Error())
	}
}

func TestPairSumSeq(t *testing.T) {
	if err := InOutTest(PairSumSeq, InOut{2, 4}); err != nil {
		t.Fatal(err.Error())
	}
}

func TestTreeInt(t *testing.T) {
	if err := InOutTest(TreeInt, InOut{4, 15}); err != nil {
		t.Fatal(err.Error())
	}
}
