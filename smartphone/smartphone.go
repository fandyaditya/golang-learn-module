package smartphone

import "fmt"

// Removing Infrared because new generation phone not using infrared technology again
type SmartPhone interface {
	Nfc() bool
	Bluetooth() bool
}

type Samsung struct {
	ModelName string
}

type Xiaomi struct {
	ModelName string
}

func (s Samsung) Bluetooth() bool {
	fmt.Println("Bluetooth for " + s.ModelName + "activated")
	return true
}

func (s Samsung) Nfc() bool {
	fmt.Println("Nfc for " + s.ModelName + "activated")
	return true
}

func (x Xiaomi) Bluetooth() bool {
	fmt.Println("Bluetooth for " + x.ModelName + "activated")
	return true
}

func (x Xiaomi) Nfc() bool {
	fmt.Println("Nfc not available in " + x.ModelName)
	return false
}
