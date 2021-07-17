// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// kuiper - Go(lang) restful starter kit
//
// API Docs for kuiper v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <ribice@gmail.com> https://ribice.ba
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package api

import (
	"crypto/sha1"
	"os"

	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/zlog"

	"github.com/soldevx/kuiper/kuipersrv/pkg/api/auth"
	al "github.com/soldevx/kuiper/kuipersrv/pkg/api/auth/logging"
	at "github.com/soldevx/kuiper/kuipersrv/pkg/api/auth/transport"
	"github.com/soldevx/kuiper/kuipersrv/pkg/api/password"
	pl "github.com/soldevx/kuiper/kuipersrv/pkg/api/password/logging"
	pt "github.com/soldevx/kuiper/kuipersrv/pkg/api/password/transport"
	"github.com/soldevx/kuiper/kuipersrv/pkg/api/user"
	ul "github.com/soldevx/kuiper/kuipersrv/pkg/api/user/logging"
	ut "github.com/soldevx/kuiper/kuipersrv/pkg/api/user/transport"

	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/config"
	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/jwt"
	authMw "github.com/soldevx/kuiper/kuipersrv/pkg/utl/middleware/auth"
	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/postgres"
	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/rbac"
	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/secure"
	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New(os.Getenv("DATABASE_URL"), cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.Service{}
	jwt, err := jwt.New(cfg.JWT.SigningAlgorithm, os.Getenv("JWT_SECRET"), cfg.JWT.DurationMinutes, cfg.JWT.MinSecretLength)
	if err != nil {
		return err
	}

	log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	authMiddleware := authMw.Middleware(jwt)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, authMiddleware)

	v1 := e.Group("/v1")
	v1.Use(authMiddleware)

	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
