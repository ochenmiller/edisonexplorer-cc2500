#!/bin/bash
# GPIO for chip select GP20_I2C_1_SDA P2_1
echo 20 > /sys/class/gpio/export
echo out > /sys/class/gpio/gpio20/direction
echo mode0 > /sys/kernel/debug/gpio_debug/gpio20/current_pinmux
echo -n "1" > /sys/class/gpio/gpio20/value
# GPIO for interrupt GP19_I2C_1_SCL P2_2
echo 19 > /sys/class/gpio/export
echo in > /sys/class/gpio/gpio19/direction
echo mode0 > /sys/kernel/debug/gpio_debug/gpio19/current_pinmux
