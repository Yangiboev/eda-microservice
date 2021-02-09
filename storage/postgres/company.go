package postgres

import (
	"fmt"
	"gitlab.udevs.io/macbro/mb_corporate_service/mb_variables/corporate_service"

	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage/repo"
)

type companyRepo struct {
	db *sqlx.DB
}

func NewCompany(db *sqlx.DB) repo.CompanyI {
	return &companyRepo{
		db: db,
	}
}

func (comp *companyRepo) Create(company *corporate_service.Company) (string, error) {

	query := `insert into companies(
                        id,
                        name,
                        telephone,
                        account_number,
                        email,
                        inn,
                        mfo,
                        address,
                        description) 
                        values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	fmt.Println(company.ID)
	_, err := comp.db.Exec(
		query,
		company.ID,
		company.Name,
		company.Telephone,
		company.AccountNumber,
		company.Email,
		company.Inn,
		company.Mfo,
		company.Address,
		company.Description,
	)

	if err != nil {
		return "", err
	}

	return company.ID, nil
}

func (comp *companyRepo) Get(id string) (*corporate_service.Company, error) {
	var (
		company corporate_service.Company
	)
	query := `select 
       	id, 
       	name, 
		telephone,
		account_number,
		email,
		inn,
		mfo,
		address,
       	description
	from companies where deleted_at is null and id=$1`


	err := comp.db.Get(&company, query, id)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (comp *companyRepo) GetAll(page, limit int32, name string) ([]*corporate_service.Company, int32, error) {
	var (
		companies []*corporate_service.Company
		filter    string
		count     int32
	)

	if name != "" {
		filter += fmt.Sprintf(" name ilike '%s'", "%"+name+"%")
	}

	query := `select
		count(1) OVER(),
       	id, 
       	name, 
		telephone,
		account_number, 
		email,
		inn,
		mfo,
		address,
       	description
	from companies where deleted_at is null `

	//stmt, err := comp.db.Prepare(query + filter)

	//if err != nil {
	//	return nil, 0, err
	//}
	err := comp.db.Select(&companies, query)

	if err != nil {
		return nil, 0, err
	}

	//_, err = stmt.Query()

	if len(companies) > 0 {
		count = companies[0].Count
	}

	return companies, count, nil
}

func (comp *companyRepo) Update(company *corporate_service.Company) error {
	query := `update companies set 
                	name = $1,
					telephone = $2,
					account_number = $3,
					email = $4,
					inn = $5,
					mfo = $6,
					address = $7,
                    description = $8,
                    updated_at = CURRENT_TIMESTAMP
			where id = $9`

	_, err := comp.db.Exec(
		query,
		company.Name,
		company.Telephone,
		company.AccountNumber,
		company.Email,
		company.Inn,
		company.Mfo,
		company.Address,
		company.Description,
		company.ID,
	)

	return err
}

func (comp *companyRepo) Delete(id string) error {
	query := "update companies set deleted_at=CURRENT_TIMESTAMP where id=$1"
	stmt, err := comp.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}
