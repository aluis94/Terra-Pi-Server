import RPi.GPIO as GPIO
import time
GPIO.setmode(GPIO.BCM)
LD_PIN = 26
GPIO.setup(LD_PIN,GPIO.OUT)

while (True):
	LD_IN =  GPIO.input(LD_PIN)
	print(LD_IN)
	if LD_IN == 1:
		print("dark")
	else:
		print("light")
	time.sleep(3)
