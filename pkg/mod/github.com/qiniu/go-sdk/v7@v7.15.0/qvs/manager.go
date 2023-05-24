package qvs

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/client"
)

// Manager 代表一个 qvs 用户的客户端
type Manager struct {
	client *client.Client
	mac    *auth.Credentials
}

// New 初始化 Client.
func NewManager(mac *auth.Credentials, tr http.RoundTripper) *Manager {
	client := client.DefaultClient
	client.Transport = newTransport(mac, nil)
	return &Manager{
		client: &client,
		mac:    mac,
	}
}

func setQuery(q url.Values, key string, v interface{}) {
	q.Set(key, fmt.Sprint(v))
}

func (manager *Manager) url(format string, args ...interface{}) string {
	return APIHTTPScheme + APIHost + fmt.Sprintf(format, args...)
}
