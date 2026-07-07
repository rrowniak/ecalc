# Electronics Calculator
This is a command-line calculator written in Go that simplifies common electronics calculations.

## Building
To download and build the calculator, follow these simple steps:
```
$ cd $GOPATH/src
$ git clone https://github.com/rrowniak/ecalc.git
$ go build ecalc
$ go install ecalc
```
You can also run optional tests by executing:
```
$ go test ecalc/logic
$ go run ecalc help
```
## Usage
###### Help
Get basic help by running:
```
$ ecalc help
```
###### Ohm's law
Calculate current, voltage, or resistance according to [Ohm's Law](https://en.wikipedia.org/wiki/Ohm%27s_law). Here are some examples:
```
$ ecalc ohm 1A 1V
1.00 Ω (1.000000 Ω), power 1.00 W

$ ecalc ohm 5V 13mA
384.62 Ω (384.615385 Ω), power 65.00 mW

$ ecalc ohm 13.7V 2k2
6.23 mA (0.006227 A), power 85.31 mW
```
###### Voltage divider
Calculate the output voltage (vout), input voltage (vin), or resistor values (r1, r2) for a simple voltage divider circuit. You can refer to the [Wiki](https://en.wikipedia.org/wiki/Voltage_divider) for more information. Here are some examples:
```
$ ecalc vdiv -vin=12V -r1=1k -r2=2k
vout = 8.00 V (8.000000 V)

$ ecalc vdiv -vin=12V -vout=5V -r1=2k2
r2 = 1.57 kΩ (1571.428571 Ω)

$ ecalc vdiv -vin=230V -vout=24V -r2=10k
r1 = 85.83 kΩ (85833.333333 Ω)
```
###### LC resonant calculator
Provide two quantities (frequency, capacitance, or inductance) and the third one will be calculated. You can learn more about LC circuits [here](https://en.wikipedia.org/wiki/LC_circuit). Here are some examples:
```
$ ecalc lc 10pF 33uH
8.76 MHz (8761191.269246 Hz)

$ ecalc lc 1kHz 1uF
25.33 mH (0.025330 H)

$ ecalc lc 1MHz 10uH
2.53 nF (0.000000 F)
```

###### Reactance calculator
Provide two of the following: frequency, capacitance, inductance, or reactance. You can find more information on the [Wiki page](https://en.wikipedia.org/wiki/Electrical_reactance). Here are some examples:
```
$ ecalc react 1kHz 1uF
Capacitive reactance: 159.15 Ω (159.154943 Ω)

$ ecalc react 1MHz 1kR
Inductance: 159.15 uH (0.000159 H)
Capacitance: 159.15 pF (0.000000 F)

$ ecalc react 5uH 225kHz
Inductive reactance: 7.07 Ω (7.068583 Ω)

$ ecalc react 10uH 33nF
Resonant frequency: 277.05 kHz (277053.194272 Hz)
```

###### dB (decibel) calculator
Provide two scalars and the calculator will provide the ratio in dB. You can learn more on the [Wiki page](https://en.wikipedia.org/wiki/Decibel). Here are some examples:
```
$ ecalc db 10 998,7
Power ratio:     -19.99 dB (-19.994350 dB)
Amplitude ratio: -39.99 dB (-39.988701 dB)
```

###### E series (resistors)
For any given value, the calculator will find the closest match in the E series of resistors. You can find more details about the E series [here](https://en.wikipedia.org/wiki/E_series_of_preferred_numbers). Here are some examples:
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

###### Wire resistance
Calculate the resistance of a copper wire given its diameter, cross-section, or AWG gauge and length. Practical ampacity estimates with voltage drop and dissipation are shown. An optional current (e.g. 5A) can be specified for load-specific V_drop and dissipation. Arguments are order-independent. Here are some examples:
```
$ ecalc wire 1.5mm 60cm
Resistance: 5.85 mΩ (0.005854 Ω)
Cross-section: 1.77 mm², diameter: 1.50 mm
Practical ampacity (power wiring): ~7.07 A (V_drop 41.38 mV, dissipation 292.47 mW)
Practical ampacity (chassis wiring): ~17.67 A (V_drop 103.44 mV, dissipation 1.83 W)

$ ecalc wire 2.5 100m
Resistance: 689.60 mΩ (0.689600 Ω)
Cross-section: 2.50 mm², diameter: 1.78 mm
Practical ampacity (power wiring): ~10.00 A (V_drop 6.90 V, dissipation 68.96 W)
Practical ampacity (chassis wiring): ~25.00 A (V_drop 17.24 V, dissipation 431.00 W)

$ ecalc wire 14AWG 100ft
Resistance: 252.52 mΩ (0.252522 Ω)
Cross-section: 2.08 mm², diameter: 1.63 mm
Practical ampacity (power wiring): ~8.32 A (V_drop 2.10 V, dissipation 17.47 W)
Practical ampacity (chassis wiring): ~20.81 A (V_drop 5.25 V, dissipation 109.28 W)

$ ecalc wire 2.5mm2 10m 5A
Resistance: 68.96 mΩ (0.068960 Ω)
Cross-section: 2.50 mm², diameter: 1.78 mm
Practical ampacity (power wiring): ~10.00 A (V_drop 689.60 mV, dissipation 6.90 W)
Practical ampacity (chassis wiring): ~25.00 A (V_drop 1.72 V, dissipation 43.10 W)
At 5.00 A: V_drop 344.80 mV, dissipation 1.72 W
```
