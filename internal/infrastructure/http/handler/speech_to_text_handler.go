package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	speech "cloud.google.com/go/speech/apiv2"
	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type SpeechToTextHandler struct {
	ctx context.Context
}

type GenerateTranscriptionRequest struct {
	AudioBase64 string `json:"audio_base64" binding:"required"`
}

func NewSpeechToTextHandler(ctx context.Context) *SpeechToTextHandler {
	return &SpeechToTextHandler{
		ctx: ctx,
	}
}

func (s *SpeechToTextHandler) GenerateTranscription(c *gin.Context) {
	var reqBody GenerateTranscriptionRequest
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	base64Audio := reqBody.AudioBase64
	if base64Audio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "audio_base64 query parameter is required"})
		return
	}
	sp, err := speech.NewClient(s.ctx, option.WithEndpoint("us-speech.googleapis.com:443"))
	if err != nil {
		log.Fatalln(fmt.Errorf("speech.NewClient error: %v", err))
	}
	defer sp.Close()
	stream, err := sp.StreamingRecognize(s.ctx)
	if err != nil {
		log.Fatal(err)
	}
	rawStr := strings.ReplaceAll(base64Audio, " ", "+")
	data, err := base64.StdEncoding.DecodeString(rawStr)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "failed to decode base64 audio"})
		return
	}

	err = stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					DecodingConfig: &speechpb.RecognitionConfig_AutoDecodingConfig{},
					Model:          "chirp_3",
					LanguageCodes:  []string{"es-ES"},
				},
			},
		},
		Recognizer: fmt.Sprintf("projects/%s/locations/us/recognizers/_", os.Getenv("GOOGLE_CLOUD_PROJECT")),
	})
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "failed to send streaming config"})
		return
	}
	go func() {
		chunkSize := 16384
		for i := 0; i < len(data); i += chunkSize {
			end := i + chunkSize
			if end > len(data) {
				end = len(data)
			}

			stream.Send(&speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_Audio{
					Audio: data[i:end],
				},
			})
		}
		stream.CloseSend()
	}()
	// 3. Recibir las respuestas
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error recibiendo: %v", err)
		}

		for _, result := range resp.Results {
			fmt.Printf("%s\n", result)
			if result.IsFinal {
				fmt.Printf("Transcripción Final: %v\n", result.Alternatives[0].Transcript)
				c.JSON(200, gin.H{"transcription": result.Alternatives[0].Transcript})
				return

			}
		}
	}
}
