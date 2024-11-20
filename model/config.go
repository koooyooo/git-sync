package model

type Config struct {
	Dirs []Dir `yaml:"dirs"`
}

type Dir struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}
