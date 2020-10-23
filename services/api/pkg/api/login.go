package api

import (
	"database/sql"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/layer-devops/livechat-example/services/api/pkg/middleware"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
	"strings"
)

func AddLoginRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", func(ginCtx *gin.Context) {
		name := strings.TrimSpace(ginCtx.PostForm("name"))
		password := strings.TrimSpace(ginCtx.PostForm("password"))

		redirectDest := "/"
		referer, err := url.Parse(ginCtx.Request.Header.Get("Referer"))
		if err == nil && referer.Scheme != "" {
			redirectDest = referer.Scheme+"://"+referer.Host
		} else {
			redirectDest = "//"
		}

		if name == "" || password == "" {
			ginCtx.Redirect(http.StatusSeeOther, redirectDest+"/login")
			return
		}

		user := &model.User{}
		err = model.Get(user, "SELECT * FROM users WHERE LOWER(name)=LOWER($1)", name)
		if err == sql.ErrNoRows {
			user.Name = name
			passBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				klog.Error(err)
				ginCtx.Redirect(http.StatusSeeOther, redirectDest+"/login?error=Internal+error+while+saving+password")
				return
			}
			user.PasswordHash = hex.EncodeToString(passBytes)
			err = model.Get(user, "INSERT INTO users(name, password_hash) VALUES ($1, $2) RETURNING *", user.Name, user.PasswordHash)
			if err != nil {
				klog.Error(err)
				ginCtx.Redirect(http.StatusSeeOther, redirectDest+"/login?error=Internal+error+while+registering")
				return
			}
		} else if err != nil {
			klog.Error(err)
			ginCtx.Redirect(http.StatusSeeOther, redirectDest+"/login?error=Internal+error+while+logging+in")
			return
		} else {
			decoded, _ := hex.DecodeString(user.PasswordHash)
			err = bcrypt.CompareHashAndPassword(decoded, []byte(password))
			if err != nil {
				klog.Warning("could not check password: ", err)
				ginCtx.Redirect(http.StatusSeeOther, redirectDest+"/login?error=Invalid+password")
				return
			}
		}

		session, err := middleware.GetSession(ginCtx)
		if err != nil {
			panic(err)
		}
		session.Values["uuid"] = user.UUID
		err = middleware.SaveSession(ginCtx, session)
		if err != nil {
			panic(err)
		}

		ginCtx.Redirect(http.StatusSeeOther, redirectDest)
	})
}
