package password

import (
	"time"

	"github.com/labstack/echo"

	andro "github.com/soldevx/kuiper/andro/pkg/utl/model"

	"github.com/soldevx/kuiper/andro/pkg/api/password"
)

// New creates new password logging service
func New(svc password.Service, logger andro.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents password logging service
type LogService struct {
	password.Service
	logger andro.Logger
}

const name = "password"

// Change logging
func (ls *LogService) Change(c echo.Context, id int, oldPass, newPass string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Change password request", err,
			map[string]interface{}{
				"req":  id,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Change(c, id, oldPass, newPass)
}
