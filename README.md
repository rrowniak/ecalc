# ecalc
A simple command line calculator for electronics that ease the most common calculations.
## Building
The calculator is written in Go. Download and build process is as simple as:
```
$ cd $GOPATH/src
$ git clone https://github.com/rrowniak/ecalc.git
$ go build ecalc
$ go install ecalc
# optional tests
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
and two resistors (r1, r2). You can refferf to [Wiki](https://en.wikipedia.org/wiki/Voltage_divider). Having these three known quantities, the calculator will calculate remaining unknown quantity.
Examles:
```
$ ecalc vdiv -vin=12V -r1=1k -r2=2k
vout = 8.00 V (8.000000 V)

$ ecalc vdiv -vin=12V -vout=5V -r1=2k2
r2 = 1.57 kΩ (1571.428571 Ω)

$ ecalc vdiv -vin=230V -vout=24V -r2=10k
r1 = 85.83 kΩ (85833.333333 Ω)
```
###### E series (resistors)
More details on E series can be found [here](https://en.wikipedia.org/wiki/E_series_of_preferred_numbers)
For any given value calculator will find out the closest match.