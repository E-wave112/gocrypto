/*
Copyright © 2023 NAME HERE <iyayiemmanuel1@gmail.com>
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/E-wave112/gocrypto/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocrypto",
	Short: "a crypto/fiat exchange cli built with go",
	Long: `a crypto/fiat exchange cli built with go, get real time USD
	exchange rates of common crypto coins such as Bitcoin (BTC), Ether(ETH)
	Dogecoin(DOGE), Solana(SOL), Shiba(SHIB) and Tether(USDT)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists the supported cryptocurrencies to get real-time exchange rates from",
	Run: func(cmd *cobra.Command, args []string) {
		supportedCurrencies := pkg.ListSupportedCryptoCurrencies()
		fmt.Println("Supported cryptocurrencies:")
		for _, currency := range supportedCurrencies {
			fmt.Println(currency)
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves the exchange rate between the supported cryptocurrencies and any specified fiat currency provided as an argument.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currency := strings.Join(args, " ")
		currency = strings.ToUpper(currency)
		response, _ := pkg.RetrieveRates(currency)
		fmt.Printf("Realtime exchange rates for %q :\n", currency)

		for crypto, rate := range response {
			crypto = fmt.Sprintf("1 %s", crypto)
			rate = fmt.Sprintf("%s %s", rate, currency)
			fmt.Printf("%s: %s\n", crypto, rate)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// add extra commands
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocrypto.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gocrypto" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gocrypto")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	Execute()
}
