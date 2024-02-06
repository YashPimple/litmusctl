package config_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/litmuschaos/litmusctl/cmd/config"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountsCmd(t *testing.T) {
	tests := []struct {
		name         string
		configPath   string
		expectedOut  string
		expectedErr  string
		shouldCreate bool
	}{
		{
			name:        "NoAccounts",
			expectedOut: "CURRENT\tENDPOINT\tUSERNAME\tEXPIRESIN\n",
			expectedErr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.shouldCreate {
				err := os.WriteFile("litmusconfig.yaml", []byte(strings.TrimSpace(test.configPath)), 0644)
				assert.NoError(t, err)
				defer os.Remove("litmusconfig.yaml")
			}
			cmd := config.NewCmdGetAccounts()
			cmd.SetArgs([]string{})
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			err := cmd.Execute()
			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedOut, buf.String())
		})
	}
}
