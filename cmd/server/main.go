package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/handlers"
	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
	"go.uber.org/zap"
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
		authHandler        = handlers.NewAuth(logger, daos.User, appSettings.GitHubAuth)
	)
	r := mux.NewRouter()
	r.Handle("/vehicle", mw.Reduce(vehicleHandler.List, mw.Standard...))
	r.Handle("/vehicle/{id:[0-9]+}", mw.Reduce(vehicleHandler.Detail, mw.Standard...))
	r.Handle("/service", mw.Reduce(serviceHandler.List, mw.Standard...))
	r.Handle("/service/{id:[0-9]+}", mw.Reduce(serviceHandler.Detail, mw.Standard...))
	r.Handle("/service_type", mw.Reduce(serviceTypeHandler.List, mw.Standard...))
	r.Handle("/service_type/{id:[0-9]+}", mw.Reduce(serviceTypeHandler.Detail, mw.Standard...))
	r.Handle("/oauth/login", http.HandlerFunc(authHandler.Login))
	r.Handle("/oauth/callback", http.HandlerFunc(authHandler.Callback))

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
