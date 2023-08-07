package alconvert

import "time"

// The functions in here are used only for returning the values of unexported variables.
// The reason they are unexported is to minimize confusion and potential corruption of the results.

// Returns the pure amount in ml, calculated from UserSet Milliliters and Percent,
// by using CalcPureAmount().
func (alcval *Alcovalues) GotPure() float32 {
	return alcval.calcGotUnits.gotPure
}

// Returns the units calculated from UserSet Milliliters and Percent,
// by using CalcGotUnits().
func (alcval *Alcovalues) GotUnits() float32 {
	return alcval.calcGotUnits.gotUnits
}

// Returns the amount of alcohol left, after removing
// GotTargUnitsRemAmount from UserSet Milliliters to reach UnitTarget.
func (alcval *Alcovalues) GotTargUnitsFinalAmount() float32 {
	return alcval.calcTargetUnits.gotTargUnitsFinalAmount
}

// Returns the amount of alcohol to be removed,
// in order to reach UserSet UnitTarget.
func (alcval *Alcovalues) GotTargUnitsRemAmount() float32 {
	return alcval.calcTargetUnits.gotTargUnitsRemAmount
}

// Returns the amount of water to add,
// in order to reach UserSet PercenTarget.
func (alcval *Alcovalues) GotTargPercAddWater() float32 {
	return alcval.calcTargetPercent.gotTargPercAddWater
}

// Returns the diluted alcohol left, after adding
// GotTargPercAddWater to reach UserSet PercenTarget.
func (alcval *Alcovalues) GotTargPercAlcLeft() float32 {
	return alcval.calcTargetPercent.gotTargPercAlcLeft
}

// Returns the percentage after adding water for UserSet TargetMl.
func (alcval *Alcovalues) GotTargMlNewAlcPerc() float32 {
	return alcval.calcTargetMl.gotTargMlNewAlcPerc
}

// Return the total amount of water that needs to be added,
// to reach UserSet TargetMl.
func (alcval *Alcovalues) GotTargMlNeededWater() float32 {
	return alcval.calcTargetMl.gotTargMlNeededWater
}

// Returns the timestamp of the last operation done.
func (alcval *Alcovalues) Timestamp() time.Time {
	return alcval.timestamp
}

// Return the name of the last operation done.
func (alcval *Alcovalues) LastOperation() string {
	return alcval.lastOperation
}
