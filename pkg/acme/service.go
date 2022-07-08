package acme

import (
	"errors"
	"legocerthub-backend/pkg/acme/nonces"

	"go.uber.org/zap"
)

// App interface is for connecting to the main app
type App interface {
	//GetAccountStorage() Storage
	GetLogger() *zap.SugaredLogger
}

// Acme service struct
type Service struct {
	logger       *zap.SugaredLogger
	dirUri       string
	dir          *acmeDirectory
	nonceManager *nonces.Manager
}

// NewService creates a new acme service based on a directory uri
func NewService(app App, dirUri string) (*Service, error) {
	service := new(Service)
	var err error

	// logger
	service.logger = app.GetLogger()
	if service.logger == nil {
		return nil, errors.New("acme: newservice requires valid logger")
	}

	// acme directory
	service.dirUri = dirUri
	service.dir = new(acmeDirectory)

	// initial population
	err = service.fetchAcmeDirectory()
	if err != nil {
		return nil, err
	}

	// start go routine to check for periodic updates
	service.backgroundDirManager()

	// nonce manager
	service.nonceManager = nonces.NewManager(&service.dir.NewNonce)

	return service, nil
}
