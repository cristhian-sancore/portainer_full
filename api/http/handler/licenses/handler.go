package licenses

import (
	"net/http"
	"time"

	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/http/security"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/response"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler for license operations.
type Handler struct {
	*mux.Router
}

// LicenseInfo matches the structure expected by the frontend
type LicenseInfo struct {
	ProductEdition portainer.SoftwareEdition `json:"productEdition"`
	Company        string                    `json:"company"`
	Email          string                    `json:"email"`
	CreatedAt      int64                     `json:"createdAt"`
	ExpiresAt      int64                     `json:"expiresAt"`
	Nodes          int                       `json:"nodes"`
	Type           int                       `json:"type"`
	Valid          bool                      `json:"valid"`
	EnforcedAt     int64                     `json:"enforcedAt"`
	Enforced       bool                      `json:"enforced"`
}

// License matches the individual license structure expected by the frontend
type License struct {
	ID             string                    `json:"id"`
	Company        string                    `json:"company"`
	Created        int64                     `json:"created"`
	Email          string                    `json:"email"`
	ExpiresAfter   int64                     `json:"expiresAfter"`
	LicenseKey     string                    `json:"licenseKey"`
	Nodes          int                       `json:"nodes"`
	ProductEdition portainer.SoftwareEdition `json:"productEdition"`
	Revoked        bool                      `json:"revoked"`
	RevokedAt      int64                     `json:"revokedAt"`
	Type           int                       `json:"type"`
	Version        int                       `json:"version"`
	Reference      string                    `json:"reference"`
	ExpiresAt      int64                     `json:"expiresAt"`
}

// NewHandler creates a handler to manage license operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}

	router := h.PathPrefix("/licenses").Subrouter()

	publicRouter := router.PathPrefix("/").Subrouter()
	publicRouter.Use(bouncer.PublicAccess)

	publicRouter.Handle("/info", httperror.LoggerHandler(h.licenseInfo)).Methods(http.MethodGet)
	publicRouter.Handle("", httperror.LoggerHandler(h.listLicenses)).Methods(http.MethodGet)
	publicRouter.Handle("", httperror.LoggerHandler(h.attachLicense)).Methods(http.MethodPost)
	publicRouter.Handle("/remove", httperror.LoggerHandler(h.removeLicense)).Methods(http.MethodPost)

	return h
}

func (handler *Handler) licenseInfo(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	now := time.Now()
	// Return a fake valid EE subscription license
	info := LicenseInfo{
		ProductEdition: portainer.PortainerEE,
		Company:        "SANCORE",
		Email:          "admin@sancore.com.br",
		CreatedAt:      now.AddDate(-1, 0, 0).Unix(),
		ExpiresAt:      now.AddDate(10, 0, 0).Unix(), // Expires in 10 years
		Nodes:          9999,
		Type:           2, // Subscription
		Valid:          true,
		EnforcedAt:     0,
		Enforced:       false,
	}

	return response.JSON(w, info)
}

func (handler *Handler) listLicenses(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	now := time.Now()
	licenses := []License{
		{
			ID:             "sancore-full-unlock-001",
			Company:        "SANCORE",
			Created:        now.AddDate(-1, 0, 0).Unix(),
			Email:          "admin@sancore.com.br",
			ExpiresAfter:   now.AddDate(10, 0, 0).Unix(),
			LicenseKey:     "3-SANCORE-FULL-2099-ACTIVA",
			Nodes:          9999,
			ProductEdition: portainer.PortainerEE,
			Revoked:        false,
			RevokedAt:      0,
			Type:           2, // Subscription
			Version:        2,
			Reference:      "SANCORE-PORTAINER-FULL",
			ExpiresAt:      now.AddDate(10, 0, 0).Unix(),
		},
	}

	return response.JSON(w, licenses)
}

func (handler *Handler) attachLicense(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	now := time.Now()
	// Accept any license and return success
	result := map[string]interface{}{
		"licenses": []License{
			{
				ID:             "sancore-full-unlock-001",
				Company:        "SANCORE",
				Created:        now.Unix(),
				Email:          "admin@sancore.com.br",
				ExpiresAfter:   now.AddDate(10, 0, 0).Unix(),
				LicenseKey:     "3-SANCORE-FULL-2099-ACTIVA",
				Nodes:          9999,
				ProductEdition: portainer.PortainerEE,
				Revoked:        false,
				RevokedAt:      0,
				Type:           2,
				Version:        2,
				Reference:      "SANCORE-PORTAINER-FULL",
				ExpiresAt:      now.AddDate(10, 0, 0).Unix(),
			},
		},
		"failedKeys": map[string]string{},
	}

	return response.JSON(w, result)
}

func (handler *Handler) removeLicense(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	result := map[string]interface{}{
		"failedKeys": map[string]string{},
	}

	return response.JSON(w, result)
}
