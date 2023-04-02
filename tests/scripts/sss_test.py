import RPi.GPIO as GPIO
import time

GPIO.setmode(GPIO.BCM)
TRIG = 21
ECHO = 26

GPIO.setup(TRIG, GPIO.OUT)
GPIO.output(TRIG,0)

GPIO.setup(ECHO, GPIO.IN)
time.sleep(0.1)

print("Starting measurement...")

GPIO.output(TRIG,1)
time.sleep(0.00001)
GPIO.output(TRIG,0)

while GPIO.input(ECHO)==0:
	pass
start = time.time()

while GPIO.input(ECHO)== 1:
	pass
stop = time.time()
distance = (stop-start)*17000
print(distance)
