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
		alcval.GotUnits == 0 &&
		alcval.FinalTargetUnitsMl == 0 &&
		alcval.FinalRemoveAmount == 0 &&
		alcval.FinalTargetPercent == 0 &&
		alcval.FinalTargetPercentAll == 0 &&
		alcval.FinalTargetMlPercent == 0 &&
		alcval.FinalTargetMlDiff == 0 {
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

	av.GotUnits = 6
	av.FinalTargetUnitsMl = 7
	av.FinalRemoveAmount = 8
	av.FinalTargetPercent = 9
	av.FinalTargetPercentAll = 10
	av.FinalTargetMlPercent = 11
	av.FinalTargetMlDiff = 12

	ResetAV(av)

	ret := checkAllZero(av)
	if ret != true {
		t.Fatal("ResetAV() does not reset all values properly")
	}
}

func BenchmarkCalcGotUnits(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		CalcGotUnits(av)
	}
}

func BenchmarkCalcTargetUnits(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		CalcTargetUnits(av)
	}
}

func BenchmarkCalcTargetPercent(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		CalcTargetPercent(av)
	}
}

func BenchmarkCalcTargetMl(b *testing.B) {
	av := NewAV()
	for i := 0; i < b.N; i++ {
		CalcTargetMl(av)
	}
}
