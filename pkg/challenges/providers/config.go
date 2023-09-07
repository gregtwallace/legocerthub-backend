package providers

import (
	"legocerthub-backend/pkg/challenges/providers/dns01acmedns"
	"legocerthub-backend/pkg/challenges/providers/dns01acmesh"
	"legocerthub-backend/pkg/challenges/providers/dns01cloudflare"
	"legocerthub-backend/pkg/challenges/providers/dns01manual"
	"legocerthub-backend/pkg/challenges/providers/http01internal"
)

// Config contains configurations for all provider types
type Config struct {
	Http01InternalConfigs  []*http01internal.Config  `yaml:"http_01_internal" json:"http_01_internal"`
	Dns01ManualConfigs     []*dns01manual.Config     `yaml:"dns_01_manual" json:"dns_01_manual"`
	Dns01AcmeDnsConfigs    []*dns01acmedns.Config    `yaml:"dns_01_acme_dns" json:"dns_01_acme_dns"`
	Dns01AcmeShConfigs     []*dns01acmesh.Config     `yaml:"dns_01_acme_sh" json:"dns_01_acme_sh"`
	Dns01CloudflareConfigs []*dns01cloudflare.Config `yaml:"dns_01_cloudflare" json:"dns_01_cloudflare"`
}

// Len returns the total number of Provider Configs, regardless of type.
func (cfg Config) Len() int {
	return len(cfg.Dns01AcmeDnsConfigs) +
		len(cfg.Dns01AcmeShConfigs) +
		len(cfg.Dns01CloudflareConfigs) +
		len(cfg.Dns01ManualConfigs) +
		len(cfg.Http01InternalConfigs)
}

// All returns an array of all configs regardless of type
func (cfg Config) All() []providerConfig {
	var all []providerConfig
	for _, cfg := range cfg.Dns01AcmeDnsConfigs {
		all = append(all, cfg)
	}
	for _, cfg := range cfg.Dns01AcmeShConfigs {
		all = append(all, cfg)
	}
	for _, cfg := range cfg.Dns01CloudflareConfigs {
		all = append(all, cfg)
	}
	for _, cfg := range cfg.Dns01ManualConfigs {
		all = append(all, cfg)
	}
	for _, cfg := range cfg.Http01InternalConfigs {
		all = append(all, cfg)
	}

	return all
}
