package cmd

import (
	"errors"
	"fmt"
	"github.com/ngyewch/go-clibase"
	"github.com/ngyewch/go-ntfsvc/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	goVersion "go.hein.dev/go-version"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:  fmt.Sprintf("%s (topic) (message)", AppName),
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := doRunE(cmd, args)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				os.Exit(1)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("url", "", "notification service URL (NTFSVC_URL)")
	rootCmd.PersistentFlags().String("api-key", "", "notification service API key (NTFSVC_APIKEY)")

	viper.SetEnvPrefix("NTFSVC")
	viper.AutomaticEnv()
	_ = viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	_ = viper.BindPFlag("apiKey", rootCmd.PersistentFlags().Lookup("api-key"))

	clibase.AddVersionCmd(rootCmd, func() *goVersion.Info {
		return VersionInfo
	})
}

func initConfig() {
	// do nothing
}

func doRunE(cmd *cobra.Command, args []string) error {
	url := viper.GetString("url")
	if url == "" {
		return errors.New("url not specified")
	}

	apiKey := viper.GetString("apiKey")
	if apiKey == "" {
		return errors.New("apiKey not specified")
	}

	topic := args[0]
	if topic == "" {
		return errors.New("topic not specified")
	}

	message := args[1]
	if message == "" {
		return errors.New("message not specified")
	}

	notificationService := client.NewNotificationService(url, apiKey)
	err := notificationService.SendNotification(topic, message)
	if err != nil {
		return err
	}

	return nil
}
