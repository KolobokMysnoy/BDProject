package stores

import (
	"Technopark_DB_Project/app/models"
	"Technopark_DB_Project/app/repositories"
	dbRequests "Technopark_DB_Project/app/requests"
	"time"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type PostStore struct {
	db *pgx.ConnPool
}

func CreatePostRepository(db *pgx.ConnPool) repositories.PostRepository {
	return &PostStore{db: db}
}

func (postStore *PostStore) GetByID(id int64) (post *models.Post, err error) {
	post = &models.Post{}
	postTime := time.Time{}
	err = postStore.db.QueryRow(dbRequests.GET_POST_BY_ID, id).
		Scan(
			&post.ID,
			&post.Parent,
			&post.Author,
			&post.Message,
			&post.IsEdited,
			&post.Forum,
			&post.Thread,
			&postTime)

	post.Created = postTime.Format(time.RFC3339)
	return
}

func (postStore *PostStore) Update(post *models.Post) (err error) {
	_, err = postStore.db.Exec(dbRequests.UPDATE_POST,
		post.Message,
		post.IsEdited,
		post.ID)
	return
}
