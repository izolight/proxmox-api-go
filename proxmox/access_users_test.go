package proxmox_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/izolight/proxmox-api-go/proxmox"
)

func TestClient_GetUsers(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		equals(t, req.URL.String(), "/access/users")
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	c, err := proxmox.NewClient(server.URL, server.Client(), server.TLS)
	ok(t, err)
	users, err := c.GetUsers()
	ok(t, err)
	equals(t, []byte("OK"), users)
}