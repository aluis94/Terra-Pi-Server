import time
import requests
import Adafruit_DHT

url = "http://192.168.1.112:8080/data-entry/add"
PIN1 = 21
PIN2 = 2
PIN3 = 0
deviceID = 1
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
	time.sleep(3);
myobj = {"DataEntry": {"Device_ID": deviceID,"Type": deviceType ,"Value": humidity,"Unit": "Percent"}}
post_response = requests.post(url, json = myobj)
post_response_json = post_response.json()


#Save data entry to DB
print("Saving to DB")
myobj = {"DataEntry": {"Device_ID": deviceID,"Type": deviceType ,"Value": value,"Unit": unit}}
post_response = requests.post(url, json = myobj)
post_response_json = post_response.json()
print(post_response_json)
