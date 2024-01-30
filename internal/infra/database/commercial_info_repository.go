package database

import (
	"database/sql"

	"github.com/k1nha/travelreviews/internal/domain/entity"
)

type CommercialInfoRepository struct {
	Db *sql.DB
}

func (r *CommercialInfoRepository) Save(address *entity.CommercialInfo) error {
	return nil
}
