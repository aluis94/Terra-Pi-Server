import time
import requests
import RPi.GPIO as GPIO

url = "http://192.168.1.112:8080/data-entry/add"
PIN1 = 26
PIN2 = 21
PIN3 = 0
deviceID = 1
deviceType ="Light"
unit = ""
value = ""
#Light sensor 
unit ="Light/DARK"
didRead=0
GPIO.setmode(GPIO.BCM)
LD_PIN = PIN1
GPIO.setup(LD_PIN,GPIO.OUT)
while (not didRead):
	LD_IN =  GPIO.input(LD_PIN)
	print(LD_IN)
	if LD_IN == 1:
		print("dark")
		didRead = 1
	else:
		print("light")
		didRead = 1
	time.sleep(3)
value = LD_IN


#Save data entry to DB
print("Saving to DB")
myobj = {"DataEntry": {"Device_ID": deviceID,"Type": deviceType ,"Value": value,"Unit": unit}}
post_response = requests.post(url, json = myobj)
post_response_json = post_response.json()
print(post_response_json)
