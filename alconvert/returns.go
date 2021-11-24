package alconvert

import "time"

// The functions in here are used only for returning the values of unexported variables.
// The reason they are unexported is to minimize confusion and potential corruption of the results.

func GotUnits(alcval *Alcovalues) float32 {
	return alcval.calcGotUnits.gotUnits
}

func FinalMl(alcval *Alcovalues) float32 {
	return alcval.calcTargetUnits.finalMl
}

func FinalRemoveAmount(alcval *Alcovalues) float32 {
	return alcval.calcTargetUnits.finalRemoveAmount
}

func FinalTargetPercent(alcval *Alcovalues) float32 {
	return alcval.calcTargetPercent.finalTargetPercent
}

func FinalTargetPercentAll(alcval *Alcovalues) float32 {
	return alcval.calcTargetPercent.finalTargetPercentAll
}

func FinalTargetMlPercent(alcval *Alcovalues) float32 {
	return alcval.calcTargetMl.finalTargetMlPercent
}

func FinalTargetMlDiff(alcval *Alcovalues) float32 {
	return alcval.calcTargetMl.finalTargetMlDiff
}

func Timestamp(alcval *Alcovalues) time.Time {
	return alcval.timestamp
}

func LastOperation(alcval *Alcovalues) string {
	return alcval.lastOperation
}
