// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fuyibing/log/v2"
	"github.com/hashicorp/consul/api"
)

// Consul struct.
type consul struct {
	err    error                // consul error.
	client *api.Client          // consul client.
	caches map[string]string    // cached services.
	times  map[string]time.Time // cached times.
	mu     *sync.RWMutex        // cached mutex.
}

// Get service.
func (o *consul) Get(ctx interface{}, service string) (string, error) {
	// Has error.
	if o.err != nil {
		return "", o.err
	}
	// From cache.
	if host, ok := o.fromCache(ctx, service); ok {
		return host, nil
	}
	// From server.
	return o.fromServer(ctx, service)
}

// Get service from server.
func (o *consul) GetServer(ctx interface{}, service string) (string, error) {
	return o.fromServer(ctx, service)
}

// Read service from cache.
func (o *consul) fromCache(ctx interface{}, service string) (host string, has bool) {
	// Begin.
	if log.Config.DebugOn() {
		log.Client.Debugfc(ctx, "[sdk=%s] cache begin.", service)
	}
	// Lock and unlock when completed.
	o.mu.RLock()
	defer func() {
		o.mu.RUnlock()
		// log with host.
		if has {
			log.Client.Infofc(ctx, "[sdk=%s] cache found {%s}.", service, host)
		} else if log.Config.DebugOn() {
			log.Client.Debugfc(ctx, "[sdk=%s] cache not found.", service)
		}
	}()
	// Return from cached.
	if t, ok := o.times[service]; ok {
		if time.Now().Sub(t).Seconds() <= Config.CacheLifetime {
			host, ok = o.caches[service]
			return host, true
		}
	}
	// Return for not found.
	return "", false
}

// Read service from server.
func (o *consul) fromServer(ctx interface{}, service string) (host string, err error) {
	// begin.
	if log.Config.DebugOn() {
		log.Client.Debugfc(ctx, "[sdk=%s] read from {%s} begin.", service, Config.Address)
	}
	// ended.
	t := time.Now()
	defer func() {
		d := time.Now().Sub(t).Seconds()
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			log.Client.Errorfc(ctx, "[sdk=%s][d=%f] read from {%s} error: %v.", service, d, Config.Address, r)
		} else if log.Config.InfoOn() {
			log.Client.Infofc(ctx, "[sdk=%s][d=%f] read from {%s} and found {%s}.", service, d, Config.Address, host)
		}
	}()
	// prepare.
	var se *api.ServiceEntry
	var ss []*api.ServiceEntry
	// read server.
	if ss, _, err = o.client.Health().Service(service, "", true, nil); err != nil {
		return
	}
	// fetch one.
	se = ss[rand.Intn(len(ss))]
	host = fmt.Sprintf("%s:%d", se.Service.Address, se.Service.Port)
	// add scheme prefix to service.
	if !regexpServiceAddress.MatchString(host) {
		host = fmt.Sprintf("%s://%s", ServiceScheme, host)
	}
	// update cache.
	o.mu.Lock()
	defer o.mu.Unlock()
	o.caches[service] = host
	o.times[service] = time.Now()
	return
}

// Initialize consul.
func (o *consul) initialize() {
	o.caches = make(map[string]string)
	o.times = make(map[string]time.Time)
	o.mu = new(sync.RWMutex)
	o.client, o.err = api.NewClient(&api.Config{
		Address:  Config.Address,
		Scheme:   Config.Scheme,
		WaitTime: time.Duration(Config.Timeout) * time.Second,
	})
}
