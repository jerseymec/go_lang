package domain

import (
	"fmt"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	config map[string]interface{}
	lock   sync.RWMutex
}

func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}
	//fmt.Printf("%T",rawConfig)
	unknownconfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("config is not a map")
	}
	config, err := convertKeystoString(unknownconfig)
	if err != nil {
		return err
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	c.config = config
	return nil
}

//returns the config for a requested service(name)
func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	//check the yaml root if its the service name thats needed
	a, ok := c.config["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("base config is not a map")
	}
	//if no config is defined for the service
	if _, ok = c.config[serviceName]; !ok {
		//return the base config
		return a, nil
	}
	b, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service %q config is not a map", serviceName)
	}
	config := make(map[string]interface{})
	for k, v := range a {
		config[k] = v
	}
	for k, v := range b {
		config[k] = v
	}
	return config, nil
}

func convertKeystoString(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("Config key is not a string")
		}
		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeystoString(vMap)
			if err != nil {
				return nil, err
			}
		}
		n[str] = v
	}
	return n, nil
}
