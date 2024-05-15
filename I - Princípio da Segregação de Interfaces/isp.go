package isp

import "fmt"

/*

I - Princípio da Segregação de Interfaces (Interface Segregation Principle)

Clientes não devem ser forçados a depender de métodos que não usam.
Devemos criar interfaces mais específicas do que uma genérica.
*/

// Device interface exige que todos os dispositivos implementem todos esses métodos.
type Device interface {
	TurnOn()
	TurnOff()
	SetTimer(int)
	ConnectToBluetooth()
}

// Lamp é um dispositivo simples que apenas liga e desliga.
type Lamp struct{}

func (l Lamp) TurnOn() {
	fmt.Println("Lamp turned on")
}

func (l Lamp) TurnOff() {
	fmt.Println("Lamp turned off")
}

// SetTimer e ConnectToBluetooth são métodos desnecessários para uma lâmpada.
func (l Lamp) SetTimer(t int) {
	fmt.Println("Setting timer not supported")
}

func (l Lamp) ConnectToBluetooth() {
	fmt.Println("Bluetooth not supported")
}

/*
	Para resolver esse problema e seguir o ISP,
	podemos dividir a interface Device em interfaces
	 menores e mais específicas que melhor correspondam
	 às capacidades de cada dispositivo
*/

// BasicDevice apenas para dispositivos que ligam e desligam.
type BasicDevice interface {
	TurnOn()
	TurnOff()
}

// ProgrammableDevice para dispositivos que podem ser programados com um timer.
type ProgrammableDevice interface {
	BasicDevice
	SetTimer(int)
}

// BluetoothCapable para dispositivos que suportam conexão Bluetooth.
type BluetoothCapable interface {
	BasicDevice
	ConnectToBluetooth()
}

// Lamp agora implementa apenas a interface BasicDevice.
type NLamp struct{}

func (l NLamp) TurnOn() {
	fmt.Println("Lamp turned on")
}

func (l NLamp) TurnOff() {
	fmt.Println("Lamp turned off")
}

// SmartSpeaker implementa BasicDevice e BluetoothCapable.
type SmartSpeaker struct{}

func (s SmartSpeaker) TurnOn() {
	fmt.Println("Speaker turned on")
}

func (s SmartSpeaker) TurnOff() {
	fmt.Println("Speaker turned off")
}

func (s SmartSpeaker) ConnectToBluetooth() {
	fmt.Println("Connecting to Bluetooth")
}

/*
func main() {
	var lamp BasicDevice = NLamp{}
	lamp.TurnOn()
	lamp.TurnOff()

	var speaker BluetoothCapable = SmartSpeaker{}
	speaker.TurnOn()
	speaker.TurnOff()
	speaker.ConnectToBluetooth()
}
*/
