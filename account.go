package putio

import (
	"context"
	"net/http"
)

// AccountService is the service to gather information about user account.
type AccountService struct {
	client *Client
}

// Info retrieves user account information.
func (a *AccountService) Info(ctx context.Context) (AccountInfo, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, "/v2/account/info", nil)
	if err != nil {
		return AccountInfo{}, err
	}

	var r struct {
		Info AccountInfo
	}
	_, err = a.client.Do(req, &r) // nolint:bodyclose
	if err != nil {
		return AccountInfo{}, err
	}
	return r.Info, nil
}

// Settings retrieves user preferences.
func (a *AccountService) Settings(ctx context.Context) (Settings, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, "/v2/account/settings", nil)
	if err != nil {
		return Settings{}, err
	}
	var r struct {
		Settings Settings
	}
	_, err = a.client.Do(req, &r) // nolint:bodyclose
	if err != nil {
		return Settings{}, err
	}

	return r.Settings, nil
}
