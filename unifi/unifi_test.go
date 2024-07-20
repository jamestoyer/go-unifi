package unifi

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func TestClient_DoRaw(t *testing.T) {
	ctx := context.Background()
	client := Client{}

	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,

		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	jar, _ := cookiejar.New(nil)
	httpClient.Jar = jar

	_ = client.SetHTTPClient(httpClient)

	require.NoError(t, client.SetBaseURL("https://localhost:8443"))
	require.NoError(t, client.Login(ctx, "admin", "admin"))

	respBytes, err := client.DoRaw(ctx, "GET", fmt.Sprintf("s/%s/stat/device/%s", "default", "00:27:22:00:00:01"), nil)
	assert.NoError(t, err)
	spew.Print(string(respBytes))
}
