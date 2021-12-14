package proxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type ProxyGetter interface {
	GetProxy() (string, error)    // return an usable proxy. if there is not an usable proxy, return "", error
	CheckProxy(proxy string) bool // check if the proxy is usable
	EraseProxy(proxy string) int  // erase the proxy from the proxy list
	PushProxy(proxy ...string)    // push the proxy into the proxy list
	LenOfProxies() int
}

type DefaultProxyGetter struct {
	now                 int // return an proxy in sequence, without using random numbers
	proxies             []string
	CheckBeforeGetProxy bool // if true, proxy will be checked for availability before returning; otherwise, it will be returned directly
	lock                sync.Mutex
}

func (p *DefaultProxyGetter) GetProxy() (string, error) {
	if p.proxies == nil || len(p.proxies) == 0 {
		return "", nil
	}
	proxy := p.proxies[p.now]
	p.now = (p.now + 1) % len(p.proxies)
	if p.CheckBeforeGetProxy {
		for len(p.proxies) > 0 && !p.CheckProxy(proxy) {
			p.EraseProxy(proxy)
			proxy = p.proxies[p.now]
			p.now = (p.now + 1) % len(p.proxies)
		}
	}
	if p.proxies == nil || len(p.proxies) == 0 {
		return "", nil
	}
	return proxy, nil
}

// The efficiency is not high when the number of proxies is large
func (p *DefaultProxyGetter) EraseProxy(proxy string) int {
	p.lock.Lock()
	for i, v := range p.proxies {
		if v == proxy {
			p.proxies = append(p.proxies[:i], p.proxies[i+1:]...)
			break
		}
	}
	p.now -= 1
	p.lock.Unlock()
	return len(p.proxies)
}

func (p *DefaultProxyGetter) PushProxy(proxy ...string) {
	p.lock.Lock()
	limitCh := make(chan struct{}, 30) //限制并发数
	wg := sync.WaitGroup{}
	wg.Add(len(proxy))
	for _, v := range proxy {
		limitCh <- struct{}{}
		go func(proxyAddr string) {
			if p.CheckProxy(proxyAddr) && p.CheckExist(proxyAddr) {
				p.proxies = append(p.proxies, proxyAddr)
			}
			<-limitCh
			wg.Done()
		}(v)
	}
	wg.Wait()
	p.lock.Unlock()
}

func (p *DefaultProxyGetter) CheckProxy(proxyAddr string) bool {
	httpUrl := "http://icanhazip.com"
	proxy, _ := url.Parse(proxyAddr)

	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	res, err := httpClient.Get(httpUrl)
	if err != nil {
		//fmt.Println("错误信息：",err)
		return false
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return false
	}
	c, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK || string(c) == "" {
		return false
	}
	return true
}

func (p *DefaultProxyGetter) CheckExist(proxyAddr string) bool {
	for _, v := range p.proxies {
		if v == proxyAddr {
			return false
		}
	}
	return true
}

func (p *DefaultProxyGetter) LenOfProxies() int {
	return len(p.proxies)
}
