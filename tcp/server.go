package tcp

import "time"

type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:""timeout`
}

var ClientCounter int32

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {

}
