package main

import "fmt"

// client.go: 客户端代码
type Client struct {
}

func (c *Client) InsertLightingConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lighting connector into computer.")
	com.InsertIntoLightingPort()
}

// computer.go: 客户端接口
type Computer interface {
	InsertIntoLightingPort()
}

// mac.go: 服务
type Mac struct {
}

func (m *Mac) InsertIntoLightingPort() {
	fmt.Println("Lighting connector is plugged into mac machine.")
}

// windows.go: 未知服务
type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// windowsAdapter.go: 适配器
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lighting signal to USB.")
	w.windowMachine.InsertIntoUSBPort()
}

// main.go
func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightingConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowMachine: windowsMachine}

	client.InsertLightingConnectorIntoComputer(windowsMachineAdapter)
}
