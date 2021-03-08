// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Configuration.
type configuration struct {
	// Cache lifetime.
	// Store request service address lifetime, when
	// read from consul.
	// default: 3 seconds.
	CacheLifetime float64 `yaml:"cache-lifetime"`

	// Consul address.
	// default: 127.0.0.1:8500.
	ConsulAddress string `yaml:"consul-address"`

	// Consul scheme.
	// default: http.
	ConsulScheme string `yaml:"consul-scheme"`

	// Consul timeout.
	// default: 3.
	ConsulTimeout int `yaml:"consul-timeout"`

	// Sdk request method.
	// default: POST.
	ServiceRequestMethod string `yaml:"service-request-method"`

	// Sdk request content type.
	// default: application/json.
	ServiceRequestContentType string `yaml:"service-request-content-type"`

	// Sdk request timeout.
	// default: 25.
	ServiceRequestTimeout int `yaml:"service-request-timeout"`

	// Sdk request scheme if not defined.
	// default: http.
	ServiceScheme string `yaml:"service-request-scheme"`

	// Sdk request user agent.
	ServiceUserAgent string `yaml:"-"`
}

// Load defaults.
func (o *configuration) LoadDefault() {
	if o.ConsulAddress == "" {
		o.ConsulAddress = DefaultConsulAddress
	}
	if o.ConsulScheme == "" {
		o.ConsulScheme = DefaultConsulScheme
	}
	if o.ConsulTimeout == 0 {
		o.ConsulTimeout = DefaultConsulTimeout
	}
	if o.CacheLifetime == 0 {
		o.CacheLifetime = DefaultCacheLifetime
	}
	if o.ServiceRequestContentType == "" {
		o.ServiceRequestContentType = DefaultContentType
	}
	if o.ServiceRequestMethod == "" {
		o.ServiceRequestMethod = DefaultRequestMethod
	}
	if o.ServiceRequestTimeout == 0 {
		o.ServiceRequestTimeout = DefaultRequestTimeout
	}
	if o.ServiceScheme == "" {
		o.ServiceScheme = DefaultServiceScheme
	}
	if o.ServiceUserAgent == "" {
		o.ServiceUserAgent = fmt.Sprintf("%s/%s", Name, Version)
	}
}

// Load config from yaml file.
func (o *configuration) LoadYaml(file string) error {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(body, o); err != nil {
		return err
	}
	return nil
}

// Load app config from yaml file.
func (o *configuration) LoadYamlApp(file string) error {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	tmp := &struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}{}
	if err = yaml.Unmarshal(body, tmp); err != nil {
		return err
	}
	if tmp.Name != "" && tmp.Version != "" {
		o.ServiceUserAgent = fmt.Sprintf("%s/%s", tmp.Name, tmp.Version)
		return nil
	}
	return fmt.Errorf("parse app.yaml file")
}

// Initialize configuration.
func (o *configuration) initialize() {
	for _, file := range []string{"./tmp/sdk.yaml", "../tmp/sdk.yaml", "./config/sdk.yaml", "../config/sdk.yaml"} {
		if o.LoadYaml(file) != nil {
			break
		}
	}
	for _, file := range []string{"./tmp/app.yaml", "../tmp/app.yaml", "./config/app.yaml", "../config/app.yaml"} {
		if o.LoadYamlApp(file) != nil {
			break
		}
	}
	o.LoadDefault()
}
