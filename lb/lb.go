package lb

import (
  "net/http"
  "net/http/httputil"
  "net/url"
  "fmt"
)

type LoadBalancer struct {
  proxies []*httputil.ReverseProxy
  hosts []string
  current int
}

func NewLoadBalancer(hosts ...string) (*LoadBalancer, error) {
  b := new(LoadBalancer)
  for _, host := range hosts {
    u, _ := url.Parse(host)
    b.proxies = append(b.proxies, httputil.NewSingleHostReverseProxy(u))
    b.hosts = append(b.hosts, host)
  }
  return b, nil
}

func (l *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  l.current = (l.current + 1) % len(l.proxies)
  fmt.Printf("Fowarding to %s\n\n", l.hosts[l.current])
  l.proxies[l.current].ServeHTTP(w, r)
}

