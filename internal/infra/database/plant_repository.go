package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const PlantsTableName = "plants"

type plant struct {
	Id          uint64           `db:"id,omitempty"`
	UserId      uint64           `db:"user_id"`
	Name        string           `db:"name"`
	Address     string           `db:"address"`
	Lat         float64          `db:"lat"`
	Lon         float64          `db:"lon"`
	Type        domain.PlantType `db:"type"`
	CreatedDate time.Time        `db:"created_date"`
	UpdatedDate time.Time        `db:"updated_date"`
	DeletedDate *time.Time       `db:"deleted_date"`
}

type PlantRepository interface {
	Save(p domain.Plant) (domain.Plant, error)
	GetForUser(uId uint64) ([]domain.Plant, error)
	GetById(id uint64) (domain.Plant, error)
	Update(p domain.Plant) (domain.Plant, error)
	Delete(id uint64) error
}

type plantRepository struct {
	coll db.Collection
	sess db.Session
}

func NewPlantRepository(session db.Session) PlantRepository {
	return plantRepository{
		coll: session.Collection(PlantsTableName),
		sess: session,
	}
}

func (r plantRepository) Save(p domain.Plant) (domain.Plant, error) {
	pl := r.mapDomainToModel(p)
	pl.CreatedDate = time.Now()
	pl.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&pl)
	if err != nil {
		return domain.Plant{}, err
	}
	p = r.mapModelToDomain(pl)
	return p, err
}

func (r plantRepository) GetForUser(uId uint64) ([]domain.Plant, error) {
	var plants []plant
	err := r.coll.
		Find(db.Cond{"user_id": uId, "deleted_date": nil}).
		OrderBy("-updated_date").
		All(&plants)
	if err != nil {
		return nil, err
	}

	result := r.mapModelToDomainCollection(plants)
	return result, nil
}

func (r plantRepository) GetById(id uint64) (domain.Plant, error) {
	var pl plant
	err := r.coll.
		Find(db.Cond{"id": id, "deleted_date": nil}).
		One(&pl)
	if err != nil {
		return domain.Plant{}, err
	}

	result := r.mapModelToDomain(pl)
	return result, nil
}

func (r plantRepository) Update(p domain.Plant) (domain.Plant, error) {
	pl := r.mapDomainToModel(p)
	pl.UpdatedDate = time.Now()
	err := r.coll.
		Find(db.Cond{"id": pl.Id, "deleted_date": nil}).
		Update(&pl)
	if err != nil {
		return domain.Plant{}, err
	}

	result := r.mapModelToDomain(pl)
	return result, nil
}

func (r plantRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r plantRepository) mapDomainToModel(p domain.Plant) plant {
	return plant{
		Id:          p.Id,
		UserId:      p.UserId,
		Name:        p.Name,
		Address:     p.Address,
		Lat:         p.Lat,
		Lon:         p.Lon,
		Type:        p.Type,
		CreatedDate: p.CreatedDate,
		UpdatedDate: p.UpdatedDate,
		DeletedDate: p.DeletedDate,
	}
}

func (r plantRepository) mapModelToDomain(p plant) domain.Plant {
	return domain.Plant{
		Id:          p.Id,
		UserId:      p.UserId,
		Name:        p.Name,
		Address:     p.Address,
		Lat:         p.Lat,
		Lon:         p.Lon,
		Type:        p.Type,
		CreatedDate: p.CreatedDate,
		UpdatedDate: p.UpdatedDate,
		DeletedDate: p.DeletedDate,
	}
}

func (r plantRepository) mapModelToDomainCollection(plants []plant) []domain.Plant {
	var ps []domain.Plant
	for _, p := range plants {
		ps = append(ps, r.mapModelToDomain(p))
	}
	return ps
}
