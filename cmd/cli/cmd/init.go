package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var exampleConfig = `[aes]
master_key="4ebca8b93b2fa70067ea92b33cdf9b4d011e3acb2ccabe768272157058e1cb9b"

[immu]
user="immudb"
password="immudb"
db="defaultdb"

[jwt]
access_token_duration="30m"
refresh_token_duration="720h"
secret_key="xbXRN2BAqdUvY4dz6srExeJldJwVFPkUSgHw5tp9cbFPgBHaWG9L9e0E27v7v92"
secure_cookie=false
access_cookie_name=".kimcha_access_token"
refresh_cookie_name=".kimcha_refresh_token"
`

var initCmd = &cobra.Command{
	Use:        "init",
	Aliases:    []string{"config"},
	SuggestFor: []string{"config"},
	Short:      "Creates .kimcha directory and example config (example.toml)",
	Long: color.New(color.Bold).Sprintln("Example config (example.toml) contents") +
		exampleConfig +
		"\n[jwt] section is " + color.New(color.Bold).Sprint("optional") + ", in case if you want to use front-end app",
	Args: nil,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		kimchaPath := filepath.Join(homeDir, ".kimcha")

		configPath := filepath.Join(kimchaPath, "example.toml")
		cliConfigPath := filepath.Join(kimchaPath, "cli")

		if _, err := os.Stat(kimchaPath); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(kimchaPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
			err = os.WriteFile(configPath, []byte(exampleConfig), 0644)
			if err != nil {
				panic(err)
			}
		}

		if _, err := os.Stat(cliConfigPath); errors.Is(err, os.ErrNotExist) {
			err = os.WriteFile(cliConfigPath, []byte{}, 0644)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(fmt.Sprintf("Example config created at path: %s\n\ncopy it to config.toml and replace with your values", configPath))
	},
}

func init() {

	rootCmd.AddCommand(initCmd)
}
