package repositories

import (
	"categories-api/models"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository{
	return &ProductRepository{db : db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error){
	query :=`
		SELECT 
			products.id,
			products.name,
			products.price,
			products.stock,
			categories.id,
			categories.title
		FROM products 
		INNER JOIN categories 
			ON products.category_id = categories.id
	`
				
	rows, err := repo.db.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product,0)
	for rows.Next(){
		var p models.Product
		err := rows.Scan(&p.ID,&p.Name,&p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}


func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}
