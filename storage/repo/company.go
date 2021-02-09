package repo

import (
	"gitlab.udevs.io/macbro/mb_corporate_service/mb_variables/corporate_service"
)

type CompanyI interface {
	Create(company *corporate_service.Company) (string, error)
	Update(company *corporate_service.Company)  error
	Get(id string) (*corporate_service.Company, error)
	GetAll(page, limit int32, name string) ([]*corporate_service.Company, int32, error)
	Delete(id string) error
}
