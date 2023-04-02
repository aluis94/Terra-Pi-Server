#!/usr/bin/python
import RPi.GPIO as GPIO
import time as t
import datetime

##Device On script for Light1
#Set device pin
DPIN= 27
#Use Braodcom pinout numbering
GPIO.setmode(GPIO.BCM)
#Setup outputs
GPIO.setup(DPIN, GPIO.OUT)
#turn off pin
GPIO.output(DPIN, GPIO.LOW)
