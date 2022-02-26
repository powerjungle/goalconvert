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
		alcval.calcTargetUnits.gotTargUnitsFinalAmount == 0 &&
		alcval.calcTargetUnits.gotTargUnitsRemAmount == 0 &&
		alcval.calcTargetPercent.gotTargPercAddWater == 0 &&
		alcval.calcTargetPercent.gotTargPercAlcLeft == 0 &&
		alcval.calcTargetMl.gotTargMlNewAlcPerc == 0 &&
		alcval.calcTargetMl.gotTargMlNeededWater == 0 {
		return true
	}
	fmt.Println(alcval)
	return false
}

// TestDefaultOutput tests whether when calling the functions
// without entering any data in Alcovalues, the return
// will have the appropriate values.
func TestDefaultOutput(t *testing.T) {
	av := NewAV()

	av.CalcGotUnits()
	ret := checkAllZero(av)
	if ret != true {
		t.Fatal("CalcGotUnits() doesn't keep all values to 0 without input")
	}

	av.CalcTargetUnits()
	ret = checkAllZero(av)
	if ret != true {
		t.Fatal("CalcTargetUnits() doesn't keep all values to 0 without input")
	}

	av.CalcTargetPercent()
	ret = checkAllZero(av)
	if ret != true {
		t.Fatal("CalcTargetPercent() doesn't keep all values to 0 without input")
	}

	av.CalcTargetMl()
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

	av.calcTargetUnits.gotTargUnitsFinalAmount = 7
	av.calcTargetUnits.gotTargUnitsRemAmount = 8

	av.calcTargetPercent.gotTargPercAddWater = 9
	av.calcTargetPercent.gotTargPercAlcLeft = 10

	av.calcTargetMl.gotTargMlNewAlcPerc = 11
	av.calcTargetMl.gotTargMlNeededWater = 12

	av.ResetAV()

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
		av.CalcGotUnits()
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
		av.CalcTargetUnits()
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
		av.CalcTargetPercent()
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
		av.CalcTargetMl()
	}
}
