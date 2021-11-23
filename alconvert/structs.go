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
	finalMl float32

	// calculated amount to remove, to get to finalMl
	// this could be a negative number indicating amount to add
	finalRemoveAmount float32
}

type calcTargetPercent struct {
	// amount of water (in ml) to add in order to reach final_target_percent_all
	finalTargetPercent float32

	// total amount after adding water for target percent
	finalTargetPercentAll float32
}

type calcTargetMl struct {
	// if water is added this is the percent it becomes
	finalTargetMlPercent float32

	// the difference between starting ml and target ml
	finalTargetMlDiff float32
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
