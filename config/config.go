package config

import (
	"flag"
	"fmt"
	"time"
)

const (
	defaultAddress = ":8080"

	defaultTLSEnable = false
	defaultTLSCert   = ""
	defaultTLSKey    = ""

	defaultGroupsLimit = 100
	defaultConnsLimit  = 1000

	defaultLogEnableJSON = false
	defaultLogLevel      = "info"

	defaultEnableBroadcast = false
	defaultEnableWeb       = false
)

const (
	flagLabelAddress = "address"

	flagLabelTLSEnable = "tls-enable"
	flagLabelTLSCert   = "tls-cert"
	flagLabelTLSKey    = "tls-key"

	flagLabelGroupsLimit = "groups-limit"
	flagLabelConnsLimit  = "conns-limit"

	flagLabelSeed = "seed"

	flagLabelLogEnableJSON = "log-json"
	flagLabelLogLevel      = "log-level"

	flagLabelEnableBroadcast = "enable-broadcast"
	flagLabelEnableWeb       = "enable-web"
)

const (
	flagUsageAddress = "address to serve"

	flagUsageTLSEnable = "enable TLS"
	flagUsageTLSCert   = "path to certificate file"
	flagUsageTLSKey    = "path to key file"

	flagUsageGroupsLimit = "game groups limit"
	flagUsageConnsLimit  = "web-socket connections limit"

	flagUsageSeed = "random seed"

	flagUsageLogEnableJSON = "use json format for logger"
	flagUsageLogLevel      = "set log level: panic, fatal, error, warning (warn), info or debug"

	flagUsageEnableBroadcast = "enable broadcasting API method"
	flagUsageEnableWeb       = "enable web client"
)

const envVarSnakeServerConfigPath = "SNAKE_SERVER_CONFIG_PATH"

func generateSeed() int64 {
	return time.Now().UnixNano()
}

type TLS struct {
	Enable bool   `yaml:"enable"`
	Cert   string `yaml:"cert"`
	Key    string `yaml:"key"`
}

type Limits struct {
	Groups int `yaml:"groups"`
	Conns  int `yaml:"conns"`
}

type Log struct {
	EnableJSON bool   `yaml:"enable_json"`
	Level      string `yaml:"level"`
}

type Server struct {
	Address string `yaml:"address"`

	TLS    `yaml:"tls"`
	Limits `yaml:"limits"`
	Seed   int64 `yaml:"seed"`
	Log    `yaml:"log"`

	EnableBroadcast bool `yaml:"enable_broadcast"`
	EnableWeb       bool `yaml:"enable_web"`
}

// Config is a server configuration structure
type Config struct {
	Server `yaml:"server"`
}

var defaultConfig = Config{
	Server: Server{
		Address: defaultAddress,

		TLS: TLS{
			Enable: defaultTLSEnable,
			Cert:   defaultTLSCert,
			Key:    defaultTLSKey,
		},

		Limits: Limits{
			Groups: defaultGroupsLimit,
			Conns:  defaultConnsLimit,
		},

		Seed: generateSeed(),

		Log: Log{
			EnableJSON: defaultLogEnableJSON,
			Level:      defaultLogLevel,
		},

		EnableBroadcast: defaultEnableBroadcast,
		EnableWeb:       defaultEnableWeb,
	},
}

func DefaultConfig() Config {
	return defaultConfig
}

func ParseFlags(fs *flag.FlagSet, args []string, defaults Config) (Config, error) {
	if fs.Parsed() {
		panic("program composition error: the provided FlagSet has been parsed")
	}

	config := defaults

	// Address
	fs.StringVar(&config.Server.Address, flagLabelAddress, defaults.Server.Address, flagUsageAddress)

	// TLS
	fs.BoolVar(&config.Server.TLS.Enable, flagLabelTLSEnable, defaults.Server.TLS.Enable, flagUsageTLSEnable)
	fs.StringVar(&config.Server.TLS.Cert, flagLabelTLSCert, defaults.Server.TLS.Cert, flagUsageTLSCert)
	fs.StringVar(&config.Server.TLS.Key, flagLabelTLSKey, defaults.Server.TLS.Key, flagUsageTLSKey)

	// Limits
	fs.IntVar(&config.Server.Limits.Groups, flagLabelGroupsLimit, defaults.Server.Limits.Groups, flagUsageGroupsLimit)
	fs.IntVar(&config.Server.Limits.Conns, flagLabelConnsLimit, defaults.Server.Limits.Conns, flagUsageConnsLimit)

	// Random
	fs.Int64Var(&config.Server.Seed, flagLabelSeed, defaults.Server.Seed, flagUsageSeed)

	// Logging
	fs.BoolVar(&config.Server.Log.EnableJSON, flagLabelLogEnableJSON, defaults.Server.Log.EnableJSON, flagUsageLogEnableJSON)
	fs.StringVar(&config.Server.Log.Level, flagLabelLogLevel, defaults.Server.Log.Level, flagUsageLogLevel)

	// Flags
	fs.BoolVar(&config.Server.EnableBroadcast, flagLabelEnableBroadcast, defaults.Server.EnableBroadcast, flagUsageEnableBroadcast)
	fs.BoolVar(&config.Server.EnableWeb, flagLabelEnableWeb, defaults.Server.EnableWeb, flagUsageEnableWeb)

	if err := fs.Parse(args); err != nil {
		return defaults, fmt.Errorf("cannot parse flags: %s", err)
	}

	return config, nil
}
