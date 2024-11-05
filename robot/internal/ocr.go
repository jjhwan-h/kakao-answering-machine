package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/viper"
)

func ReqOCR(bitmap *robotgo.CBitmap) []string {
	img := robotgo.ToImage(*bitmap)
	if img == nil {
		log.Printf("Failed to convert CBitmap to image.Image")
		return nil
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("img", "screenshot.png") // The filename is just a placeholder
	if err != nil {
		log.Printf("Failed to create form file: %v", err)
		return nil
	}

	if err := png.Encode(fw, img); err != nil {
		log.Printf("Failed to encode image: %v", err)
		return nil
	}
	w.Close()
	uri := viper.GetString("OCRAPI")
	req, err := http.NewRequest("POST", uri, &b)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return nil
	}

	var res []string
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Printf("req : %s\n", res)
	defer resp.Body.Close()
	return res
}

// // Save the image to a file
// filePath := "screenshot.png"
// outFile, err := os.Create(filePath)
// if err != nil {
// 	log.Printf("Failed to create file: %v", err)
// 	return nil
// }

// if err := png.Encode(outFile, img); err != nil {
// 	log.Printf("Failed to encode image: %v", err)
// 	return nil
// }

// outFile.Close()
