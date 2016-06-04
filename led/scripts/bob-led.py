#!/usr/bin/python

import time
import RPi.GPIO as GPIO
import sys

leds = {}
leds["blue"] = 6
leds["red"] = 12
leds["green"] = 20
leds["yellow"] = 21

mode = sys.argv[1]
led = leds[sys.argv[2]]

GPIO.setmode(GPIO.BCM)
GPIO.setup(led, GPIO.OUT)

print(mode)
print(led)

if (mode == "on"):
  print("on!")
  GPIO.output(led, True)
else:
  print("off!")
  GPIO.output(led, False)
  GPIO.cleanup()
