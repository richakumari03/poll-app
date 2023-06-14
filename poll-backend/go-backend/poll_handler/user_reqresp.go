package poll_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type (
	SignInRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		user User
	}

	LogoutRequest struct {
		Token string
	}

	GenericResponse struct {
		Status int
	}

	LoginResponse struct {
		Cookie Cookie
		Status int
	}

	CreatePollRequest struct {
		Question string   `json:"question"`
		Option   []string `json:"option"`
		Token    string
	}

	GetPollsRequest struct {
		Username string `json:"username"`
		Token    string
	}

	GetPollsResponse struct {
		Result     []PollDetails
		UserResult []PollByUserDetails
		Status     int
	}

	UpdateVoteRequest struct {
		QuestionId int
		OptionId   int
		Username   string
		Token      string
	}
)

func getAuthToken(r *http.Request) string {
	var token string = ""

	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	}

	return token
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, _ := response.(GenericResponse)
	w.WriteHeader(res.Status)

	return json.NewEncoder(w).Encode(response)
}

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, _ := response.(LoginResponse)
	w.WriteHeader(res.Status)
	return json.NewEncoder(w).Encode(res.Cookie)
}

func encodePollsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, _ := response.(GetPollsResponse)
	w.WriteHeader(res.Status)

	return json.NewEncoder(w).Encode(response)
}

func encodeLogoutResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, _ := response.(GenericResponse)
	w.WriteHeader(res.Status)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func decodeLogoutReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req LogoutRequest

	token := getAuthToken(r)
	req.Token = token

	return req, nil
}

func decodeCreatePollReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreatePollRequest

	token := getAuthToken(r)
	req.Token = token

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodePollReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetPollsRequest

	token := getAuthToken(r)
	req.Token = token

	if token != "" {
		req.Username = sessions[token].Username
	}
	return req, nil
}

func decodeUpdateVoteReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateVoteRequest

	token := getAuthToken(r)
	req.Token = token
	if token != "" {
		req.Username = sessions[token].Username
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func decodeOptionsCall(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeOptionsCall(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return nil
}
