package cclient

import (
	"time"

	http "github.com/useflyent/fhttp"
	"golang.org/x/net/proxy"

	utls "github.com/Titanium-ctrl/utls"
)

func NewClient(clientHello utls.ClientHelloID, proxyUrl string, allowRedirect bool, timeout time.Duration) (http.Client, error) {
	if len(proxyUrl) > 0 {
		dialer, err := newConnectDialer(proxyUrl)
		if err != nil {
			if allowRedirect {
				return http.Client{
					Timeout: time.Second * timeout,
				}, err
			}
			return http.Client{
				Timeout: time.Second * timeout,
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}, err
		}
		if allowRedirect {
			return http.Client{
				Transport: newRoundTripper(clientHello, dialer),
				Timeout:   time.Second * timeout,
			}, nil
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, dialer),
			Timeout:   time.Second * timeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}, nil
	} else {
		if allowRedirect {
			return http.Client{
				Transport: newRoundTripper(clientHello, proxy.Direct),
				Timeout:   time.Second * timeout,
			}, nil
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct),
			Timeout:   time.Second * timeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}, nil

	}
}
