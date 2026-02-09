// Copyright 2026 The MathWorks, Inc.

package sdk_test

import (
	"context"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/matlab/matlab-mcp-core-server/internal/adaptors/time/retry"
	"github.com/matlab/matlab-mcp-core-server/tests/functional/sdk/testbinaries"
	"github.com/matlab/matlab-mcp-core-server/tests/testutils/mcpclient"
	"github.com/stretchr/testify/suite"
)

// ServerWithCustomParametersTestSuite tests the custom parameters functionality of the SDK.
type ServerWithCustomParametersTestSuite struct {
	suite.Suite

	serverDetails testbinaries.ServerDetails
}

// SetupSuite runs once before all tests in a suite
func (s *ServerWithCustomParametersTestSuite) SetupSuite() {
	s.serverDetails = testbinaries.BuildServerWithCustomParameters(s.T())
}

func TestServerWithCustomParametersTestSuite(t *testing.T) {
	suite.Run(t, new(ServerWithCustomParametersTestSuite))
}

func (s *ServerWithCustomParametersTestSuite) TestSDK_CustomParameter_HappyPath() {
	// Connect to a session
	expectedID := "custom-param-id"
	expectedValue := "someValue"

	expectedRecordedID := "custom-recorded-param-id"
	expectedRecordedValue := "someOtherValue"

	client := mcpclient.NewClient(s.T().Context(), s.serverDetails.BinaryLocation(), nil,
		"--log-level=debug",
		"--custom-param="+expectedValue,
		"--custom-recorded-param="+expectedRecordedValue,
	)

	session, err := client.CreateSession(s.T().Context())
	s.Require().NoError(err, "should create MCP session")
	defer func() {
		s.Require().NoError(session.Close(), "closing session should not error")
	}()

	// Setup regexp to look into the logs
	anyCharacterButNewLines := `[^\n]+`
	configStateLogMessage := "Configuration state"
	configStateRegExp, err := regexp.Compile(configStateLogMessage + anyCharacterButNewLines + expectedRecordedID + anyCharacterButNewLines + expectedRecordedValue)
	s.Require().NoError(err)

	dependenciesProviderLogMessage := "Config value from dependency provider"
	dependenciesProviderRegExp, err := regexp.Compile(dependenciesProviderLogMessage + anyCharacterButNewLines + expectedID + anyCharacterButNewLines + expectedValue)
	s.Require().NoError(err)

	toolsProviderLogMessage := "Config value from tools provider"
	toolsProviderRegExp, err := regexp.Compile(toolsProviderLogMessage + anyCharacterButNewLines + expectedID + anyCharacterButNewLines + expectedValue)
	s.Require().NoError(err)

	// Check that the log features the custom parameter values
	ctx, cancel := context.WithTimeout(s.T().Context(), 2*time.Second) // Timeout for the logs to appear in the stream
	defer cancel()

	_, err = retry.Retry(ctx, func() (struct{}, bool, error) {
		logContent := session.Stderr()

		foundAllLogEntries := configStateRegExp.MatchString(logContent) &&
			dependenciesProviderRegExp.MatchString(logContent) &&
			toolsProviderRegExp.MatchString(logContent)

		return struct{}{}, foundAllLogEntries, nil
	}, retry.NewLinearRetryStrategy(200*time.Millisecond))

	s.Require().NoError(err, "Failed to find the log entries in :\n\n%s", session.Stderr())
}

func (s *ServerWithCustomParametersTestSuite) TestSDK_CustomParameter_Recorded_ByEnvVar() {
	// Connect to a session
	expectedRecordedID := "custom-recorded-param-id"
	expectedRecordedValue := "someValue"

	env := append(os.Environ(), "CUSTOM_RECORDED_PARAM="+expectedRecordedValue)
	client := mcpclient.NewClient(s.T().Context(), s.serverDetails.BinaryLocation(), env)

	session, err := client.CreateSession(s.T().Context())
	s.Require().NoError(err, "should create MCP session")
	defer func() {
		s.Require().NoError(session.Close(), "closing session should not error")
	}()

	// Check that the log features the custom parameter values
	anyCharacterButNewLines := `[^\n]+`
	configStateLogMessage := "Configuration state"
	configStateRegExp, err := regexp.Compile(configStateLogMessage + anyCharacterButNewLines + expectedRecordedID + anyCharacterButNewLines + expectedRecordedValue)
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(s.T().Context(), 2*time.Second)
	defer cancel()

	_, err = retry.Retry(ctx, func() (struct{}, bool, error) {
		return struct{}{}, configStateRegExp.MatchString(session.Stderr()), nil
	}, retry.NewLinearRetryStrategy(200*time.Millisecond))

	s.Require().NoError(err, "custom recorded param should be logged with expected value:\n\n%s", session.Stderr())
}
