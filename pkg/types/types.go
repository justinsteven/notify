package types

import "github.com/projectdiscovery/goflags"

type Options struct {
	BIID string `yaml:"burp_biid,omitempty"`

	Verbose        bool                `yaml:"verbose,omitempty"`
	NoColor        bool                `yaml:"no_color,omitempty"`
	Silent         bool                `yaml:"silent,omitempty"`
	Version        bool                `yaml:"version,omitempty"`
	Interval       int                 `yaml:"interval,omitempty"`
	ProviderConfig string              `yaml:"provider_config,omitempty"`
	Providers      goflags.StringSlice `yaml:"providers,omitempty"`
	Profiles       goflags.StringSlice `yaml:"profiles,omitempty"`

	HTTPMessage string `yaml:"http_message,omitempty"`
	DNSMessage  string `yaml:"dns_message,omitempty"`
	CLIMessage  string `yaml:"cli_message,omitempty"`
	SMTPMessage string `yaml:"smtp_message,omitempty"`

	Stdin bool
	Data  string `yaml:"data,omitempty"`
}