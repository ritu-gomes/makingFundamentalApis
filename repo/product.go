package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageUrl    string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(productId int) (*Product, error) {
	var prod Product

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

func (r *productRepo) List() ([]*Product, error) {
	var prodList []*Product

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

func (r *productRepo) Update(p Product) (*Product, error) {
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