package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	rest_server "github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/cmd/rest"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generated code example",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// Registering CMD
	{% if cookiecutter.add_rest_server != "none" %}
	rest_server.RestCmd.PersistentFlags().StringVar(&rest_server.ENV, "env", "dev", "--env=dev")
	rest_server.RestCmd.PersistentFlags().StringVar(&rest_server.Address, "port", "dev", "--env=dev")
	rootCmd.AddCommand(rest_server.RestCmd)
	{% endif %}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
