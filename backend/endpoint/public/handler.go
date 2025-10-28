package publicEndpoint

import (
	"backend/common/config"
	"backend/type/common"
	"context"

	"net/url"

	"github.com/bsthun/gut"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Handler struct {
	config       *config.Config
	database     common.Database
	OidcProvider *oidc.Provider
	OidcVerifier *oidc.IDTokenVerifier
	Oauth2Config *oauth2.Config
}

func Handle(
	config *config.Config,
	database common.Database,
) *Handler {
	handler := &Handler{
		config:       config,
		database:     database,
		OidcProvider: nil,
		OidcVerifier: nil,
		Oauth2Config: nil,
	}

	redirectUrl, err := url.JoinPath(*config.FrontendUrl, "/entry/callback")
	if err != nil {
		gut.Fatal("unable to join url path", err)
	}

	handler.OidcProvider, err = oidc.NewProvider(context.Background(), *config.OauthEndpoint)
	if err != nil {
		gut.Fatal("unable to create oidc provider", err)
	}

	handler.Oauth2Config = &oauth2.Config{
		ClientID:     *config.OauthClientId,
		ClientSecret: *config.OauthClientSecret,
		RedirectURL:  redirectUrl,
		Endpoint:     handler.OidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "email", "profile"},
	}

	return handler
}
