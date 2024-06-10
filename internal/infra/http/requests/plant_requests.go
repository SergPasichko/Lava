package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type AddPlantRequest struct {
	Name    string           `json:"name" validate:"required"`
	Address string           `json:"address" validate:"required"`
	Lat     float64          `json:"lat" validate:"required"`
	Lon     float64          `json:"lon" validate:"required"`
	Type    domain.PlantType `json:"type" validate:"required,oneof=SOLAR WIND"`
}

func (r AddPlantRequest) ToDomainModel() (interface{}, error) {
	return domain.Plant{
		Name:    r.Name,
		Address: r.Address,
		Lat:     r.Lat,
		Lon:     r.Lon,
		Type:    r.Type,
	}, nil
}
