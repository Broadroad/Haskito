package commands

import (
	"os"
	"path"

	cfg "github.com/haskito/config"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var initFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize blockchain",
	Run:   initFiles,
}

var config = cfg.DefaultConfig()

func init() {
	initFilesCmd.Flags().String("chain_id", config.ChainID, "Select [mainnet] or [testnet] or [solonet]")
	HaskitodCmd.AddCommand(initFilesCmd)
}

func initFiles(cmd *cobra.Command, args []string) {
	configFilePath := path.Join(config.RootDir, "Config.toml")
	if _, err := os.Stat(configFilePath); !os.IsNotExist(err) {
		log.WithField("config", configFilePath).Info("Already exists config file.")
		return
	}
}
