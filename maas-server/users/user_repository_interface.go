package users

type UserRepositoryInterface interface {
	Get(id uint32) (User, error)
	GetByUsername(username string) (User, error)
	List() ([]User, error)
	Save(user User) (User, error)
	Delete(id uint32) (User, error)
}
