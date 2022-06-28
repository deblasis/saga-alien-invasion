package app

type World struct {
	config *Config
}

func NewWorld(config *Config) *World {
	return &World{
		config: config,
	}
}

func (w *World) Setup() error {
	return nil
}

func (w *World) Spin() error {
	return nil
}
