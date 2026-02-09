// Copyright 2025-2026 The MathWorks, Inc.

package config

import (
	"reflect"

	"github.com/matlab/matlab-mcp-core-server/internal/adaptors/application/inputs/defaultparameters"
	"github.com/matlab/matlab-mcp-core-server/internal/messages"
)

func get[OutputType any](cfg GenericConfig, parameter *defaultparameters.ParameterDef[OutputType]) (OutputType, messages.Error) {
	var zeroValue OutputType

	value, err := cfg.Get(parameter.GetID())
	if err != nil {
		return zeroValue, err
	}

	castValue, ok := value.(OutputType)
	if !ok {
		expectedType := reflect.TypeOf(zeroValue).String()
		return zeroValue, messages.New_StartupErrors_InvalidParameterType_Error(parameter.GetID(), expectedType)
	}

	return castValue, nil
}
