package net

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// error
var (
	ErrProxyNotConfigured = errors.New("Proxy is not configured correctly")
)

// ProxySettings todo
type ProxySettings struct {
	ProxyServer   string
	ProxyOverride string // aka no proxy
	sep           string
}

func getEnvAny(names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return ""
}

// Executor download executor
type Executor struct {
	client *http.Client
	OutDir string
}

func isTrue(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "on" || s == "yes" || s == "1"
}

// NewExecutor new executor
func NewExecutor() *Executor {
	ps, err := ResolveProxy()
	transport := &http.Transport{}
	if err == nil {
		proxyurl := ps.ProxyServer
		if !strings.Contains(proxyurl, "://") {
			proxyurl = "http://" + proxyurl // avoid proxy url parse failed
		}
		if u, err := url.Parse(proxyurl); err == nil {
			transport.Proxy = http.ProxyURL(u)
		}
	}
	bkznetdir := os.Getenv("BKZ_DOWNLOAD_OUTDIR")
	if len(bkznetdir) == 0 {
		bkznetdir = os.ExpandEnv(fmt.Sprintf("${TEMP}/bkz_net_%d", os.Getpid()))
	}
	_ = os.MkdirAll(bkznetdir, 0755)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: isTrue(os.Getenv("BKZ_INSECURE_TLS"))} //set ssl
	return &Executor{client: &http.Client{Transport: transport}}
}

func resolveFileName(resp *http.Response, rawurl string) string {
	if disp := resp.Header.Get("Content-Disposition"); disp != "" {
		if _, params, err := mime.ParseMediaType(disp); err == nil {
			if filename := params["filename"]; len(filename) > 0 {
				return filename
			}
		}
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return "index.html"
	}
	if filename := path.Base(u.Path); filename != "" && filename != "." {
		return filename
	}
	return "index.html"
}

// Get get file from network
func (e *Executor) Get(rawurl string) (string, error) {
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "bkz/1.0")
	resp, err := e.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	filename := resolveFileName(resp, rawurl)
	fullname := filepath.Join(e.OutDir, filename)
	fd, err := os.Create(fullname)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	io.Copy(fd, resp.Body)
	return fullname, nil
}
