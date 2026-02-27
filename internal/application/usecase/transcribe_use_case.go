package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	speech "cloud.google.com/go/speech/apiv2"
	"cloud.google.com/go/speech/apiv2/speechpb"
	"google.golang.org/api/option"
)

type TranscribeUseCase struct {
	ctx context.Context
}

func NewTranscribeUseCase(ctx context.Context) *TranscribeUseCase {
	return &TranscribeUseCase{
		ctx: ctx,
	}
}

func (tuc *TranscribeUseCase) SpeechToText(base64Audio string) (*speechpb.RecognizeResponse, error) {
	sp, err := speech.NewClient(tuc.ctx, option.WithEndpoint("us-speech.googleapis.com:443"))
	if err != nil {
		return nil, err
	}
	defer sp.Close()
	rawStr := strings.ReplaceAll(base64Audio, " ", "+")
	data, err := base64.StdEncoding.DecodeString(rawStr)
	if err != nil {
		return nil, err
	}
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			DecodingConfig: &speechpb.RecognitionConfig_AutoDecodingConfig{},
			Model:          os.Getenv("CHIRP_MODEL"),
			LanguageCodes:  []string{"pt-BR", "es-ES"},
		},
		AudioSource: &speechpb.RecognizeRequest_Content{
			Content: data,
		},
		Recognizer: fmt.Sprintf("projects/%s/locations/us/recognizers/_", os.Getenv("GOOGLE_CLOUD_PROJECT")),
	}
	resp, err := sp.Recognize(tuc.ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
