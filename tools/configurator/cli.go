package configurator

import (
	"flag"
)

var (
	input    bool
	nodebug  bool
	broker   string
	username string
	password string
)

func init() {
	flag.BoolVar(&input, "c", false, "turn args mode")
	flag.BoolVar(&input, "config", false, "turn args mode")
	flag.BoolVar(&nodebug, "n", false, "debug")
	flag.BoolVar(&nodebug, "nodebug", false, "debug")
	flag.StringVar(&broker, "b", "localhost", "broker ip")
	flag.StringVar(&broker, "broker", "localhost", "broker ip")
	flag.StringVar(&username, "u", "admin", "mqtt username")
	flag.StringVar(&username, "user", "admin", "mqtt username")
	flag.StringVar(&password, "p", "admin", "mqtt password")
	flag.StringVar(&password, "pass", "admin", "mqtt password")
}

func ArgsInit() (*conf, bool, bool) {
	flag.Parse()

	cfg := &conf{}
	cfg.Broker = broker
	cfg.Username = username
	cfg.Password = password

	return cfg, input, !nodebug
}
