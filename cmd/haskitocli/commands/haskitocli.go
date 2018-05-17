package commands

import (
	"os"

	"github.com/spf13/cobra"
)

// haskitocli usage template
var usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletDisable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

  available with wallet enable:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletEnable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// HaskitocliCmd is haskitocli's root Command
var HaskitocliCmd = &cobra.Command{
	Use:   "haskitocli",
	Short: "Haskitocli is a command line client for haskito core (a.k.a haskitod)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.SetUsageTemplate(usageTemplate)
			cmd.Usage()
		}
	},
}

// Execute adds all child commands to the root command HaskitocliCmd and set flags appropriately.
func Execute() {
	AddCommands()
	AddTemplateFunc()

	if _, err := HaskitocliCmd.ExecuteC(); err != nil {
		os.Exit(-1)
	}
}

// AddCommands adds child commands to the root command HaskitocliCmd.
func AddCommands() {
	HaskitocliCmd.AddCommand(versionCmd)
}

// AddTemplateFunc adds usage template to the root command HaskitocliCmd
func AddTemplateFunc() {
	walletEnableCmd := []string{}

	cobra.AddTemplateFunc("WalletEnable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return true
			}
		}
		return false
	})

	cobra.AddTemplateFunc("WalletDisable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return false
			}
		}
		return true
	})
}
