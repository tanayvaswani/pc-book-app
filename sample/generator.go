package sample

import "github.com/tanayvaswani/pc-book-app/pb"

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout: randonKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCpu returns a new sample CPU
func NewCpu() *pb.CPU {
	brand := randomCPUBrand()
	name := randomName(brand)
	cores := getRandomNumber(2, 8)
	threads := getRandomNumber(cores, 12)

	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		Cores: uint32(cores),
		Threads: uint32(threads),

	}
	return cpu
}