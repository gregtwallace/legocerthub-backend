package authorizations

import (
	"errors"
	"legocerthub-backend/pkg/acme"
	"legocerthub-backend/pkg/domain/challenges"

	"go.uber.org/zap"
)

var errServiceComponent = errors.New("necessary authorizations service component is missing")

// App interface is for connecting to the main app
type App interface {
	GetLogger() *zap.SugaredLogger
	GetAcmeProdService() *acme.Service
	GetAcmeStagingService() *acme.Service
	GetChallengesService() *challenges.Service
}

// service struct
type Service struct {
	logger      *zap.SugaredLogger
	acmeProd    *acme.Service
	acmeStaging *acme.Service
	challenges  *challenges.Service
}

// NewService creates a new service
func NewService(app App) (*Service, error) {
	service := new(Service)

	// logger
	service.logger = app.GetLogger()
	if service.logger == nil {
		return nil, errServiceComponent
	}

	// acme services
	service.acmeProd = app.GetAcmeProdService()
	if service.acmeProd == nil {
		return nil, errServiceComponent
	}
	service.acmeStaging = app.GetAcmeStagingService()
	if service.acmeStaging == nil {
		return nil, errServiceComponent
	}

	// challenge solver service
	service.challenges = app.GetChallengesService()
	if service.challenges == nil {
		return nil, errServiceComponent
	}

	return service, nil
}
