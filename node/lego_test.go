package node

import (
	"os"
	"testing"

	"github.com/InazumaV/V2bX/conf"
)

func TestLego_CreateCertByDns(t *testing.T) {
	token := os.Getenv("CF_DNS_API_TOKEN")
	if token == "" {
		t.Skip("CF_DNS_API_TOKEN not set; skipping DNS integration test")
	}

	l, err := NewLego(&conf.CertConfig{
		CertMode:   "dns",
		Email:      "test@test.com",
		CertDomain: "test.test.com",
		Provider:   "cloudflare",
		DNSEnv: map[string]string{
			"CF_DNS_API_TOKEN": token,
		},
		CertFile: "./cert/1.pem",
		KeyFile:  "./cert/1.key",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = l.CreateCert()
	if err != nil {
		t.Error(err)
	}
}

func TestLego_RenewCert(t *testing.T) {
	t.Skip("DNS renewal integration test skipped in automated runs")
}
