package jsonified

import (
	"github.com/v4fly/v4ray-core/v0/app/subscription/containers"
	"github.com/v4fly/v4ray-core/v0/app/subscription/containers/jsonfieldarray"
	"github.com/v4fly/v4ray-core/v0/common"
	jsonConf "github.com/v4fly/v4ray-core/v0/infra/conf/json"
)

func newJsonifiedYamlParser() containers.SubscriptionContainerDocumentParser {
	return &jsonifiedYAMLParser{}
}

type jsonifiedYAMLParser struct{}

func (j jsonifiedYAMLParser) ParseSubscriptionContainerDocument(rawConfig []byte) (*containers.Container, error) {
	parser := jsonfieldarray.NewJSONFieldArrayParser()
	jsonified, err := jsonConf.FromYAML(rawConfig)
	if err != nil {
		return nil, newError("failed to parse as yaml").Base(err)
	}
	container, err := parser.ParseSubscriptionContainerDocument(jsonified)
	if err != nil {
		return nil, newError("failed to parse as jsonfieldarray").Base(err)
	}
	container.Kind = "Yaml2Json+" + container.Kind

	for _, value := range container.ServerSpecs {
		value.KindHint = "Yaml2Json+" + value.KindHint
	}
	return container, nil
}

func init() {
	common.Must(containers.RegisterParser("Yaml2Json", newJsonifiedYamlParser()))
}
