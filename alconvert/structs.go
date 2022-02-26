package alconvert

import "time"

type UserSet struct {
	// starting amount
	Milliliters float32
	// alcohol percentage (concentration)
	Percent float32
	// needed units (target units)
	UnitTarget float32
	// needed percentage (target percentage/concentration)
	PercenTarget float32
	// needed milliliters (target ml/amount)
	TargetMl float32
}

type calcGotUnits struct {
	// the calculated units using the present Ml and UserSet.Percent
	gotUnits float32
}

type calcTargetUnits struct {
	// calculated milliliters needed for target units at same concentration
	gotTargUnitsFinalAmount float32

	// calculated amount to remove, to get to gotTargUnitsFinalAmount
	// this could be a negative number indicating amount to add
	gotTargUnitsRemAmount float32
}

type calcTargetPercent struct {
	// amount of water (in ml) to add in order to reach gotTargPercAlcLeft
	gotTargPercAddWater float32

	// total amount after adding water for target percent
	gotTargPercAlcLeft float32
}

type calcTargetMl struct {
	// if water is added this is the percent it becomes
	gotTargMlNewAlcPerc float32

	// the difference between starting ml and target ml
	gotTargMlNeededWater float32
}

// Alcovalues contains all of the variables used when calculating and initial setup
type Alcovalues struct {
	// values that should be used as initializing inputs////////////////////
	UserSet UserSet

	// values that are set by the functions in this file////////////////////
	calcGotUnits      calcGotUnits
	calcTargetUnits   calcTargetUnits
	calcTargetPercent calcTargetPercent
	calcTargetMl      calcTargetMl

	// time of last calculation/////////////////////////////////////////////
	timestamp time.Time

	// the last function used for calculating///////////////////////////////
	lastOperation string
}
