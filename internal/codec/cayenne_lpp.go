package codec

import (
	"bytes"
	"encoding/gob"
	"github.com/pkg/errors"
	"github.com/selvakn/go-cayenne-lib/cayennelpp"
)

func init() {
	gob.Register(CayenneLPP{})
}

// Accelerometer defines the accelerometer data.
type Accelerometer struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// Gyrometer defines the gyrometer data.
type Gyrometer struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// GPSLocation defines the GPS location data.
type GPSLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
}

// CayenneLPP defines the Cayenne LPP data structure.
type CayenneLPP struct {
	DigitalInput      map[byte]uint8         `json:"digitalInput,omitempty" influxdb:"digital_input"`
	DigitalOutput     map[byte]uint8         `json:"digitalOutput,omitempty" influxdb:"digital_output"`
	AnalogInput       map[byte]float32       `json:"analogInput,omitempty" influxdb:"analog_input"`
	AnalogOutput      map[byte]float32       `json:"analogOutput,omitempty" influxdb:"analog_output"`
	IlluminanceSensor map[byte]uint16        `json:"illuminanceSensor,omitempty" influxdb:"illuminance_sensor"`
	PresenceSensor    map[byte]uint8         `json:"presenceSensor,omitempty" influxdb:"presence_sensor"`
	TemperatureSensor map[byte]float32       `json:"temperatureSensor,omitempty" influxdb:"temperature_sensor"`
	HumiditySensor    map[byte]float32       `json:"humiditySensor,omitempty" influxdb:"humidity_sensor"`
	Accelerometer     map[byte]Accelerometer `json:"accelerometer,omitempty" influxdb:"accelerometer"`
	Barometer         map[byte]float32       `json:"barometer,omitempty" influxdb:"barometer"`
	Gyrometer         map[byte]Gyrometer     `json:"gyrometer,omitempty" influxdb:"gyrometer"`
	GPSLocation       map[byte]GPSLocation   `json:"gpsLocation,omitempty" influxdb:"gps_location"`
	Voltage           map[byte]float32       `json:"voltage,omitempty" influxdb:"voltage"`
	Current           map[byte]float32       `json:"current,omitempty" influxdb:"current"`
	Frequency         map[byte]float32       `json:"frequency,omitempty" influxdb:"frequency"`
	Energy            map[byte]float32       `json:"energy,omitempty" influxdb:"energy"`
}

type target struct {
	values *CayenneLPP
}

// Object returns the CayenneLPP data object.
func (c CayenneLPP) Object() interface{} {
	return c
}

func (t *target) Port(channel uint8, value float32) {
}

func (t *target) DigitalInput(channel, value uint8) {
	if t.values.DigitalInput == nil {
		t.values.DigitalInput = make(map[uint8]uint8)
	}
	t.values.DigitalInput[channel] = value
}

func (t *target) DigitalOutput(channel, value uint8) {
	if t.values.DigitalOutput == nil {
		t.values.DigitalOutput = make(map[uint8]uint8)
	}
	t.values.DigitalOutput[channel] = value
}

func (t *target) AnalogInput(channel uint8, value float32) {
	if t.values.AnalogInput == nil {
		t.values.AnalogInput = make(map[uint8]float32)
	}
	t.values.AnalogInput[channel] = value
}

func (t *target) AnalogOutput(channel uint8, value float32) {
	if t.values.AnalogOutput == nil {
		t.values.AnalogOutput = make(map[uint8]float32)
	}
	t.values.AnalogOutput[channel] = value
}

func (t *target) Luminosity(channel uint8, value uint16) {
	if t.values.IlluminanceSensor == nil {
		t.values.IlluminanceSensor = make(map[uint8]uint16)
	}
	t.values.IlluminanceSensor[channel] = value
}

func (t *target) Presence(channel, value uint8) {
	if t.values.PresenceSensor == nil {
		t.values.PresenceSensor = make(map[uint8]uint8)
	}
	t.values.PresenceSensor[channel] = value
}

func (t *target) Temperature(channel uint8, celcius float32) {
	if t.values.TemperatureSensor == nil {
		t.values.TemperatureSensor = make(map[uint8]float32)
	}
	t.values.TemperatureSensor[channel] = celcius
}

