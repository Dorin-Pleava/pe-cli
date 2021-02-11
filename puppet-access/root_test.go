package cmd

import (
	"path/filepath"
	"testing"

	"github.com/puppetlabs/pe-cli/testdata"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGlobalConfigFileAbsent(t *testing.T) {
	assert := assert.New(t)
	err := readGlobalConfigFile()
	assert.NoError(err)
}

func TestDefaultConfigFileAbsent(t *testing.T) {
	assert := assert.New(t)
	err := readConfigFile(getDefaultConfigFile())
	assert.NoError(err)
}

func TestCLIConfigFileAbsent(t *testing.T) {
	assert := assert.New(t)
	err := readConfigFile("/path/to/absent/config")
	assert.Error(err)
}

func TestCanReadAndAliasConfigParameters(t *testing.T) {
	assert := assert.New(t)
	initConfig(filepath.Join(testdata.FixturePath(), "puppet-access.conf"))
	assert.Equal("https://<CONSOLE HOSTNAME>:4433/rbac-api", viper.GetString("service-url"))
	assert.Equal("/path/to/cacert", viper.GetString("certificate-file"))
	assert.Equal("/path/to/token", viper.GetString("token-file"))
}
