package products

import "database/sql"

type MySqlRepository struct {
	db *sql.DB
}

func (m *MySqlRepository) GetAll() ([]entities.Product, error) {
	return nil, nil
}

func (m *MySqlRepository) Store(name, category string, count int, price float64) (entities.Product, error) {
	return entities.Product{}, nil
}

func (m *MySqlRepository) LastID() (uint64, error) {
	return 0, nil
}
