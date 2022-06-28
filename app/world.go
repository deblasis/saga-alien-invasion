package app

type World struct {
	config *Config

	Map *Map
}

func NewWorld(config *Config) *World {
	return &World{
		config: config,
	}
}

func (w *World) Setup() error {
	//TODO: parseMapFile
	//TODO: spawnAliens
	return nil
}

func (w *World) Spin() error {
	//TODO: loop
	//  TODO: handleBattles
	//  TODO: moveAliens
	//TODO: outputMap

	return nil
}
