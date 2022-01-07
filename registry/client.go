package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

func RegisterService(r Registration) error {

	serviceUpdateURL, err := url.Parse(r.ServiceUpdateUrl)
	if err != nil {
		return err
	}
	http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{})
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	err = enc.Encode(r)

	if err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", buf)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register %v", res.StatusCode)
	}

	return nil
}

type serviceUpdateHandler struct{}

func (suh serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prov.Update(p)
}

func ShutDownService(serviceUrl string) error {
	req, err := http.NewRequest(http.MethodDelete, serviceUrl, bytes.NewBuffer([]byte(serviceUrl)))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service %v", serviceUrl)
	}
	return err
}

type Providers struct {
	services map[ServiceName][]string
	mutex    *sync.RWMutex
}

func (p *Providers) Update(pat patch) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, v := range pat.Added {
		if _, ok := p.services[v.Name]; !ok {
			p.services[v.Name] = make([]string, 0)
		}
		p.services[v.Name] = append(p.services[v.Name], v.URL)
	}

	for _, v := range pat.Removed {
		if providerUrls, ok := p.services[v.Name]; ok {
			for i := range providerUrls {
				if providerUrls[i] == v.URL {
					p.services[v.Name] = append(providerUrls[:i], providerUrls[i+1:]...)
				}
			}
		}
	}
}

func (p Providers) get(name ServiceName) (string, error) {
	providers, ok := p.services[name]

	if !ok {
		return "", fmt.Errorf("no providers available")
	}
	idx := int(rand.Float32() * float32(len(providers)))
	return providers[idx], nil
}

func GetProviders(name ServiceName) (string, error) {
	return prov.get(name)
}

var prov = Providers{
	services: make(map[ServiceName][]string, 0),
	mutex:    new(sync.RWMutex),
}
