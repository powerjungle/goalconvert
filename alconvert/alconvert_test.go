package alconvert

import (
	"fmt"
	"math/rand"
	"testing"
)

func checkAllZero(alcval *Alcovalues) bool {
	if alcval.UserSet.Milliliters == 0 &&
		alcval.UserSet.Percent == 0 &&
		alcval.UserSet.UnitTarget == 0 &&
		alcval.UserSet.PercenTarget == 0 &&
		alcval.UserSet.TargetMl == 0 &&
		alcval.calcGotUnits.gotUnits == 0 &&
		alcval.calcTargetUnits.finalMl == 0 &&
		alcval.calcTargetUnits.finalRemoveAmount == 0 &&
		alcval.calcTargetPercent.finalTargetPercent == 0 &&
		alcval.calcTargetPercent.finalTargetPercentAll == 0 &&
		alcval.calcTargetMl.finalTargetMlPercent == 0 &&
		alcval.calcTargetMl.finalTargetMlDiff == 0 {
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

	av.UserSet.Milliliters = 1
	av.UserSet.Percent = 2
	av.UserSet.UnitTarget = 3
	av.UserSet.PercenTarget = 4
	av.UserSet.TargetMl = 5

	av.calcGotUnits.gotUnits = 6

	av.calcTargetUnits.finalMl = 7
	av.calcTargetUnits.finalRemoveAmount = 8

	av.calcTargetPercent.finalTargetPercent = 9
	av.calcTargetPercent.finalTargetPercentAll = 10

	av.calcTargetMl.finalTargetMlPercent = 11
	av.calcTargetMl.finalTargetMlDiff = 12

	ResetAV(av)

	ret := checkAllZero(av)
	if ret != true {
		t.Fatal("ResetAV() does not reset all values properly")
	}
}

func BenchmarkCalcGotUnits(b *testing.B) {
	av := NewAV()
	var randval float32
	for i := 0; i < b.N; i++ {
		// since no seed is set in Go by default
		// the seed gets set to 1, which means
		// the random values don't change
		// https://pkg.go.dev/math/rand#Seed
		randval = rand.Float32()
		av.UserSet.Milliliters = float32(randval * 100)
		av.UserSet.Percent = float32(randval * 10)
		CalcGotUnits(av)
	}
}

func BenchmarkCalcTargetUnits(b *testing.B) {
	av := NewAV()
	var randval float32
	for i := 0; i < b.N; i++ {
		randval = rand.Float32()
		av.UserSet.Milliliters = float32(randval * 100)
		av.UserSet.Percent = float32(randval * 10)
		av.UserSet.UnitTarget = float32(randval * 10)
		CalcTargetUnits(av)
	}
}

func BenchmarkCalcTargetPercent(b *testing.B) {
	av := NewAV()
	var randval float32
	for i := 0; i < b.N; i++ {
		randval = rand.Float32()
		av.UserSet.Milliliters = float32(randval * 100)
		av.UserSet.Percent = float32(randval * 10)
		av.UserSet.PercenTarget = float32(randval * 10)
		CalcTargetPercent(av)
	}
}

func BenchmarkCalcTargetMl(b *testing.B) {
	av := NewAV()
	var randval float32
	for i := 0; i < b.N; i++ {
		randval = rand.Float32()
		av.UserSet.Milliliters = float32(randval * 100)
		av.UserSet.Percent = float32(randval * 10)
		av.UserSet.TargetMl = float32(randval * 1000)
		CalcTargetMl(av)
	}
}
