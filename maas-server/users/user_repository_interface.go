package users

type UserRepositoryInterface interface {
	Get(id uint32) (User, error)
	GetByUsername(username string) (User, error)
	List() (map[uint32]User, error)
	Save(user User) (User, error)
	Delete(id uint32) (User, error)
}

const ERROR_UNKOWN_USER string = "Unkown user"
