package poll_handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key_123")

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func validateSession(sessionToken string) int {

	if sessionToken == "" {
		return http.StatusUnauthorized
	}

	userSession, exists := sessions[sessionToken]
	if !exists {
		return http.StatusUnauthorized
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		return http.StatusUnauthorized
	}

	return http.StatusOK
}

// Sign up a new user
func (s service) Signin(ctx context.Context, user User) (int, error) {
	logger := log.With(s.logger, "method", "SignIn")

	if exists, err := s.repostory.UserExists(user.username); err != nil || exists {
		level.Error(logger).Log("err", err)
		return http.StatusInternalServerError, err
	}

	if err := s.repostory.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// Login an existing user
func (s service) Login(ctx context.Context, user User) (Cookie, int, error) {

	expectedPassword, err := s.repostory.GetPassword(user.username)
	var userCookie Cookie

	if err != nil || expectedPassword != user.password {
		return userCookie, http.StatusUnauthorized, err
	}

	expiresAt := time.Now().Add(120 * time.Second)

	claims := &Claims{
		Username: user.username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return userCookie, http.StatusInternalServerError, err
	}

	sessions[tokenString] = Session{
		Username: user.username,
		Expiry:   expiresAt,
	}

	userCookie = Cookie{
		Value:  tokenString,
		Expiry: expiresAt,
	}
	return userCookie, http.StatusOK, err
}

// Logout a loggedin user
func (s service) Logout(ctx context.Context, token string) int {

	status := validateSession(token)
	if status != http.StatusOK {
		return status
	}

	sessionToken := token

	delete(sessions, sessionToken)

	return http.StatusOK
}

// Create a poll
func (s service) CreatePoll(ctx context.Context, question string, option []string, token string) (int, error) {
	logger := log.With(s.logger, "method", "CreatePoll")

	status := validateSession(token)
	if status != http.StatusOK {
		return status, nil
	}

	var id int
	var err error
	if id, err = s.repostory.AddPollQuestion(ctx, question); err != nil {
		level.Error(logger).Log("err", err)
		return http.StatusInternalServerError, err
	}

	if err := s.repostory.AddPollOptions(ctx, id, option); err != nil {
		level.Error(logger).Log("err", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// Fetch all the polls
func (s service) GetPolls(ctx context.Context, username string, token string) ([]PollDetails, []PollByUserDetails, int, error) {
	logger := log.With(s.logger, "method", "GetPolls")

	status := validateSession(token)
	if status != http.StatusOK {
		return nil, nil, status, nil
	}

	var result []PollDetails
	var res []PollByUserDetails
	var err error
	if username != "" {
		if result, err = s.repostory.GetPolls(ctx); err != nil {
			level.Error(logger).Log("err", err)
			return result, res, http.StatusInternalServerError, err
		}

		if res, err = s.repostory.GetPollsByUser(ctx, username); err != nil {
			level.Error(logger).Log("err", err)
			return result, res, http.StatusInternalServerError, err
		}
	}

	return result, res, http.StatusOK, nil
}

func (s service) UpdateVote(ctx context.Context, questionId int, optionId int, username string, token string) (int, error) {
	logger := log.With(s.logger, "method", "UpdateVote")

	status := validateSession(token)
	if status != http.StatusOK {
		return status, nil
	}

	if username != "" {
		if err := s.repostory.UpdateVote(ctx, questionId, optionId, username); err != nil {
			level.Error(logger).Log("err", err)
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
