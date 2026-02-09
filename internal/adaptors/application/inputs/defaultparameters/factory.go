// Copyright 2026 The MathWorks, Inc.

package defaultparameters

import (
	"github.com/matlab/matlab-mcp-core-server/internal/entities"
	"github.com/matlab/matlab-mcp-core-server/internal/messages"
)

type MessageCatalog interface {
	Get(message messages.MessageKey) string
}

type Factory struct {
	messageCatalog MessageCatalog
}

func NewFactory(
	messageCatalog MessageCatalog,
) *Factory {
	return &Factory{
		messageCatalog: messageCatalog,
	}
}

func (f *Factory) DefaultParameters() []entities.Parameter {
	parameterDefs := []parameterWithDescriptionFromMessageCatalog{
		HelpMode(),
		VersionMode(),
		DisableTelemetry(),
		PreferredLocalMATLABRoot(),
		PreferredMATLABStartingDirectory(),
		BaseDir(),
		LogLevel(),
		InitializeMATLABOnStartup(),
		MATLABDisplayMode(),
		UseSingleMATLABSession(),
		WatchdogMode(),
		ServerInstanceID(),
	}

	parameters := make([]entities.Parameter, len(parameterDefs))
	for i, parameterDef := range parameterDefs {
		f.resolveDescription(parameterDef)
		parameters[i] = parameterDef
	}
	return parameters
}

func (f *Factory) resolveDescription(p parameterWithDescriptionFromMessageCatalog) {
	p.setDescription(f.messageCatalog.Get(p.getDescriptionKey()))
}