func (t *target) RelativeHumidity(channel uint8, rh float32) {
	if t.values.HumiditySensor == nil {
		t.values.HumiditySensor = make(map[uint8]float32)
	}
	t.values.HumiditySensor[channel] = rh
}

func (t *target) Accelerometer(channel uint8, x, y, z float32) {
	if t.values.Accelerometer == nil {
		t.values.Accelerometer = make(map[uint8]Accelerometer)
	}
	t.values.Accelerometer[channel] = Accelerometer{X: x, Y: y, Z: z}
}

func (t *target) BarometricPressure(channel uint8, hpa float32) {
	if t.values.Barometer == nil {
		t.values.Barometer = make(map[uint8]float32)
	}
	t.values.Barometer[channel] = hpa
}

func (t *target) Gyrometer(channel uint8, x, y, z float32) {
	if t.values.Gyrometer == nil {
		t.values.Gyrometer = make(map[uint8]Gyrometer)
	}
	t.values.Gyrometer[channel] = Gyrometer{X: x, Y: y, Z: z}
}

func (t *target) GPS(channel uint8, latitude, longitude, altitude float32) {
	if t.values.GPSLocation == nil {
		t.values.GPSLocation = make(map[uint8]GPSLocation)
	}
	t.values.GPSLocation[channel] = GPSLocation{Latitude: latitude, Longitude: longitude, Altitude: altitude}
}

func (t *target) Voltage(channel uint8, value float32) {
	if t.values.Voltage == nil {
		t.values.Voltage = make(map[uint8]float32)
	}
	t.values.Voltage[channel] = value
}

func (t *target) Current(channel uint8, value float32) {
	if t.values.Current == nil {
		t.values.Current = make(map[uint8]float32)
	}
	t.values.Current[channel] = value
}

func (t *target) Frequency(channel uint8, value float32) {
	if t.values.Frequency == nil {
		t.values.Frequency = make(map[uint8]float32)
	}
	t.values.Frequency[channel] = value
}

func (t *target) Energy(channel uint8, value float32) {
	if t.values.Energy == nil {
		t.values.Energy = make(map[uint8]float32)
	}
	t.values.Energy[channel] = value
}

// DecodeBytes decodes the payload from a slice of bytes.
func (c *CayenneLPP) DecodeBytes(data []byte) error {
	decoder := cayennelpp.NewDecoder(bytes.NewBuffer(data))
	target := &target{c}
	err := decoder.DecodeUplink(target)

	if err != nil {
		return errors.Wrap(err, "decode error")
	}

	return nil
}

// EncodeToBytes encodes the payload to a slice of bytes.
func (c CayenneLPP) EncodeToBytes() ([]byte, error) {
	w := bytes.NewBuffer([]byte{})

	encoder := cayennelpp.NewEncoder()

	for k, v := range c.DigitalInput {
		encoder.AddDigitalInput(k, v)
	}
	for k, v := range c.DigitalOutput {
		encoder.AddDigitalOutput(k, v)
	}
	for k, v := range c.AnalogInput {
		encoder.AddAnalogInput(k, v)
	}
	for k, v := range c.AnalogOutput {
		encoder.AddAnalogOutput(k, v)
	}
	for k, v := range c.IlluminanceSensor {
		encoder.AddLuminosity(k, v)
	}
	for k, v := range c.PresenceSensor {
		encoder.AddPresence(k, v)
	}
	for k, v := range c.TemperatureSensor {
		encoder.AddTemperature(k, v)
	}
	for k, v := range c.HumiditySensor {
		encoder.AddRelativeHumidity(k, v)
	}
	for k, v := range c.Accelerometer {
		encoder.AddAccelerometer(k, v.X, v.Y, v.Z)
	}
	for k, v := range c.Barometer {
		encoder.AddBarometricPressure(k, v)
	}
	for k, v := range c.Gyrometer {
		encoder.AddGyrometer(k, v.X, v.Y, v.Z)
	}
	for k, v := range c.GPSLocation {
		encoder.AddGPS(k, v.Latitude, v.Longitude, v.Altitude)
	}
	for k, v := range c.Voltage {
		encoder.AddVoltage(k, v)
	}
	for k, v := range c.Current {
		encoder.AddCurrent(k, v)
	}
	for k, v := range c.Energy {
		encoder.AddEnergy(k, v)
	}
	for k, v := range c.Frequency {
		encoder.AddFrequency(k, v)
	}

	return w.Bytes(), nil
}
