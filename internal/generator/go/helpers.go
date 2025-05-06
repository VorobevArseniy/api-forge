package generator

import "api-generator/pkg/spec"

type EndpointTemplate struct {
	Name     string
	Request  map[string]string
	Response map[string]string
}

func convertEndpoints(endpoints map[string]spec.Endpoint) []EndpointTemplate {
	var result []EndpointTemplate
	for name, ep := range endpoints {
		result = append(result, EndpointTemplate{
			Name:     name,
			Request:  ep.Request,
			Response: ep.Response,
		})
	}
	return result
}
