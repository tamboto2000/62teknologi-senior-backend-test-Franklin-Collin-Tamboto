package zaplog

import (
	"os"

	"github.com/tamboto2000/62teknologi-senior-backend-test-Franklin-Collin-Tamboto/constants"
	"go.uber.org/zap"
)

// New() instatiate logger from go.uber.org/zap
func New() (*zap.Logger, error) {
	var conf zap.Config
	mode := os.Getenv(constants.DeployMode)

	if mode == constants.DeployModeDev {
		conf = zap.NewDevelopmentConfig()
	} else if mode == constants.DeployModeProd {
		conf = zap.NewProductionConfig()
	}

	if _, err := os.Stat(constants.LogFile); err == os.ErrNotExist {
		f, err := os.Create(constants.LogFile)
		if err != nil {
			return nil, err
		}

		defer f.Close()
	}

	return conf.Build()
}
