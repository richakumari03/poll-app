package poll_handler

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Signin          endpoint.Endpoint
	Login           endpoint.Endpoint
	Logout          endpoint.Endpoint
	CreatePoll      endpoint.Endpoint
	GetPolls        endpoint.Endpoint
	UpdateVote      endpoint.Endpoint
	Optionsendpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Signin:          makeSignInEndpoint(s),
		Login:           makeLoginEndpoint(s),
		Logout:          makeLogoutEndpoint(s),
		CreatePoll:      makeCreatePollEndpoint(s),
		GetPolls:        makeGetPollsEndpoint(s),
		UpdateVote:      makeUpdateVoteEndpoint(s),
		Optionsendpoint: makeOptionsEndpoint(s),
	}
}

func makeSignInEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignInRequest)
		user := User{
			username: req.Username,
			email:    req.Email,
			password: req.Password,
		}
		status, err := s.Signin(ctx, user)

		return GenericResponse{Status: status}, err
	}
}

func makeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignInRequest)
		user := User{
			username: req.Username,
			email:    req.Email,
			password: req.Password,
		}
		Cookie, status, err := s.Login(ctx, user)

		return LoginResponse{Cookie: Cookie, Status: status}, err
	}
}

func makeLogoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogoutRequest)

		status := s.Logout(ctx, req.Token)

		return GenericResponse{Status: status}, nil
	}
}

func makeCreatePollEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePollRequest)

		status, err := s.CreatePoll(ctx, req.Question, req.Option, req.Token)

		return GenericResponse{Status: status}, err
	}
}

func makeGetPollsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPollsRequest)

		result, res, status, err := s.GetPolls(ctx, req.Username, req.Token)

		return GetPollsResponse{Result: result, UserResult: res, Status: status}, err
	}
}

func makeUpdateVoteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateVoteRequest)

		status, err := s.UpdateVote(ctx, req.QuestionId, req.OptionId, req.Username, req.Token)

		return GenericResponse{Status: status}, err
	}
}

func makeOptionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}
