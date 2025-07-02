package types

type ParsedPSResult struct {
	PID          int
	InternalPort int
	ExternalPort int
}

type CreateTunnelDTO struct {
	Name         string `json:"name" validate:"required"`
	InternalPort int    `json:"internalPort" validate:"required,gt=0"`
	ExternalPort int    `json:"externalPort" validate:"required,gt=0"`
}

type KillTunnelDTO struct {
	PID int `json:"pid" validate:"required,gt=0"`
}
