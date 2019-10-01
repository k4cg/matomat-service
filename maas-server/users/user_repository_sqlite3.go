package users

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepoSqlite3 struct {
	db *sql.DB
}

func NewUserRepoSqlite3(sqlite3DbFilePath string) *UserRepoSqlite3 {
	db, err := sql.Open("sqlite3", sqlite3DbFilePath)
	if err == nil {
		return &UserRepoSqlite3{db: db}
	} else {
		panic(err)
	}
}

func (r *UserRepoSqlite3) Get(userID uint32) (User, error) {
	var user User
	var err error

	rows, err := r.db.Query("SELECT ID, username, password, credits, admin FROM users WHERE id=?", userID)
	if err == nil {
		for rows.Next() {
			rows.Scan(&user.ID, &user.Username, &user.Password, &user.Credits, &user.admin)
			break
		}
		err = rows.Close()
	}

	return user, err
}

func (r *UserRepoSqlite3) GetByUsername(username string) (User, error) {
	var user User
	var err error

	rows, err := r.db.Query("SELECT ID, username, password, credits, admin FROM users WHERE username=?", username)
	if err == nil {
		for rows.Next() {
			rows.Scan(&user.ID, &user.Username, &user.Password, &user.Credits, &user.admin)
			break
		}
		err = rows.Close()
	}

	return user, err
}

func (r *UserRepoSqlite3) List() ([]User, error) {
	users := make([]User, 0)
	var err error

	rows, err := r.db.Query("SELECT ID, username, password, credits, admin FROM users")
	if err == nil {
		for rows.Next() {
			var id uint32
			var username string
			var password string
			var credits int32
			var adminInt uint32

			err = rows.Scan(&id, &username, &password, &credits, &adminInt)
			if err == nil {
				var adminBool bool
				if adminInt == 1 {
					adminBool = true
				}

				users = append(users, User{ID: id, Username: username, Password: password, Credits: credits, admin: adminBool})
			}
		}
		err = rows.Close()
	}

	return users, err
}

func (r *UserRepoSqlite3) Save(user User) (User, error) {
	var returnedUser User
	var err error
	oldUser, err := r.Get(user.ID)
	if err == nil {
		if oldUser == (User{}) {
			//create new item
			stmt, err := r.db.Prepare("INSERT INTO users (username, password, credits, admin) VALUES (?, ?, ?, ?)")
			if err == nil {
				adminInt := 0
				if user.IsAdmin() {
					adminInt = 1
				}
				res, err := stmt.Exec(user.Username, user.Password, user.Credits, adminInt)
				if err == nil {
					id, err := res.LastInsertId()
					if err == nil {
						//evil cast of int64 => uint32 .... TODO solve this better
						returnedUser = User{ID: uint32(id), Username: user.Username, Password: user.Password, Credits: user.Credits, admin: user.IsAdmin()}
					}
					err = stmt.Close()
				}
			}
		} else {
			stmt, err := r.db.Prepare("UPDATE users SET username=?, password=?, credits=?, admin=? WHERE ID=?")
			if err == nil {
				adminInt := 0
				if user.IsAdmin() {
					adminInt = 1
				}
				_, err = stmt.Exec(user.Username, user.Password, user.Credits, adminInt, user.ID)
				if err == nil {
					returnedUser = user
				}
				err = stmt.Close()
			}
		}
	}
	return returnedUser, err
}

func (r *UserRepoSqlite3) Delete(userID uint32) (User, error) {
	var user User
	var err error
	user, err = r.Get(userID)
	if err == nil {
		stmt, err := r.db.Prepare("DELETE FROM users WHERE ID=?")
		if err == nil {
			_, err = stmt.Exec(user.ID)
			err = stmt.Close()
		}
	}
	return user, err
}
