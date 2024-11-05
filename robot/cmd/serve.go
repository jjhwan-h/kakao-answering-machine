package cmd

import (
	"chatbot/internal"
	"time"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts a weather-bot",
	Long:  "starts a weather-bot",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			Start()
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func Start() {
	messageBox := &internal.DIR{X1: 490, Y1: 900, X2: 570, Y2: 955}
	speechBox := &internal.SpeechBox{X: 0, Y: 760, Width: 560, Height: 72}
	//testBox := []int{500, 145}
	redPixels := internal.CollectRedPixels(messageBox)
	if len(redPixels) > 0 {
		internal.Click(redPixels[0].X, redPixels[0].Y)
		//internal.Click(testBox[0], testBox[1])
		time.Sleep(2 * time.Second)
		bit := internal.CaptureScreen(speechBox)
		defer internal.FreeBitmap(bit)

		region := internal.ReqOCR(bit)

		if len(region) > 0 {
			res := internal.GetWeather(region)
			//res = internal.Joke(res, region)
			internal.Click(40, 850)
			internal.EnterInput(res)
		}
		internal.Input("esc")
	}
}
