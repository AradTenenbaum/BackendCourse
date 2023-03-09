package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/AradTenenbaum/BackendCourse/db/sqlc"
	"github.com/AradTenenbaum/BackendCourse/val"
	"github.com/rs/zerolog/log"
)

const (
	ID          = "id"
	SECRET_CODE = "secret_code"
	MESSAGE     = "message"
)

func (httpServer *HttpServer) VerifyEmailHandlerFunc(res http.ResponseWriter, req *http.Request) {

	resp := make(map[string]string)
	secretCode := req.URL.Query().Get(SECRET_CODE)
	id, err := strconv.Atoi(req.URL.Query().Get(ID))

	if err != nil {
		resp[MESSAGE] = fmt.Sprintf("Error in convertion: %s", err.Error())
		jsonData, err := json.Marshal(resp)
		if err != nil {
			log.Fatal().Err(err).Msg("Cannot marshal")
		}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonData)
		return
	}

	// validation for secret code to avoid sql injection
	err = val.ValidateSecretCode(secretCode)
	if err != nil {
		resp[MESSAGE] = fmt.Sprintf("Error in code: %s", err.Error())
		jsonData, err := json.Marshal(resp)
		if err != nil {
			log.Fatal().Err(err).Msg("Cannot marshal")
		}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonData)
		return
	}

	// Verify the email
	_, err = httpServer.store.VerifyEmailTx(context.Background(), db.SetUsedVerifyEmailParams{
		SecretCode: secretCode,
		ID:         int64(id),
	})
	if err != nil {
		resp[MESSAGE] = err.Error()
		jsonData, err := json.Marshal(resp)
		if err != nil {
			log.Fatal().Err(err).Msg("Cannot marshal")
		}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonData)
		return
	}
	res.Write([]byte(`<h1 style="text-align:center;padding:50px;font-family:verdana;">Your Email is Verified✅</h1>`))
}
