package resolvers

import "github.com/ajaka-the-wizard/dnsgo/internal/config"

type Resolvers struct {
	zones *config.Zones
	roots *config.Roots
}

func CreateResolver(z *config.Zones, r *config.Roots) *Resolvers {
	return &Resolvers{
		zones: z,
		roots: r,
	}
}
