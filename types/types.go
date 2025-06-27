package types

type RouteTypeDTO struct {
	Subpath string `json:"subpath" validate:"required"`
	Port    int    `json:"port" validate:"required,gt=0"`
}

type RouteType struct {
	Subpath string `json:"subpath" validate:"required"`
	PID     int    `json:"pid" validate:"required,gt=0"`
}

type RoutesType map[int]RouteType
