package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface{
	product.ProductRepo
}
type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (
			title,
			description, 
			price, 
			image_url)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4)
		RETURNING id `

		row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl)
		err := row.Scan(&p.Id)
		if err != nil {
			return nil, err
		}

		return &p, nil
}

func (r *productRepo) Get(productId int) (*domain.Product, error) {
	var prod domain.Product

	query := `
		SELECT id, title, description, price, img_url
		FROM products
		WHERE id = $1
		LIMIT 1
	`

	err := r.db.Get(&prod, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prod, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	var prodList []*domain.Product

	query := `SELECT * FROM products;`

	err := r.db.Select(&prodList, query)
	if err != nil {
		return nil, err
	}

	return prodList, nil
}

func (r *productRepo) Delete(productId int) error {
	query := `
		DELETE FROM products WHERE id = $1
	`
	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products 
		SET title=$1, description=$2, price=$3, img_url=$4
		WHERE id = $5
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &p, nil
}