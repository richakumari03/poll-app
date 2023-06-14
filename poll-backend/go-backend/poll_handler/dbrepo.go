package poll_handler

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle db Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)`

	if user.email == "" || user.password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.username, user.email, user.password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) UserExists(username string) (bool, error) {
	var exists int
	err := repo.db.QueryRow("SELECT count(*) FROM users WHERE username=$1", username).Scan(&exists)

	if err != nil {
		return false, RepoErr
	}
	var customErr error = nil
	if exists > 0 {
		customErr = errors.New("User already exists")
	}
	return exists > 0, customErr
}

func (repo *repo) GetPassword(username string) (string, error) {
	var password string
	err := repo.db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&password)
	if err != nil {
		return "", RepoErr
	}

	return password, nil
}

func (repo *repo) AddPollQuestion(ctx context.Context, question string) (int, error) {

	sql := `
	INSERT INTO question (id, description)
	VALUES ((SELECT COALESCE(MAX(id) + 1, 1) FROM question), $1)`

	_, err := repo.db.ExecContext(ctx, sql, question)
	if err != nil {
		return 0, err
	}

	var id int
	err = repo.db.QueryRow("SELECT MAX(id) FROM question").Scan(&id)

	return id, nil
}

func (repo *repo) AddPollOptions(ctx context.Context, id int, option []string) error {

	sql := `
	INSERT INTO option (optionId, questionId, description)
	VALUES ($1, $2, $3)`

	for i := 0; i < len(option); i++ {
		_, err := repo.db.ExecContext(ctx, sql, i+1, id, option[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *repo) GetPolls(ctx context.Context) ([]PollDetails, error) {

	sql := `SELECT q.id, q.description, o.optionId, o.description, COUNT(DISTINCT p.username) AS totalUserCount
	FROM question q
	INNER JOIN option o ON q.id = o.questionId
	LEFT JOIN poll p ON q.id = p.questionId AND o.optionId = p.optionId
	GROUP BY q.id, o.optionId, q.description, o.description;`

	rows, err := repo.db.Query(sql)

	var result []PollDetails

	for rows.Next() {
		var row PollDetails
		err := rows.Scan(&row.QuestionId, &row.QDesc, &row.OptionId, &row.ODesc, &row.TotalCount)
		if err != nil {
			return result, err
		}
		result = append(result, row)
	}

	return result, err
}

func (repo *repo) GetPollsByUser(ctx context.Context, username string) ([]PollByUserDetails, error) {

	sql := `SELECT questionId, optionId
	FROM poll where username = ($1)`

	rows, err := repo.db.Query(sql, username)

	var result []PollByUserDetails

	for rows.Next() {
		var row PollByUserDetails
		err := rows.Scan(&row.QuestionId, &row.OptionId)
		if err != nil {
			return result, err
		}
		result = append(result, row)
	}

	return result, err
}

func (repo *repo) UpdateVote(ctx context.Context, questionId int, optionId int, username string) error {

	sql := `INSERT INTO poll (optionid, questionid, username)
	VALUES ($1, $2, $3)`

	_, err := repo.db.ExecContext(ctx, sql, optionId, questionId, username)
	if err != nil {
		return err
	}

	return nil
}
