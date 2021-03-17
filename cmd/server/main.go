package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/handlers"
	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
)

func main() {
	appSettings, err := settings.NewSettings()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := newLogger(appSettings.Debug, appSettings.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	daos, err := dao.New(appSettings.DB)
	if err != nil {
		logger.Fatal("error creating daos", zap.Error(err))
	}
	var (
		mw                 = handlers.NewMiddleware(logger, daos.User)
		serviceHandler     = handlers.NewService(logger, daos.Service)
		serviceTypeHandler = handlers.NewServiceType(logger, daos.ServiceType)
		vehicleHandler     = handlers.NewVehicle(logger, daos.Vehicle)
		userHandler        = handlers.NewUser(logger, daos.User)
		authHandler        = handlers.NewAuth(logger, daos.User, appSettings.GitHubAuth)
	)

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	vehicles := api.PathPrefix("/vehicles").Subrouter()
	services := api.PathPrefix("/services").Subrouter()
	serviceTypes := api.PathPrefix("/service_types").Subrouter()
	users := api.PathPrefix("/users").Subrouter()
	oauth := api.PathPrefix("/oauth").Subrouter()

	vehicles.Handle("", mw.Reduce(vehicleHandler.List, mw.Standard...))
	vehicles.Handle("/{id:[0-9]+}", mw.Reduce(vehicleHandler.Detail, mw.Standard...))

	services.Handle("", mw.Reduce(serviceHandler.List, mw.Standard...))
	services.Handle("/{id:[0-9]+}", mw.Reduce(serviceHandler.Detail, mw.Standard...))

	serviceTypes.Handle("", mw.Reduce(serviceTypeHandler.List, mw.Standard...))
	serviceTypes.Handle("/{id:[0-9]+}", mw.Reduce(serviceTypeHandler.Detail, mw.Standard...))

	users.Handle("/me", mw.Reduce(userHandler.Me, mw.Standard...))

	oauth.Handle("/login", mw.Reduce(authHandler.Login, mw.Log))
	oauth.Handle("/callback", mw.Reduce(authHandler.Callback, mw.Log))

	fs := http.FileServer(http.Dir(appSettings.StaticDir))
	r.PathPrefix("/").Handler(fs)

	logger.Info("running", zap.Any("settings", appSettings))
	err = http.ListenAndServe(":80", r)
	if err != nil {
		logger.Fatal("error running server", zap.Error(err))
	}
}

func newLogger(debug bool, logLevel string) (logger *zap.Logger, err error) {
	var loggerConfig zap.Config
	if debug {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	if err := loggerConfig.Level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, err
	}

	if logger, err = loggerConfig.Build(); err != nil {
		return nil, err
	}
	return logger, nil
}
