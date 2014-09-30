package v1

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"code.google.com/p/go.net/context"

	"github.com/dgrijalva/jwt-go"
	"github.com/hackedu/backend/v1/database"
	"github.com/hackedu/backend/v1/user"
	"github.com/hackedu/backend/httputil"
)

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rv := recover(); rv != nil {
			err := errors.New("handler panic")
			logError(r, err, rv)
			handleAPIError(w, r, http.StatusInternalServerError, err, false)
		}
	}()
	var (
		rb  httputil.ResponseBuffer
		err error
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if r.Header.Get("Authorization") != "" {
		u, err := getUserFromToken(r)
		if err != nil {
			if e, ok := err.(*httputil.HTTPError); ok {
				if e.Status >= 500 {
					logError(r, err, nil)
				}
				handleAPIError(w, r, e.Status, e.Err, true)
			} else {
				logError(r, err, nil)
				handleAPIError(w, r, http.StatusInternalServerError, err, false)
			}
		}

		ctx = user.NewContext(ctx, u)
	}

	err = h(ctx, &rb, r)
	if err == nil {
		rb.WriteTo(w)
	} else if e, ok := err.(*httputil.HTTPError); ok {
		if e.Status >= 500 {
			logError(r, err, nil)
		}
		handleAPIError(w, r, e.Status, e.Err, true)
	} else {
		logError(r, err, nil)
		handleAPIError(w, r, http.StatusInternalServerError, err, false)
	}
}

func getUserFromToken(r *http.Request) (*user.User, error) {
	token, err := jwt.ParseFromRequest(r,
		func(t *jwt.Token) (interface{}, error) {
			// TODO: Use real secret
			return []byte("secret"), nil
		})
	if err != nil {
		return nil, &httputil.HTTPError{http.StatusUnauthorized,
			errors.New("bad authorization token")}
	}
	userID := int64(token.Claims["id"].(float64))
	user, err := database.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &httputil.HTTPError{http.StatusNotFound,
				errors.New("user from token not found")}
		}
		return nil, err
	}
	return user, nil
}

func logError(req *http.Request, err error, rv interface{}) {
	if err != nil {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "Error serving %s: %v\n", req.URL, err)
		if rv != nil {
			fmt.Fprintln(&buf, rv)
			buf.Write(debug.Stack())
		}
		log.Println(buf.String())
	}
}

func handleAPIError(resp http.ResponseWriter, req *http.Request,
	status int, err error, showErrorMsg bool) {
	var data struct {
		Error struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		} `json:"error"`
	}
	data.Error.Status = status
	if showErrorMsg {
		data.Error.Message = err.Error()
	} else {
		data.Error.Message = http.StatusText(status)
	}
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(status)
	json.NewEncoder(resp).Encode(&data)
}
