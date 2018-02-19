package core

type IO interface {
	Read(port uint8) uint8
	Write(port uint8, value uint8)
}
