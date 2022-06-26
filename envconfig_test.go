package envconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FromEnvs struct {
	Host string
	Port int
	Live bool
	Dead bool
	Math float64
	Text string `default:"example"`
	TheZ string `underscore:"false"`
	Just string `env:"JUST"`
}

type EnvTestSuite struct {
	suite.Suite
	StructWithVariables FromEnvs
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}

func (suite *EnvTestSuite) SetupSuite() {
	os.Setenv("MYAPP_HOST", "localhost")
	os.Setenv("MYAPP_PORT", "9999")
	os.Setenv("MYAPP_LIVE", "true")
	os.Setenv("MYAPP_DEAD", "0")
	os.Setenv("MYAPP_THEZ", "Z")
	os.Setenv("MYAPP_MATH", "3.14")
	os.Setenv("JUST", "alone")

	Process("MYAPP", &suite.StructWithVariables)
}

func (suite *EnvTestSuite) TestAssertingAllVariablesWereLoadedWithTheRightCasting() {
	assert.Equal(suite.T(), "localhost", suite.StructWithVariables.Host)
	assert.Equal(suite.T(), 9999, suite.StructWithVariables.Port)
	assert.Equal(suite.T(), 3.14, suite.StructWithVariables.Math)
	assert.Equal(suite.T(), true, suite.StructWithVariables.Live)
	assert.Equal(suite.T(), false, suite.StructWithVariables.Dead)
}

func (suite *EnvTestSuite) TestShouldSetADefaultValueWhenUsingAnAnnotation() {
	assert.Equal(suite.T(), "example", suite.StructWithVariables.Text)
}

func (suite *EnvTestSuite) TestShouldNotIncludeUnderscoreDuringTheLookupForAnEnv() {
	assert.Equal(suite.T(), "Z", suite.StructWithVariables.TheZ)
}

func (suite *EnvTestSuite) TestShouldIgnoreThePrefixWhenDefinedWithTheKeywordEnv() {
	assert.Equal(suite.T(), "alone", suite.StructWithVariables.Just)
}

func (suite *EnvTestSuite) TestShouldRaisePanicWhenANotCoveredTypedIsUsedAsEnv() {
	os.Setenv("MYAPP_ELSE", "{}")

	type ShouldPanic struct {
		Else struct{}
	}

	var sp ShouldPanic

	assert.Panics(suite.T(), func() {
		Process("MYAPP", &sp)
	})

}
