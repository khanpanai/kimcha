package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"kimcha/types"
	"os"
	"path/filepath"
)

type CliConfig struct {
	MasterKeyHash string `toml:"master_key_hash"`
	MasterKeySalt string `toml:"master_key_salt"`
}

func ReadConfigFromHomeDirToViper() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	kimchaPath := filepath.Join(homeDir, ".kimcha")

	if _, err := os.Stat(kimchaPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("app is not initialized"))
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(kimchaPath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(types.KhakiStyle.Render(fmt.Sprintf("fatal error config file at path %s. Ensure that you have run kimcha init and created config.toml file", kimchaPath)))
		os.Exit(1)
	}
}

func ReadCliConfigFromHomeDirToViper() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	kimchaPath := filepath.Join(homeDir, ".kimcha")

	if _, err := os.Stat(kimchaPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("app is not initialized"))
	}

	configPath := filepath.Join(kimchaPath, "cli")

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("cli config does not exist"))
	}

	viper.SetConfigName("cli")
	viper.SetConfigType("toml")
	viper.AddConfigPath(kimchaPath)
}

func ReadCliConfigToStruct() CliConfig {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	kimchaPath := filepath.Join(homeDir, ".kimcha")

	if _, err := os.Stat(kimchaPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("app is not initialized"))
	}

	configPath := filepath.Join(kimchaPath, "cli")

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("cli config does not exist"))
	}

	var cliConfig CliConfig

	_, err = toml.DecodeFile(configPath, &cliConfig)
	if err != nil {
		return CliConfig{}
	}

	return cliConfig
}

func DumpCliConfig(cliConfig *CliConfig) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	kimchaPath := filepath.Join(homeDir, ".kimcha")

	if _, err := os.Stat(kimchaPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("app is not initialized"))
	}

	configPath := filepath.Join(kimchaPath, "cli")

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		panic(errors.New("cli config does not exist"))
	}

	file, err := os.OpenFile(configPath, os.O_CREATE, 0644)
	if err != nil {
		panic(errors.New("can't open cli config file to dump"))
	}

	enc := toml.NewEncoder(file)

	err = enc.Encode(cliConfig)
	if err != nil {
		return err
	}

	return nil
}
