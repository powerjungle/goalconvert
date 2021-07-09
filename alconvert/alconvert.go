package alconvert

import (
	"fmt"
	"encoding/json"
	"os"
)

type Alcovalues struct {
	// values that should be used as initializing inputs////////////////////

	// starting amount
	Milliliters int16
	// alcohol percentage (concentration)
	Percent float32
	// needed units (target units)
	UnitTarget float32
	// needed percentage (target percentage/concentration)
	PercenTarget float32
	// needed milliliters (target ml/amount)
	TargetMl int16
	////////////////////////////////////////////////////////////////////////

	// values that are set by the functions in this file///////////////////

	// the calculated units using the present Ml and Percent
	GotUnits float32
	// calculated milliliters needed for target units at same concentration
	FinalTargetUnitsMl int16
	// calculated amount to remove, to get to FinalTargetUnitsMl
	// this could be a negative number indicating amount to add
	FinalRemoveAmount int16
	// amount of water (in ml) to add in order to reach final_target_percent_all
	FinalTargetPercent int16
	// total amount after adding water for target percent
	FinalTargetPercentAll int16
	// if water is added this is the percent it becomes
	FinalTargetMlPercent float32
	// the difference between starting ml and target ml
	FinalTargetMlDiff int16
	////////////////////////////////////////////////////////////////////////
}

// Create a new instance of Alcovalues
func NewAV() *Alcovalues {
	return &Alcovalues{}
}

// Reset values from an existing instance by going through all fields
// No need for a new instance every time you need a fresh calculation
func ResetAV(alcval *Alcovalues) {
	alcval.Milliliters = 0
	alcval.Percent = 0
	alcval.UnitTarget = 0
	alcval.PercenTarget = 0
	alcval.TargetMl = 0

	alcval.GotUnits = 0
	alcval.FinalTargetUnitsMl = 0
	alcval.FinalRemoveAmount = 0
	alcval.FinalTargetPercent = 0
	alcval.FinalTargetPercentAll = 0
	alcval.FinalTargetMlPercent = 0
	alcval.FinalTargetMlDiff = 0
}

// Print Alcohol Values Human Readable (sorta)
// Use this carefully as it might give you a headache if you constantly spam it
func PrintForHumans(alcval *Alcovalues) {
	fmt.Println("---BEGINNING Of Printed Alcohol Values Human Readable---")
	if alcval.Milliliters != 0 {
		fmt.Printf("milliliters:\n%d\n", alcval.Milliliters)
	}
	if alcval.Percent != 0 {
		fmt.Printf("alcohol percentage (concentration):\n%g\n", alcval.Percent)
	}
	if alcval.UnitTarget != 0 {
		fmt.Printf("needed units (target units):\n%g\n", alcval.UnitTarget)
	}
	if alcval.PercenTarget != 0 {
		fmt.Printf("needed percentage (target percentage/concentration):\n%g\n", alcval.PercenTarget)
	}
	if alcval.TargetMl != 0 && alcval.TargetMl != alcval.Milliliters {
		fmt.Printf("needed milliliters (target ml/amount):\n%d\n", alcval.TargetMl)
	}
	fmt.Println("------")
	if alcval.GotUnits != 0 {
		fmt.Printf("calculated units using milliliters and percentage:\n%g\n", alcval.GotUnits)
	}
	if alcval.FinalRemoveAmount != 0 {
		fmt.Printf("calculated amount of alcohol (in ml) to remove in order to reach the target units (at the same percentage):\n%d\n", alcval.FinalRemoveAmount)
	}
	if alcval.FinalTargetUnitsMl != 0 && alcval.FinalTargetUnitsMl != alcval.Milliliters {
		fmt.Printf("total amount alcohol left after removing calculated alcohol for target units:\n%d\n", alcval.FinalTargetUnitsMl)
	}
	if alcval.FinalTargetPercent != 0 {
		fmt.Printf("calculated amount of water (in ml) to add, to reach target percentage:\n%d\n", alcval.FinalTargetPercent)
	}
	if alcval.FinalTargetPercentAll != 0 && alcval.FinalTargetPercentAll != alcval.Milliliters {
		fmt.Printf("total amount alcohol left after adding calculated water:\n%d\n", alcval.FinalTargetPercentAll)
	}
	if alcval.FinalTargetMlPercent != 0 {
		fmt.Printf("alcohol becomes this percentage(concentration) after adding water for target ml:\n%g\n", alcval.FinalTargetMlPercent)
	}
	if alcval.FinalTargetMlDiff != 0 && alcval.FinalTargetMlDiff != alcval.Milliliters {
		fmt.Printf("total amount of water added in alcohol for target ml:\n%d\n", alcval.FinalTargetMlDiff)
	}
	fmt.Println("---END Of Printed Alcohol Values Human Readable---")
}

func PrintJson(alcval *Alcovalues) {
    ret, err := json.Marshal(alcval)
    if err == nil {
        os.Stdout.Write(ret)
    } else {
        fmt.Println(err)
    }
}

// calculate units from the basic milliliters and percentage in the Alcovalues struct
func CalcGotUnits(alcval *Alcovalues) {
	alcval.GotUnits = (float32(alcval.Milliliters) * (alcval.Percent / 100)) / 10
}

// calculate amount of alcohol that needs to be removed so that the target units can be reached
func CalcTargetUnits(alcval *Alcovalues) {
	alcval.FinalTargetUnitsMl = int16((alcval.UnitTarget * 10) / (alcval.Percent / 100))
	alcval.FinalRemoveAmount = alcval.Milliliters - alcval.FinalTargetUnitsMl
}

// calculate amount of alcohol (diluted) that needs to be reached so that the target percentage is reached
func CalcTargetPercent(alcval *Alcovalues) {
	alcval.FinalTargetPercent = int16((alcval.Percent / alcval.PercenTarget) * float32(alcval.Milliliters) - float32(alcval.Milliliters))
	alcval.FinalTargetPercentAll = alcval.FinalTargetPercent + alcval.Milliliters
}

// calculate the amount of dilution and final percentage if we want to reach the target milliliters
func CalcTargetMl(alcval *Alcovalues) {
	alcval.FinalTargetMlPercent = (float32(alcval.Milliliters) / float32(alcval.TargetMl)) * alcval.Percent
	alcval.FinalTargetMlDiff = alcval.TargetMl - alcval.Milliliters
}
