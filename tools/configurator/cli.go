package configurator

import (
	"flag"
)

var (
	cmd      *bool
	broker   *string
	username *string
	password *string
)

func init() {
	cmd = flag.Bool("c", false, "turn args mode")
	broker = flag.String("b", "localhost", "broker ip")
	username = flag.String("u", "admin", "mqtt username")
	password = flag.String("p", "admin", "mqtt password")
}

func CmdInit() (*conf, bool) {
	flag.Parse()

	cfg := &conf{}
	cfg.Broker = *broker
	cfg.Username = *username
	cfg.Password = *password

	return cfg, *cmd
}
