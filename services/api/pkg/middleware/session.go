package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"github.com/pkg/errors"
	"io"
	"k8s.io/klog/v2"
	"os"
)

func Session() gin.HandlerFunc {
	key := make([]byte, 64)
	err := os.MkdirAll("/secret", 0755)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("/secret/session.key", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	n, err := hex.NewDecoder(f).Read(key)
	if err != nil && errors.Cause(err) != io.EOF {
		panic(errors.Wrap(err, "invalid contents at /secret/session.key"))
	}
	if n == 0 {
		klog.Info("Session key doesn't exist, generating it...")
		f.Seek(0, io.SeekStart)
		n, err = rand.Read(key)
		if err != nil {
			panic(err)
		}
		if n != 64 {
			panic(fmt.Errorf("could not generate 64 bytes, only got %v", n))
		}
		n, err = hex.NewEncoder(f).Write(key)
		if err != nil {
			panic(err)
		}
		if n != 64 {
			panic(fmt.Errorf("could not write 64 bytes, only wrote %v", n))
		}
	} else if n != 64 {
		panic(errors.Wrap(err, "not 64 bytes of content at /secret/session.key"))
	}

	store := sessions.NewCookieStore(key)

	return func(ginCtx *gin.Context) {
		ginCtx.Set("default_cookie_store", store)
	}
}

func GetCookieStore(ginCtx *gin.Context) *sessions.CookieStore {
	x, ok := ginCtx.Get("default_cookie_store")
	if !ok {
		panic(fmt.Errorf("tried to get session when session middleware was not added"))
	}
	return x.(*sessions.CookieStore)
}

func SaveSession(ginCtx *gin.Context, session *sessions.Session) error {
	return GetCookieStore(ginCtx).Save(ginCtx.Request, ginCtx.Writer, session)
}

func GetSession(ginCtx *gin.Context) (*sessions.Session, error) {
	return GetCookieStore(ginCtx).Get(ginCtx.Request, "default")
}

func GetUser(ginCtx *gin.Context) (*model.User, error) {
	session, err := GetSession(ginCtx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get session")
	}

	userUuid, ok := session.Values["uuid"]
	if !ok || userUuid == "" {
		return nil, fmt.Errorf( "user has not logged in yet")
	}

	user := &model.User{}
	err = model.Get(user, "SELECT * FROM users WHERE uuid=$1", userUuid)
	if err != nil {
		return nil, errors.Wrap(err, "could not get user by that uuid")
	}
	return user, nil
}