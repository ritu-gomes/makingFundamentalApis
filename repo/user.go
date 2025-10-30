package repo

import (
	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
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

func (r userRepo) Create(user domain.User) (*domain.User, error) {
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

func (r *userRepo) Get(email, pass string) (*domain.User, error) {
	var user domain.User
 
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
