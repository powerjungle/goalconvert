# Go Alcohol Converter

Using this you can convert alcohol (drinkable) milliliters, percentage and units.

To use the terminal program, after compiling just do:

`./goalconvert -help`

To compile run this in the current directory (where the `README.md` file is):

`go build`

You will need Go 1.16 or later

If you want to run the tests, enter the "alconvert" directory and run:

`go test`

To run the benchmarks, enter the "alconvert" directory and run:

`go test -bench=.`

## Using Terminal Program Example:

Calculate units:

`./goalconvert -ml 200 -perc 40`

Calculate target units:

`./goalconvert -ml 200 -perc 40 -taruni 2`

## Using Module Example:

Calculate units:

```go
package main

import "github.com/powerjungle/goalconvert/alconvert"

func main() {
  av := alconvert.NewAV()
  av.Milliliters = 200
  av.Percent = 40
  alconvert.CalcGotUnits(av)
  alconvert.PrintForHumans(av)
}
```

Calculate target units:

```go
av := alconvert.NewAV()
av.Milliliters = 200
av.Percent = 40
av.UnitTarget = 2
alconvert.CalcTargetUnits(av)
alconvert.PrintForHumans(av)
```

## Functions

- `NewAV()`

  - Create a new instance with alcohol values set as 0

- `ResetAV(avInstanceHere)`

  - Reset an existing instance values back to 0

- `PrintForHumans(avInstanceHere)`

  - Print a human readable-ish text explaining values which aren't 0 and which have changed in relation to the starter values (if calculations were done on them)

- `PrintJson(avInstanceHere)`

  - Print a JSON representation of all alcohol values

- `CalcGotUnits(avInstanceHere)`

  - Calculate `GotUnits` based on set `Milliliters` and `Percent`, these are the units as pure alcohol content present, 1 unit = 10ml pure alcohol

- `CalcTargetUnits(avInstanceHere)`

  - Calculate the amount of alcohol that needs to be removed so that the set `UnitTarget` can be reached

- `CalcTargetPercent(avInstanceHere)`

  - Calculate the diluted alcohol amount that needs to be reached in order to reach `PercenTarget`

- `CalcTargetMl(avInstanceHere)`

  - Calculate if adding amount of water that is needed to reach `TargetMl`, what the percentage will be, and how much water needs to be added

## Alcohol values

Used to initialize calculations

- `Milliliters`

  - Starting milliliters of alcohol (for example beer: 500)

- `Percent`

  - Starting percentage of alcohol (for example beer: 5)

- `UnitTarget`

  - Units you want to reach by using `CalcTargetUnits()`

- `PercenTarget`

  - Percentage you want to reach by using `CalcTargetPercent()`

- `TargetMl`

  - Milliliters you want to reach by using `CalcTargetMl()`

Used for calculation results (unexported) and only used internally

- `gotUnits`

  - The units calculated from `Milliliters` and `Percent` by using `CalcGotUnits()`

- `finalTargetUnitsMl`

  - Amount of alcohol left after removing `FinalRemoveAmount` from `Milliliters` to reach `UnitTarget`

- `finalRemoveAmount`

  - Amount of alcohol to be removed in order to reach `UnitTarget`

- `finalTargetPercent`

  - Amount of water to add in order to reach `PercenTarget`

- `finalTargetPercentAll`

  - Diluted alcohol left after adding `FinalTargetPercent` to reach `PercenTarget`

- `finalTargetMlPercent`

  - Alcohol becomes this percentage after adding water for `TargetMl`

- `finalTargetMlDiff`

  - Total amount of water that needs to be added to reach `TargetMl`

- `timestamp`

  - Time of last calculation

- `lastOperation`

  - The last function used for calculating
