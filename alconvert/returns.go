package alconvert

import "time"

// The functions in here are used only for returning the values of unexported variables.
// The reason they are unexported is to minimize confusion and potential corruption of the results.

func (alcval *Alcovalues) GotUnits() float32 {
	return alcval.calcGotUnits.gotUnits
}

func (alcval *Alcovalues) FinalMl() float32 {
	return alcval.calcTargetUnits.finalMl
}

func (alcval *Alcovalues) FinalRemoveAmount() float32 {
	return alcval.calcTargetUnits.finalRemoveAmount
}

func (alcval *Alcovalues) FinalTargetPercent() float32 {
	return alcval.calcTargetPercent.finalTargetPercent
}

func (alcval *Alcovalues) FinalTargetPercentAll() float32 {
	return alcval.calcTargetPercent.finalTargetPercentAll
}

func (alcval *Alcovalues) FinalTargetMlPercent() float32 {
	return alcval.calcTargetMl.finalTargetMlPercent
}

func (alcval *Alcovalues) FinalTargetMlDiff() float32 {
	return alcval.calcTargetMl.finalTargetMlDiff
}

func (alcval *Alcovalues) Timestamp() time.Time {
	return alcval.timestamp
}

func (alcval *Alcovalues) LastOperation() string {
	return alcval.lastOperation
}
