package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "weather-bot",
	Short: "날씨 [지역]을 입력하면 날씨를 알려주는 챗봇입니다.",
	Long:  "날씨 [지역]을 입력하면 날씨를 알려주는 챗봇입니다. 이 봇은 사용자가 지역을 입력하면 그 지역의 현재 날씨를 알려줍니다.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}
