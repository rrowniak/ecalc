# ecalc
A simple command line calculator for electronics that ease the most common calculations.
## Building
The calculator is written in Go. Download and build process is as simple as:
```
$ cd $GOPATH/src
$ git clone https://github.com/rrowniak/ecalc.git
$ go build ecalc
$ go install ecalc
```
Optional tests can be done:
```
$ go test ecalc/logic
$ go run ecalc help
```
## Usage
###### Help
Get basic help:
```
$ ecalc help
```
###### Ohm's law
Calculate current, voltage or resistance according to [the Ohm's law](https://en.wikipedia.org/wiki/Ohm%27s_law).
Examples:
```
$ ecalc ohm 1A 1V
1.00 Ω (1.000000 Ω), power 1.00 W

$ ecalc ohm 5V 13mA
384.62 Ω (384.615385 Ω), power 65.00 mW

$ ecalc ohm 13.7V 2k2
6.23 mA (0.006227 A), power 85.31 mW
```
The above examples are self-explanatory.
###### Voltage divider
Consider a simple voltage divider that consist of input voltage (vin), output voltage (vout)
and two resistors (r1, r2). You can refferr to [Wiki](https://en.wikipedia.org/wiki/Voltage_divider). Having any of these three known quantities, the calculator will calculate remaining forth unknown quantity.
Examles:
```
$ ecalc vdiv -vin=12V -r1=1k -r2=2k
vout = 8.00 V (8.000000 V)

$ ecalc vdiv -vin=12V -vout=5V -r1=2k2
r2 = 1.57 kΩ (1571.428571 Ω)

$ ecalc vdiv -vin=230V -vout=24V -r2=10k
r1 = 85.83 kΩ (85833.333333 Ω)
```
###### LC resonant calculator
Privide two quantities (frequency, capacitance or inductance) and the third one will be calculated.
More about LC circuit [here](https://en.wikipedia.org/wiki/LC_circuit)

Examples:
```
$ ecalc lc 10pF 33uH
8.76 MHz (8761191.269246 Hz)

$ ecalc lc 1kHz 1uF
25.33 mH (0.025330 H)

$ ecalc lc 1MHz 10uH
2.53 nF (0.000000 F)
```

###### Reactance calculator
Privide two quantites out of frequency, capacitance, inductance or reactance.
[Wiki page](https://en.wikipedia.org/wiki/Electrical_reactance).

Examples:
```
ecalc react 1kHz 1uF
```

###### dB (decibell) calculator
Provide two scalars and the caclulator will provide the ratio in dB.
[Wiki page](https://en.wikipedia.org/wiki/Decibel).

Examples:
```
$ ecalc db 10 998,7
Power ratio:     -19.99 dB (-19.994350 dB)
Amplitude ratio: -39.99 dB (-39.988701 dB)
```

###### E series (resistors)
More details on E series can be found [here](https://en.wikipedia.org/wiki/E_series_of_preferred_numbers)
For any given value calculator will find out the closest match.

Examples:
```
$ ecalc eseries 7.482kR
Closest match to 7482.000000 in series E3 (tolerance 40.0%):
        lower boundary: 4.70 (4.70 kΩ), error: 37.2%, diff: 2.78 kΩ
        upper boundary: 1.00 (10.00 kΩ), error: 33.7%, diff: 2.52 kΩ
Closest match to 7482.000000 in series E6 (tolerance 20.0%):
        lower boundary: 6.80 (6.80 kΩ), error: 9.1%, diff: 682.00 Ω
        upper boundary: 1.00 (10.00 kΩ), error: 33.7%, diff: 2.52 kΩ
Closest match to 7482.000000 in series E12 (tolerance 10.0%):
        lower boundary: 6.80 (6.80 kΩ), error: 9.1%, diff: 682.00 Ω
        upper boundary: 8.20 (8.20 kΩ), error: 9.6%, diff: 718.00 Ω
Closest match to 7482.000000 in series E24 (tolerance 5.0%):
        lower boundary: 6.80 (6.80 kΩ), error: 9.1%, diff: 682.00 Ω
        upper boundary: 7.50 (7.50 kΩ), error: 0.2%, diff: 18.00 Ω
Closest match to 7482.000000 in series E48 (tolerance 2.0%):
        lower boundary: 7.15 (7.15 kΩ), error: 4.4%, diff: 332.00 Ω
        upper boundary: 7.50 (7.50 kΩ), error: 0.2%, diff: 18.00 Ω
Closest match to 7482.000000 in series E96 (tolerance 1.0%):
        lower boundary: 7.32 (7.32 kΩ), error: 2.2%, diff: 162.00 Ω
        upper boundary: 7.50 (7.50 kΩ), error: 0.2%, diff: 18.00 Ω
Closest match to 7482.000000 in series E128 (tolerance 0.5%):
        lower boundary: 7.41 (7.41 kΩ), error: 1.0%, diff: 72.00 Ω
        upper boundary: 7.50 (7.50 kΩ), error: 0.2%, diff: 18.00 Ω
```