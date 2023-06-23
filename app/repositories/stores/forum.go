package stores

import (
	"Technopark_DB_Project/app/models"
	"Technopark_DB_Project/app/repositories"
	dbRequests "Technopark_DB_Project/app/requests"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type ForumStore struct {
	db *pgx.ConnPool
}

func CreateForumRepository(db *pgx.ConnPool) repositories.ForumRepository {
	return &ForumStore{db: db}
}

func (forumStore *ForumStore) Create(forum *models.Forum) (err error) {
	_, err = forumStore.db.Exec(dbRequests.CREATE_FORUM,
		forum.Title,
		forum.User,
		forum.Slug)
	return
}

func (forumStore *ForumStore) GetBySlug(slug string) (forum *models.Forum, err error) {
	forum = new(models.Forum)
	err = forumStore.db.QueryRow(dbRequests.GET_FORUM_BY_SLUG, slug).
		Scan(
			&forum.Title,
			&forum.User,
			&forum.Slug,
			&forum.Posts,
			&forum.Threads)
	return
}

func (forumStore *ForumStore) GetUsers(slug string, limit int, since string, desc bool) (users *[]models.User, err error) {
	var usersSlice []models.User

	var resultRows *pgx.Rows

	query := dbRequests.DEFAULT_GET_USERS +
		"LEFT JOIN user_forum as UF ON U.nickname = UF.nickname WHERE UF.forum = $1"

	if since != "" {
		if desc {
			query += " AND U.nickname < $2 ORDER BY U.nickname DESC"
		} else {
			query += " AND U.nickname > $2 ORDER BY U.nickname"
		}
		query += " LIMIT $3;"
		resultRows, err = forumStore.db.Query(query,
			slug,
			since,
			limit)
	} else {
		if desc {
			query += " ORDER BY U.nickname DESC"
		} else {
			query += " ORDER BY U.nickname"
		}
		query += " LIMIT $2;"
		resultRows, err = forumStore.db.Query(query,
			slug,
			limit)
	}

	if err != nil {
		return
	}
	defer resultRows.Close()

	for resultRows.Next() {
		user := models.User{}
		err = resultRows.Scan(&user.Nickname,
			&user.Fullname,
			&user.About,
			&user.Email)
		if err != nil {
			return
		}
		usersSlice = append(usersSlice, user)
	}
	return &usersSlice, nil
}

func (forumStore *ForumStore) GetThreads(slug string, limit int, since string, desc bool) (threads *[]models.Thread, err error) {
	var threadsSlice []models.Thread

	var resultRows *pgx.Rows

	query := dbRequests.DEFAULT_GET_THREADS

	if since != "" {
		if desc {
			query += " AND created <= $2 ORDER BY created DESC"
		} else {
			query += " AND created >= $2 ORDER BY created ASC"
		}
		query += " LIMIT $3;"
		resultRows, err = forumStore.db.Query(query, slug, since, limit)
	} else {
		if desc {
			query += " ORDER BY created DESC"
		} else {
			query += " ORDER BY created ASC"
		}
		query += " LIMIT $2;"
		resultRows, err = forumStore.db.Query(query, slug, limit)
	}

	if err != nil {
		return
	}
	defer resultRows.Close()

	for resultRows.Next() {
		thread := models.Thread{}
		err = resultRows.Scan(
			&thread.ID,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&thread.Slug,
			&thread.Created)

		if err != nil {
			return
		}
		threadsSlice = append(threadsSlice, thread)
	}
	return &threadsSlice, nil
}
