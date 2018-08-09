package m

import (
	"testing"
	//"fmt"
)

func TestAvg(t *testing.T) {
	if v := Avg([]float64{1,2}); v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
	if v := Sum([]float64{1,2}); v != 3 {
		t.Error("Expected 3, got ", v)
	}
}

func TestSum(t *testing.T) {
}