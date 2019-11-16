package main

//Packet - basic struct of captured mqtt package
type Packet struct {
	ID         int
	TimeRel    float32
	IPSrc      string
	IPDest     string
	PortSrc    string
	PortDest   string
	PacketType string
}
