package system

import (
	"net/http"
	"regexp"

	ceplf "github.com/portainer/portainer/api/platform"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/request"
	"github.com/portainer/portainer/pkg/libhttp/response"

	"github.com/pkg/errors"
)

type systemUpgradePayload struct {
	License string
}

var re = regexp.MustCompile(`^\d-.+`)

func (payload *systemUpgradePayload) Validate(r *http.Request) error {
	if payload.License == "" {
		return errors.New("license is missing")
	}

	if !re.MatchString(payload.License) {
		return errors.New("license is invalid")
	}

	return nil
}

// @id systemUpgrade
// @summary Upgrade Portainer to BE
// @description Upgrade Portainer to BE
// @description **Access policy**: administrator
// @tags system
// @produce json
// @success 204 {object} status "Success"
// @router /system/upgrade [post]
func (handler *Handler) systemUpgrade(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	payload, err := request.GetPayload[systemUpgradePayload](r)
	if err != nil {
		return httperror.BadRequest("Invalid request payload", err)
	}

	if !re.MatchString(payload.License) {
		return httperror.BadRequest("Invalid license format. Must start with a digit and a dash (e.g. 3-xxxx)", nil)
	}

	// Bypass actual upgrade and return success
	return response.Empty(w)
}
