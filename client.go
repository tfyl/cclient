package cclient

import (
	"github.com/useflyent/fhttp/cookiejar"
	"time"

	http "github.com/useflyent/fhttp"
	"golang.org/x/net/proxy"

	utls "github.com/tfyl/utls"
)

//NewClient creates new http client with cookie jar
func NewClient(clientHello utls.ClientHelloID, userAgent string, proxyUrl string, allowRedirect bool, timeout time.Duration) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	if len(proxyUrl) > 0 {
		dialer, err := newConnectDialer(userAgent, proxyUrl)
		if err != nil {
			if allowRedirect {
				return &http.Client{
					Jar:     jar,
					Timeout: time.Second * timeout,
				}, err
			}
			return &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Jar:     jar,
				Timeout: time.Second * timeout,
			}, err
		}
		if allowRedirect {
			return &http.Client{
				Transport: newRoundTripper(clientHello, dialer),
				Jar:       jar,
				Timeout:   time.Second * timeout,
			}, nil
		}
		return &http.Client{
			Transport: newRoundTripper(clientHello, dialer),
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Jar:     jar,
			Timeout: time.Second * timeout,
		}, nil
	} else {
		if allowRedirect {
			return &http.Client{
				Transport: newRoundTripper(clientHello, proxy.Direct),
				Jar:       jar,
				Timeout:   time.Second * timeout,
			}, nil
		}
		return &http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct),
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Jar:     jar,
			Timeout: time.Second * timeout,
		}, nil

	}
}
