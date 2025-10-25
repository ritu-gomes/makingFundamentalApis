package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"Email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Get(email, pass string) (*User, error)
	// List() []*User
	// Delete(UserId int) error
	// Update(p User) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}

	return repo
}

func (r userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1

	r.users = append(r.users, user)
	return &user, nil
}

func (r *userRepo) Get(email, pass string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email && u.Password == pass {
			return &u, nil
		}
	}

	return nil, nil
}
