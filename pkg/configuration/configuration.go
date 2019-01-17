package configuration

func New() Configuration {
	config := Configuration{dict: map[string]string{}}
	return config
}

func (c Configuration) Get(key string) string {
	return c.dict[key]
}

func (c Configuration) Set(key string, value string) {
	c.dict[key] = value
}

type Configuration struct {
	dict map[string]string
}
