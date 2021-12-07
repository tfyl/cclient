package cclient

import (
	"net/url"
	"strings"
)

//authProxy proxy with user:pass auth
type authProxy struct {
	username string
	password string
	ip       string
	port     string
	raw      string
}

//ipProxy proxy without user:pass auth
type ipProxy struct {
	ip   string
	port string
	raw  string
}

func (proxy *ipProxy) Url() (u *url.URL) {
	u, _ = url.Parse("http://" + proxy.ip + ":" + proxy.port)
	return
}
func (proxy *ipProxy) Raw() string {
	return proxy.raw
}

func (proxy *ipProxy) Save(s string) {
	proxy.raw = s
}

func (proxy *authProxy) Url() (u *url.URL) {
	u, _ = url.Parse("http://" + proxy.username + ":" + proxy.password + "@" + proxy.ip + ":" + proxy.port)
	return u
}

func (proxy *authProxy) Raw() string {
	return proxy.raw
}

func (proxy *authProxy) Save(s string) {
	proxy.raw = s
}

func (proxy *ipProxy) makeProxy(ip, port string) error {
	_, err := url.Parse("http://" + ip + ":" + port)
	if err != nil {
		return err
	} else {
		proxy.ip = ip
		proxy.port = port
		return nil
	}
}

func (proxy *authProxy) makeProxy(ip, port, username, password string) error {
	_, err := url.Parse("http://" + username + ":" + password + "@" + ip + ":" + port)
	if err != nil {
		return err
	} else {
		proxy.username = username
		proxy.password = password
		proxy.ip = ip
		proxy.port = port
		return nil
	}
}

// MakeProxy makes proxy struct
// must be in ip:port:user:pass or ip:port
// can use proxy.Url().String() to use in client
func MakeProxy(u string) (Proxy, error) {
	uslice := strings.Split(u, ":")
	if len(uslice) == 2 {
		var proxy ipProxy
		proxy.Save(u)
		err := proxy.makeProxy(uslice[0], uslice[1])

		if err != nil {
			return nil, err
		}

		return &proxy, nil
	}

	if len(uslice) == 4 {
		var proxy authProxy
		proxy.Save(u)
		err := proxy.makeProxy(uslice[0], uslice[1], uslice[2], uslice[3])

		if err != nil {
			return nil, err
		}

		return &proxy, nil
	}

	return nil, nil
}

type Proxy interface {
	Url() *url.URL
	Raw() string
}
