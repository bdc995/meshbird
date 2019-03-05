package config

type Config struct {
	Key              string   `default:""`
	SeedAddrs        []string `default:""`
	HostAddr         string   `default:""`
	PublicAddrs      []string `default:""`
	BindAddrs        []string `default:""`
	TransportThreads int      `default:"1"`
	Ip               string   `default:"10.237.0.1/16"`
	Mtu              int      `default:"9000"`
	Verbose          int      `default:"0"`
}
