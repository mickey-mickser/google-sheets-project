package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

// fetchToken запрашивает у Metadata Server OAuth2 access_token с заданными scope.
func fetchToken(ctx context.Context, scopes []string) (*oauth2.Token, error) {
	url := "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token?scopes=" +
		strings.Join(scopes, ",")
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating metadata request: %w", err)
	}
	req.Header.Add("Metadata-Flavor", "Google")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting metadata token: %w", err)
	}
	defer resp.Body.Close()

	// Декодируем JSON-ответ в oauth2.Token
	var tok oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return nil, fmt.Errorf("decoding metadata token: %w", err)
	}
	return &tok, nil
}

// metadataTokenSource возвращает oauth2.TokenSource, который опрашивает Metadata Server
// и автоматически обновляет токен с нужными scope при каждом использовании.
func metadataTokenSource(ctx context.Context, scopes []string) oauth2.TokenSource {
	return oauth2.ReuseTokenSource(
		nil,
		oauth2.TokenSource(oauth2.StaticTokenSource(&oauth2.Token{})),
	)
}
