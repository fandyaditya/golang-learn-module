package smartphone

import "fmt"

type SmartPhone interface {
	Nfc() bool
	Bluetooth() bool
	Infrared() bool
}

type Samsung struct {
	modelName string
}

type Xiaomi struct {
	modelName string
}

func (s Samsung) Bluetooth() bool {
	fmt.Println("Bluetooth for " + s.modelName + "activated")
	return true
}

func (s Samsung) Nfc() bool {
	fmt.Println("Nfc for " + s.modelName + "activated")
	return true
}

func (s Samsung) Infrared() bool {
	fmt.Println("Infrared not available in " + s.modelName)
	return false
}

func (x Xiaomi) Bluetooth() bool {
	fmt.Println("Bluetooth for " + x.modelName + "activated")
	return true
}

func (x Xiaomi) Nfc() bool {
	fmt.Println("Nfc not available in " + x.modelName)
	return false
}

func (x Xiaomi) Infrared() bool {
	fmt.Println("Infrared not available in" + x.modelName)
	return false
}
