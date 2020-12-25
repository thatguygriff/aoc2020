package twentyfive

import "fmt"

type key struct {
	public   int
	sub      int
	loopsize int
}

func (k *key) generatePublic() {
	v := 1

	for i := 0; i < k.loopsize; i++ {
		v *= k.sub
		v %= 20201227
	}

	k.public = v
}

func (k *key) encryption(public int) int {
	v := 1

	for i := 0; i < k.loopsize; i++ {
		v *= public
		v %= 20201227
	}

	return v
}

func (k *key) determineLoopSize() {
	if k.public == 0 {
		return
	}

	v := 1
	loopsize := 0
	for v != k.public {
		v *= k.sub
		v %= 20201227
		loopsize++
	}

	k.loopsize = loopsize
}

// PartOne What encryption key is the handshake trying to establish?
func PartOne() string {
	card := key{
		public: 2069194,
		sub:    7,
	}

	door := key{
		public: 16426071,
		sub:    7,
	}

	card.determineLoopSize()
	door.determineLoopSize()

	return fmt.Sprintf("The handshake is trying to establish the key %d", door.encryption(card.public))
}
