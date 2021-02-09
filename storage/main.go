package storage

import (
	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage/postgres"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage/repo"
)

type StorageI interface {
	Company() repo.CompanyI
}

type storagePg struct {
	companyRepo repo.CompanyI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		companyRepo: postgres.NewCompany(db),
	}
}

func (s *storagePg) Company() repo.CompanyI {
	return s.companyRepo
}
