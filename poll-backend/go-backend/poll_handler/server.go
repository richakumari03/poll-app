package poll_handler

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("OPTIONS").PathPrefix("/").Handler(httptransport.NewServer(
		endpoints.Optionsendpoint,
		decodeOptionsCall,
		encodeOptionsCall,
	))

	r.Methods("POST").Path("/signup").Handler(httptransport.NewServer(
		endpoints.Signin,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		decodeUserReq,
		encodeLoginResponse,
	))

	r.Methods("GET").Path("/logout").Handler(httptransport.NewServer(
		endpoints.Logout,
		decodeLogoutReq,
		encodeLogoutResponse,
	))

	r.Methods("POST").Path("/createPoll").Handler(httptransport.NewServer(
		endpoints.CreatePoll,
		decodeCreatePollReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/polls").Handler(httptransport.NewServer(
		endpoints.GetPolls,
		decodePollReq,
		encodePollsResponse,
	))

	r.Methods("POST").Path("/updateVote").Handler(httptransport.NewServer(
		endpoints.UpdateVote,
		decodeUpdateVoteReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(w, r)
	})
}
