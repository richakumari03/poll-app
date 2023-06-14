package poll_handler

import (
	"context"
)

type Service interface {
	Signin(ctx context.Context, user User) (int, error)
	Login(ctx context.Context, user User) (Cookie, int, error)
	Logout(ctx context.Context, token string) int
	CreatePoll(ctx context.Context, question string, option []string, token string) (int, error)
	GetPolls(ctx context.Context, username string, token string) ([]PollDetails, []PollByUserDetails, int, error)
	UpdateVote(ctx context.Context, questionId int, optionId int, username string, token string) (int, error)
}
