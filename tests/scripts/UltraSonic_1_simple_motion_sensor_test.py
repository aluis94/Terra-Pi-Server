import time
import requests
import RPi.GPIO as GPIO

url = "http://192.168.1.112:8080/data-entry/add"
PIN1 = 26
PIN2 = 21
PIN3 = 0
deviceID = 1
deviceType ="Motion"
unit = ""
value = ""

#Read motion sensor distance
Unit ="Motion"
GPIO.setmode(GPIO.BCM)
unit = "CM"
TRIG = PIN1
ECHO = PIN2
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
value = distance


#Save data entry to DB
print("Saving to DB")
myobj = {"DataEntry": {"Device_ID": deviceID,"Type": deviceType ,"Value": value,"Unit": unit}}
post_response = requests.post(url, json = myobj)
post_response_json = post_response.json()
print(post_response_json)
