package security

import (
	"context"

	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/transport/internet"
)

func CreateSecurityEngineFromSettings(context context.Context, settings *internet.MemoryStreamConfig) (Engine, error) {
	if settings == nil || settings.SecurityType == "" {
		return nil, nil
	}
	securityEngine, err := common.CreateObject(context, settings.SecuritySettings)
	if err != nil {
		return nil, newError("unable to create security engine from security settings").Base(err)
	}
	securityEngineTyped, ok := securityEngine.(Engine)
	if !ok {
		return nil, newError("type assertion error when create security engine from security settings")
	}
	return securityEngineTyped, nil
}
