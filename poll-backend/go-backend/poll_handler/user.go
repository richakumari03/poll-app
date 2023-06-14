package poll_handler

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	username string
	email    string
	password string
}

var sessions = map[string]Session{}

type Session struct {
	Username string
	Expiry   time.Time
}

type Cookie struct {
	Value  string
	Expiry time.Time
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type PollDetails struct {
	QuestionId int
	QDesc      string
	OptionId   int
	ODesc      string
	TotalCount int
}

type PollByUserDetails struct {
	QuestionId int
	OptionId   int
}

func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}

type Repository interface {
	CreateUser(context.Context, User) error
	UserExists(username string) (bool, error)
	GetPassword(username string) (string, error)
	AddPollQuestion(context.Context, string) (int, error)
	AddPollOptions(context.Context, int, []string) error
	GetPolls(context.Context) ([]PollDetails, error)
	GetPollsByUser(context.Context, string) ([]PollByUserDetails, error)
	UpdateVote(context.Context, int, int, string) error
}
