package repo

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Get(email, pass string) (*User, error)
	// List() []*User
	// Delete(UserId int) error
	// Update(p User) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	repo := &userRepo{
		dbCon: dbCon,
	}

	return repo
}

func (r userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner)
		VALUES (
			:first_name, 
			:last_name, 
			:email, 
			:password, 
			:is_shop_owner)
		RETURNING id;
	`

	var id int
	rows, err := r.dbCon.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	// defer rows.Close()

	if rows.Next() {
		rows.Scan(&id)
	}

	user.ID = id

	return &user, nil

}

func (r *userRepo) Get(email, pass string) (*User, error) {
	var user User

	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`

	err := r.dbCon.Get(&user, query, email, pass)
	if err != nil {
		return nil, err // returns sql.ErrNoRows if not found
	}

	return &user, nil
}
