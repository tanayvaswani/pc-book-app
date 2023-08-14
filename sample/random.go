package sample

import (
	"math/rand"

	"github.com/tanayvaswani/pc-book-app/pb"
)

func randonKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet("i3", "i5", "i7", "i9")
	}
	return randomStringFromSet("Ryzen3", "Ryzen5", "Ryzen7", "Ryzen9")
}

func getRandomNumber(min, max int) int {
	return rand.Intn(max - min + 1) + min
}