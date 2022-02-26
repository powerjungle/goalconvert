# Go Alcohol Converter

Using this you can convert alcohol (drinkable) milliliters, percentage and units.

It can be used as a module for your own application or as a CLI or GUI application by itself.

After installing/compiling of the CLI app don't forget to run: `goalconvert -help`

For installing or compiling the GUI, you'll need to install these packages:
https://developer.fyne.io/started/#prerequisites

#### Installation

CLI: `go install github.com/powerjungle/goalconvert@latest`

GUI: `go install github.com/powerjungle/goalconvert/alconvgui@latest`

Package: `go get github.com/powerjungle/goalconvert/alconvert`

#### Compilation

`go build .`

For CLI run the command where the README is!

For GUI run the command in the `alconvgui` directory!

You will need Go 1.16 or later

##### Build for Android

Install the fyne CLI utility: `go install fyne.io/fyne/v2/cmd/fyne@latest`

Run inside the `alconvgui` directory:
`fyne package -os android -appID testing.alconvert`

This will create an APK. You'll need to use "Android Debug Bridge".

https://developer.android.com/studio/command-line/adb

After connecting `adb` to your phone, run:
`adb install alconvgui.apk`

#### Testing

If you want to run the tests, enter the "alconvert" directory and run:

`go test`

#### Benchmarks

To run the benchmarks, enter the "alconvert" directory and run:

`go test -bench=.`

## Using Terminal Program Example:

Calculate units:

`goalconvert -ml 200 -perc 40`

Calculate target units:

`goalconvert -ml 200 -perc 40 -taruni 2`

## Using Module Example:

Calculate units:

```go
package main

import "github.com/powerjungle/goalconvert/alconvert"

func main() {
  av := alconvert.NewAV()
  av.UserSet.Milliliters = 200
  av.UserSet.Percent = 40
  av.CalcGotUnits()
  av.PrintForHumans()
}
```

## Documentation

https://pkg.go.dev/github.com/powerjungle/goalconvert/alconvert

