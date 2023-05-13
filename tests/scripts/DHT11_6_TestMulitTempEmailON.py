
import time
import requests
import json
import smtplib
from appConfig import *
import ssl
import Adafruit_DHT
import RPi.GPIO as GPIO
import time as t
import datetime
#Message
smtp_port = 587                 # Standard secure SMTP port
smtp_server = "smtp.gmail.com"  # Google SMTP Server
message = "This is a test email- Multi"
#conditional Device

url = "http://192.168.1.112:8080/data-entry/add"
PIN1 = 26
PIN2 = 0
PIN3 = 0
deviceID = 6
deviceType ="Temp_Hum"
unit = ""
value = ""

#Temperature_Humidity sensor read
DHT_SENSOR = Adafruit_DHT.DHT11
DHT_PIN = PIN1
didRead = 0
temp_f = 0
unit = "F"
while not didRead:
	humidity, temperature = Adafruit_DHT.read(DHT_SENSOR,DHT_PIN)
	if humidity is not None and temperature is not None:
		temp_f = ((temperature*(9/5))+32)
		print("Temp={0:0.1f}C Humidity={1:0.1f}%".format(temp_f,humidity))
		didRead = 1
		value = temp_f
	else:
		print("Sensor failure. Check wiring.")
	time.sleep(3)
myobj = {"DataEntry": {"Device_ID": deviceID,"Type": deviceType ,"Value": humidity,"Unit": "Percent"}}
post_response = requests.post(url, json = myobj)
post_response_json = post_response.json()


#Save data entry to DB
print("Saving to DB")
myobj = {"DataEntry": {"Device_ID": deviceID,"Type": deviceType ,"Value": value,"Unit": unit}}
post_response = requests.post(url, json = myobj)
post_response_json = post_response.json()
print(post_response_json)


#condition
conditionString = {"Type":"Temperature","Operator":">","Value":"-70"}
cdevice_ID = 7
ndevice_ID = 1
cdevice_on_off = "ON"
condition = json.loads(conditionString)
msgType = "Email"
type = condition["Type"]
operator = condition["Operator"]
cvalue = condition["Value"]

isConditionMet = False
if (operator == ">"):
	isConditionMet = value > cvalue
elif (operator == "<"):
	isConditionMet = value > cvalue
elif (operator == "=="):
	isConditionMet = value == cvalue
elif (operator == "!="):
	isConditionMet = value != cvalue
elif (operator == ">="):
	isConditionMet = value >= cvalue
elif (operator == "<="):
	isConditionMet = value <= cvalue
else:
	isConditionMet = False

#If condition met, run jobs

if (isConditionMet):
	if (cdevice_ID != 0):
		if(cdevice_on_off.lower() == "off"):
			#Set device pin
			DPIN= 17
			#Use Braodcom pinout numbering
			GPIO.setmode(GPIO.BCM)
			#Setup outputs
			GPIO.setup(DPIN, GPIO.OUT)
			#turn off pin
			GPIO.output(DPIN, GPIO.HIGH)
		elif (cdevice_on_off.lower() == "on"):
			#Set device pin
			DPIN= 26
			#Use Braodcom pinout numbering
			GPIO.setmode(GPIO.BCM)
			#Setup outputs
			GPIO.setup(DPIN, GPIO.OUT)
			#turn off pin
			GPIO.output(DPIN, GPIO.LOW)
	
	if (ndevice_ID!=0):
		if (msgType.lower() =="email"):
			# Create context
			simple_email_context = ssl.create_default_context()
			try:
				# Connect to the server
				print("Connecting to server...")
				TIE_server = smtplib.SMTP(smtp_server, smtp_port)
				TIE_server.starttls(context=simple_email_context)
				TIE_server.login(email_from, pswd)
				print("Connected to server :-)")
				# Send the actual email
				print(f"Sending email to - {email_to}")
				TIE_server.sendmail(email_from, email_to, message)
				print(f"Email successfully sent to - {email_to}")
				
			except Exception as e:
				# If there's an error, print it out
				print(e)

			
			finally:
				# Close the port
				TIE_server.quit()
		
		if (msgType.lower() =="sms"):
			from twilio.rest import Client
			client = Client(account_sid, auth_token)
			message = client.api.messages.create(
    			to="+13109772732",
    			from_="+18665254691",
    			body="This is a test email- Multi"

			)
		

	



