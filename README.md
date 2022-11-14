# LED-Blink

In this post, you will be introduced to the basics of programming hardware using the Viam SDK by making an LED blink. This will allow you to write Python or Golang code to make an LED connected to the GPIO of a Raspberry Pi blink. This tutorial is a good introduction to Python programming, and developing code for hardware like robots and IoT devices.

![A GIF of the completed project showing a blinking blue LED connected to a Raspberry Pi with jumper cables.](https://user-images.githubusercontent.com/4650739/190014115-78f89892-ee67-4a19-846a-33c86e8e6272.gif)

## Circuit Diagram

For reference, the circuit you are building for this tutorial looks like this:

![Circuit diagram showing a Raspberry Pi with a red connector running out of GPIO pin 8 to a 100-ohm* resistor*. The resistor is connected to the long lead of a red LED bulb. Finally, a blue connector connects the short lead of the LED to the ground connection on pin 6 of the Raspberry Pi GPIO pins.](https://user-images.githubusercontent.com/4650739/190014128-1949ef55-47b0-4706-8d6f-2aac99bc6fd7.png)

## Prerequisites

You will need the following hardware, tools, and software to complete this project:

### Hardware

- [Raspberry Pi 3 or 4](https://a.co/d/5Tn67G3)

- [Solderless breadboard](https://amzn.to/2Q4Z5Ta)

- [Jumper wires for easy hookup](http://amzn.to/2qVhd4y)

- [Resistor pack](http://amzn.to/2Dmainw)

  - You will be using a 100 Ohms resistor, which is the resistor colored with brown-black-brown

- [LED](http://amzn.to/2Ex2v5q)

### Software

- [Python3](https://www.python.org/download/releases/3.0/) or [Golang](https://golang.org/dl/)

- [Viam Server](https://www.viam.com/)

- [Viam Python SDK](https://python.viam.dev/) or [Viam Golang SDK](https://pkg.go.dev/go.viam.com/rdk/robot/client#section-readme)

## Installation

1) [Install Viam Server on your Raspberry Pi](https://docs.viam.com/getting-started/rpi-setup/).

2) Clone this repo onto your development machine:

```bash
git clone git@github.com:viam-labs/LED-Blink.git
```

### Python

1) [Install Viam Python SDK on your computer](https://python.viam.dev/).

2) Update the payload and address information in `blink.py`. You can your information from the **`CONNECT`** tab of the [Viam App](https://app.viam.com/robots).

3) Run your code!

```bash
python3 python/blink.py
```

### Go

1) [Install Viam Go SDK on your computer](https://pkg.go.dev/go.viam.com/rdk/robot/client#section-readme).

2) Update the payload and address information in `blink.go`. You can your information from the **`CONNECT`** tab of the [Viam app](https://app.viam.com/robots).

3) Run your code!

```bash
go run go/blink.go
```

### Contributing TLDR

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

### License

#### [Apache 2.0](./LICENSE)
