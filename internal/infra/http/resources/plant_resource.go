package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type PlantDto struct {
	Id      uint64           `json:"id"`
	UserId  uint64           `json:"userId"`
	Name    string           `json:"name"`
	Address string           `json:"address"`
	Lat     float64          `json:"lat"`
	Lon     float64          `json:"lon"`
	Type    domain.PlantType `json:"type"`
}

type PlantsDto struct {
	Plants []PlantDto `json:"plants"`
}

func (d PlantDto) DomainToDto(plant domain.Plant) PlantDto {
	return PlantDto{
		Id:      plant.Id,
		UserId:  plant.UserId,
		Name:    plant.Name,
		Address: plant.Address,
		Lat:     plant.Lat,
		Lon:     plant.Lon,
		Type:    plant.Type,
	}
}

func (d PlantsDto) DomainToDtoCollection(ps []domain.Plant) PlantsDto {
	var plants []PlantDto
	for _, p := range ps {
		pDto := PlantDto{}.DomainToDto(p)
		plants = append(plants, pDto)
	}
	return PlantsDto{Plants: plants}
}
