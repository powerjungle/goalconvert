package alconvert

import (
	"fmt"
	"testing"
)

func checkAllZero(alcval *Alcovalues) bool {
	if alcval.Milliliters == 0 &&
		alcval.Percent == 0 &&
		alcval.UnitTarget == 0 &&
		alcval.PercenTarget == 0 &&
		alcval.TargetMl == 0 &&
		alcval.gotUnits == 0 &&
		alcval.finalTargetUnitsMl == 0 &&
		alcval.finalRemoveAmount == 0 &&
		alcval.finalTargetPercent == 0 &&
		alcval.finalTargetPercentAll == 0 &&
		alcval.finalTargetMlPercent == 0 &&
		alcval.finalTargetMlDiff == 0 {
		return true
	}
	fmt.Println(alcval)
	return false
}

// TestDefaultOutput tests whether when calling the functions
// without entering any data in Alcovalues would return
// the appropriate value
func TestDefaultOutput(t *testing.T) {
	av := NewAV()

	CalcGotUnits(av)
	ret := checkAllZero(av)
	if ret != true {
		t.Fatal("CalcGotUnits() doesn't keep all values to 0 without input")
	}

	CalcTargetUnits(av)
	ret = checkAllZero(av)
	if ret != true {
		t.Fatal("CalcTargetUnits() doesn't keep all values to 0 without input")
	}

	CalcTargetPercent(av)
	ret = checkAllZero(av)
	if ret != true {
		t.Fatal("CalcTargetPercent() doesn't keep all values to 0 without input")
	}

	CalcTargetMl(av)
	ret = checkAllZero(av)
	if ret != true {
		t.Fatal("CalcTargetMl() doesn't keep all values to 0 without input")
	}
}

func TestResetAV(t *testing.T) {
	av := NewAV()

	av.Milliliters = 1
	av.Percent = 2
	av.UnitTarget = 3
	av.PercenTarget = 4
	av.TargetMl = 5

	av.gotUnits = 6
	av.finalTargetUnitsMl = 7
	av.finalRemoveAmount = 8
	av.finalTargetPercent = 9
	av.finalTargetPercentAll = 10
	av.finalTargetMlPercent = 11
	av.finalTargetMlDiff = 12

	ResetAV(av)

	ret := checkAllZero(av)
	if ret != true {
		t.Fatal("ResetAV() does not reset all values properly")
	}
}

func BenchmarkCalcGotUnits(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		av.Milliliters = float32(i)
		av.Percent = float32(i)
		CalcGotUnits(av)
	}
}

func BenchmarkCalcTargetUnits(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		av.Milliliters = float32(i)
		av.Percent = float32(i)
		av.UnitTarget = float32(i)
		CalcTargetUnits(av)
	}
}

func BenchmarkCalcTargetPercent(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		av.Milliliters = float32(i)
		av.Percent = float32(i)
		av.PercenTarget = float32(i)
		CalcTargetPercent(av)
	}
}

func BenchmarkCalcTargetMl(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		av.Milliliters = float32(i)
		av.Percent = float32(i)
		av.TargetMl = float32(i)
		CalcTargetMl(av)
	}
}
