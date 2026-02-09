// Copyright 2026 The MathWorks, Inc.

package defaultparameters

import (
	"github.com/matlab/matlab-mcp-core-server/internal/entities"
	"github.com/matlab/matlab-mcp-core-server/internal/messages"
)

type parameterWithDescriptionFromMessageCatalog interface {
	entities.Parameter
	setDescription(description string)
	getDescriptionKey() messages.MessageKey
}

type ParameterDef[ValueType any] struct {
	id             string
	flagName       string
	hiddenFlag     bool
	envVarName     string
	descriptionKey messages.MessageKey
	description    string
	defaultValue   ValueType
	recordToLog    bool
}

func (p ParameterDef[ValueType]) GetID() string {
	return p.id
}

func (p ParameterDef[ValueType]) GetFlagName() string {
	return p.flagName
}

func (p ParameterDef[ValueType]) GetHiddenFlag() bool {
	return p.hiddenFlag
}

func (p ParameterDef[ValueType]) GetEnvVarName() string {
	return p.envVarName
}

func (p ParameterDef[ValueType]) GetDescription() string {
	return p.description
}

func (p ParameterDef[ValueType]) GetDefaultValue() any {
	return p.defaultValue
}

func (p *ParameterDef[ValueType]) setDescription(description string) {
	p.description = description
}

func (p ParameterDef[ValueType]) getDescriptionKey() messages.MessageKey {
	return p.descriptionKey
}

func (p ParameterDef[ValueType]) GetRecordToLog() bool {
	return p.recordToLog
}
