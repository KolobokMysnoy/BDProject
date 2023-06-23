package stores

import (
	"Technopark_DB_Project/app/models"
	"Technopark_DB_Project/app/repositories"
	dbRequests "Technopark_DB_Project/app/requests"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type UserStore struct {
	db *pgx.ConnPool
}

func CreateUserRepository(db *pgx.ConnPool) repositories.UserRepository {
	return &UserStore{db: db}
}

func (userStore *UserStore) Create(user *models.User) (err error) {
	_, err = userStore.db.Exec(dbRequests.CREATE_USER,
		user.Nickname, user.Fullname, user.About, user.Email)
	return
}

func (userStore *UserStore) Update(user *models.User) (err error) {
	return userStore.db.QueryRow(dbRequests.UPDATE_USER,
		user.Fullname,
		user.About,
		user.Email,
		user.Nickname).
		Scan(&user.Fullname, &user.About, &user.Email)
}

func (userStore *UserStore) GetByNickname(nickname string) (user *models.User, err error) {
	user = new(models.User)
	err = userStore.db.QueryRow(dbRequests.GET_USER_BY_NICKNAME, nickname).Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
	return
}

func (userStore *UserStore) GetAllMatchedUsers(user *models.User) (users *[]models.User, err error) {
	var usersSlice []models.User

	resultRows, err := userStore.db.Query(dbRequests.GET_MATCHED_USERS, user.Nickname, user.Email)
	if err != nil {
		return
	}
	defer resultRows.Close()

	for resultRows.Next() {
		user := models.User{}
		err = resultRows.Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
		if err != nil {
			return
		}
		usersSlice = append(usersSlice, user)
	}
	return &usersSlice, nil
}
