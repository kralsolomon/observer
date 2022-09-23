package main

func main() {
	var WeatherStation = NewSubject("Weather staion", 21, 31) //observable item
	var Mobile = NewListener("Xiaomi")                        //observer
	var Computer = NewListener("PC")
	WeatherStation.Subscribe(*Mobile)
	WeatherStation.SetPressure(35)
	WeatherStation.Unsubscribe(*Mobile)
	WeatherStation.Subscribe(*Computer)
	WeatherStation.SetTemperature(58)
}
