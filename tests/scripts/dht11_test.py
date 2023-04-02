import Adafruit_DHT
import time

DHT_SENSOR = Adafruit_DHT.DHT11
DHT_PIN = 21

while True:
	humidity, temperature = Adafruit_DHT.read(DHT_SENSOR,DHT_PIN)
	if humidity is not None and temperature is not None:
		temp_f = ((temperature*(9/5))+32)
		print("Temp={0:0.1f}C Humidity={1:0.1f}%".format(temp_f,humidity))
	else:
		print("Sensor failure. Check wiring.")
	time.sleep(3);

