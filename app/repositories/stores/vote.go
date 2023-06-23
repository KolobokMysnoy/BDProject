package stores

import (
	"Technopark_DB_Project/app/models"
	"Technopark_DB_Project/app/repositories"
	dbRequests "Technopark_DB_Project/app/requests"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type VoteStore struct {
	db *pgx.ConnPool
}

func CreateVoteRepository(db *pgx.ConnPool) repositories.VoteRepository {
	return &VoteStore{db: db}
}

func (voteStore *VoteStore) Vote(threadID int64, vote *models.Vote) (err error) {
	_, err = voteStore.db.Exec(dbRequests.CREATE_VOTE,
		vote.Nickname, threadID, vote.Voice)
	return
}
