// Copyright 2026 The MathWorks, Inc.

package defaultparameters

import "github.com/matlab/matlab-mcp-core-server/internal/messages"

func HelpMode() *ParameterDef[bool] {
	return &ParameterDef[bool]{
		id:             "HelpMode",
		flagName:       "help",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_HelpDescription,
		defaultValue:   false,
		recordToLog:    false,
	}
}

func VersionMode() *ParameterDef[bool] {
	return &ParameterDef[bool]{
		id:             "VersionMode",
		flagName:       "version",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_VersionDescription,
		defaultValue:   false,
		recordToLog:    false,
	}
}

func DisableTelemetry() *ParameterDef[bool] {
	return &ParameterDef[bool]{
		id:             "DisableTelemetry",
		flagName:       "disable-telemetry",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_DisableTelemetryDescription,
		defaultValue:   false,
		recordToLog:    true,
	}
}

func PreferredLocalMATLABRoot() *ParameterDef[string] {
	return &ParameterDef[string]{
		id:             "PreferredLocalMATLABRoot",
		flagName:       "matlab-root",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_PreferredLocalMATLABRootDescription,
		defaultValue:   "",
		recordToLog:    true,
	}
}

func PreferredMATLABStartingDirectory() *ParameterDef[string] {
	return &ParameterDef[string]{
		id:             "PreferredMATLABStartingDirectory",
		flagName:       "initial-working-folder",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_PreferredMATLABStartingDirectoryDescription,
		defaultValue:   "",
		recordToLog:    true,
	}
}

func BaseDir() *ParameterDef[string] {
	return &ParameterDef[string]{
		id:             "BaseDir",
		flagName:       "log-folder",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_BaseDirDescription,
		defaultValue:   "",
		recordToLog:    false,
	}
}

func LogLevel() *ParameterDef[string] {
	return &ParameterDef[string]{
		id:             "LogLevel",
		flagName:       "log-level",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_LogLevelDescription,
		defaultValue:   "info",
		recordToLog:    true,
	}
}

func InitializeMATLABOnStartup() *ParameterDef[bool] {
	return &ParameterDef[bool]{
		id:             "InitializeMATLABOnStartup",
		flagName:       "initialize-matlab-on-startup",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_InitializeMATLABOnStartupDescription,
		defaultValue:   false,
		recordToLog:    true,
	}
}

func MATLABDisplayMode() *ParameterDef[string] {
	return &ParameterDef[string]{
		id:             "MATLABDisplayMode",
		flagName:       "matlab-display-mode",
		hiddenFlag:     false,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_DisplayModeDescription,
		defaultValue:   "desktop",
		recordToLog:    true,
	}
}

func UseSingleMATLABSession() *ParameterDef[bool] {
	return &ParameterDef[bool]{
		id:             "UseSingleMATLABSession",
		flagName:       "use-single-matlab-session",
		hiddenFlag:     true,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_UseSingleMATLABSessionDescription,
		defaultValue:   true,
		recordToLog:    true,
	}
}

func WatchdogMode() *ParameterDef[bool] {
	return &ParameterDef[bool]{
		id:             "WatchdogMode",
		flagName:       "watchdog",
		hiddenFlag:     true,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_InternalUseDescription,
		defaultValue:   false,
	}
}

func ServerInstanceID() *ParameterDef[string] {
	return &ParameterDef[string]{
		id:             "ServerInstanceID",
		flagName:       "server-instance-id",
		hiddenFlag:     true,
		envVarName:     "",
		descriptionKey: messages.CLIMessages_InternalUseDescription,
		defaultValue:   "",
	}
}
