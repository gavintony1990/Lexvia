package registry

import "github.com/QuantumNous/new-api/relay/channel"

var adaptorFactories = map[int]func() channel.Adaptor{}

func RegisterAdaptor(apiType int, factory func() channel.Adaptor) {
	adaptorFactories[apiType] = factory
}

func GetAdaptor(apiType int) channel.Adaptor {
	if factory, ok := adaptorFactories[apiType]; ok {
		return factory()
	}
	return nil
}
