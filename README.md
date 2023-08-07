# Go Alcohol Converter (module and CLI)

Using this you can convert
[alcohol (drinkable)](https://en.wikipedia.org/wiki/Alcohol_(drug))
[millilitres](https://en.wikipedia.org/wiki/Litre#SI_prefixes_applied_to_the_litre),
[percentage](https://en.wikipedia.org/wiki/Alcohol_by_volume) and
[units](https://en.wikipedia.org/wiki/Unit_of_alcohol).

The units are the UK definition which you can checkout here:
https://www.nhs.uk/live-well/alcohol-advice/calculating-alcohol-units/

From the above link:
"One unit equals 10ml or 8g of pure alcohol, which is around the amount of
alcohol the average adult can process in an hour."

It can be used as a module for your own application or
as a CLI by itself.

You can checkout dosages for units here:
https://psychonautwiki.org/wiki/Alcohol

This project is not for the endorsement of alcohol consumption.
The aim is to help people get a perspective of what different amounts mean.

Stay safe!

## Using Terminal Program Example

Calculate units:

`goalconvert -ml 200 -perc 40`

Calculate target units:

`goalconvert -ml 200 -perc 40 -taruni 2`

Others:

`goalconvert -help`

## Using Module Example

Calculate units:

```go
package main

import "github.com/powerjungle/goalconvert"

func main() {
  av := alconvert.NewAV()
  av.UserSet.Milliliters = 200
  av.UserSet.Percent = 40
  av.CalcGotUnits()
  av.PrintForHumans()
}
```

## Module Documentation

https://pkg.go.dev/github.com/powerjungle/goalconvert

## Dependencies

If you have an already built binary from the "Releases" page,
you don't need these.

For Go compilation/installation you'll need this.

## Installation

If you don't want or can't use the already built binaries
in the "Releases" page.

CLI: `go install github.com/powerjungle/goalconvert/cmd@latest`

If you want to use as a module for your code.

Module: `go get github.com/powerjungle/goalconvert`

## Compilation

`go build .`

Run the command where the `README.md` file is!

## Testing

If you want to run the tests: `go test`

## Benchmarks

To run the benchmarks: `go test -bench=.`

## Releases

The approprate packages need to be installed.
Checkout [Dependencies](#dependencies)!

To install Goreleaser: https://goreleaser.com/install/

To do a release: https://goreleaser.com/cmd/goreleaser/

The `.goreleaser.yaml` file is already done and is in the repo.

If you don't want to release for all OSs and architecture or want for more,
edit the `.goreleaser.yaml` file! Info on how to edit here:

https://goreleaser.com/customization/build/

All Go OS and architecture combos here:

https://go.dev/doc/install/source#environment

This needs to be tested every time it's changed, as not all builds work
without some preparation.

