// Copyright 2025-2026 The MathWorks, Inc.

package config

import (
	"github.com/matlab/matlab-mcp-core-server/internal/adaptors/application/inputs/defaultparameters"
	"github.com/matlab/matlab-mcp-core-server/internal/messages"
)

func Get[OutputType any](cfg GenericConfig, parameter *defaultparameters.ParameterDef[OutputType]) (OutputType, messages.Error) {
	return get(cfg, parameter)
}
