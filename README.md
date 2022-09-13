# LED-Blink

In this post, you will be introduced to the basics of programming hardware using the [Viam Python SDK](https://python.viam.dev/) by making an LED blink. This will allow you to write Python code to make an LED connected to the GPIO of a Raspberry Pi blink. This tutorial is a good introduction to Python programming, and developing code for hardware like robots and IoT devices.

![A GIF of the completed project showing a blinking blue LED connected to a Raspberry Pi with jumper cables.](https://user-images.githubusercontent.com/4650739/190014115-78f89892-ee67-4a19-846a-33c86e8e6272.gif)

## Circuit Diagram

For reference, the circuit you are building for this tutorial looks like this:

![Circuit diagram showing a Raspberry Pi with a red connector running out of GPIO pin 8 to a 100-ohm* resistor*. The resistor is connected to the long lead of a red LED bulb. Finally, a blue connector connects the short lead of the LED to the ground connection on pin 6 of the Raspberry Pi GPIO pins.](https://user-images.githubusercontent.com/4650739/190014128-1949ef55-47b0-4706-8d6f-2aac99bc6fd7.png)

## Prerequisites

You will need the following hardware, tools, and software to complete this project:

### Hardware:

-   [Raspberry Pi 3 or 4](https://a.co/d/5Tn67G3)

    -   [Check out the Viam Raspberry Pi Setup Guide for steps on how to get started](https://docs.viam.com/getting-started/installation/)

    -   [Be sure that you have set up Viam Server on your Raspberry Pi as well.](https://docs.viam.com/getting-started/installation/#installing-viam-server)

    -   [Be sure that you are running Raspbian on your Pi.](https://docs.viam.com/getting-started/installation/#installing-raspian-on-the-raspberry-pi)

    -   SSH must also be enabled on your Pi.

-   [Solderless breadboard](https://amzn.to/2Q4Z5Ta)

-   [Jumper wires for easy hookup](http://amzn.to/2qVhd4y)

-   [Resistor pack](http://amzn.to/2Dmainw)

    -   You will be using a 100 Ohms resistor, which is the resistor colored with brown-black-brown

-   [Blue LED](http://amzn.to/2Ex2v5q)

-   [Multimeter](http://amzn.to/2qWurxS) (optional)

### Software:

-   [Python3](https://www.python.org/download/releases/3.0/)

-   [Pip](https://pip.pypa.io/en/stable/#)

-   [Viam Server](https://github.com/viamrobotics/rdk/tree/0c550c246739b87b4d5a9e8d96d2b6fdb3948e2b)

-   [Viam Python SDK](https://python.viam.dev/)

-   Install [Visual Studio Code](https://code.visualstudio.com/) or [Visual Studio Code Insiders](https://code.visualstudio.com/insiders/) on your development machine (not your Raspberry Pi).

## Installation

1) [Install Viam Server on your Raspberry Pi](https://docs.viam.com/getting-started/installation/#installing-viam-server).

2) [Install the Viam Python SDK](https://python.viam.dev/) on your development machine.

3) Clone this repo onto your development machine:

```bash
git clone https://github.com/viamrobotics/LED-Blink.git
```

4) Update the payload and address information in `blink.py`. You can your information from the **`Connect`** tab of the [Viam App](https://app.viam.com/robots).


## Contributing

Please read [CONTRIBUTING.md](https://github.com/JoeKarlsson/iot-kitty-litter-box/blob/develop/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Contributing TLDR;

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

### License

#### [Apache 2.0](./LICENSE)