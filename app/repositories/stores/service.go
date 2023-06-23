package stores

import (
	"Technopark_DB_Project/app/models"
	"Technopark_DB_Project/app/repositories"
	dbRequests "Technopark_DB_Project/app/requests"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type ServiceStore struct {
	db *pgx.ConnPool
}

func CreateServiceRepository(db *pgx.ConnPool) repositories.ServiceRepository {
	return &ServiceStore{db: db}
}

func (serviceStore *ServiceStore) Clear() (err error) {
	_, err = serviceStore.db.Exec(dbRequests.CLEAR_ALL)
	return
}

func (serviceStore *ServiceStore) GetStatus() (status *models.Status, err error) {
	status = &models.Status{}
	err = serviceStore.db.QueryRow(dbRequests.GET_STATUS).
		Scan(
			&status.User,
			&status.Forum,
			&status.Thread,
			&status.Post)

	return
}
