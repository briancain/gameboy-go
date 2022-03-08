package gbcore

type Sound struct {
	// Gameboy supports 4 sound channels
	Channel0 Channel
	Channel1 Channel
	Channel2 Channel
	Channel3 Channel

	Enabled bool

	// volume
	LeftVolume  int
	RightVolume int
}

type Channel struct {
}

func (s *Sound) Init() error {
	return nil
}

func (s *Sound) Play() {
}
