package main

import (
	"context"
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/genai"
)

//go:embed system-prompt.txt
var systemPrompt string

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path_to_png_image>", os.Args[0])
	}
	if err := run(context.Background(), os.Args[1]); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}

func run(ctx context.Context, imagePath string) error {
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return fmt.Errorf("failed to read image file %s: %w", imagePath, err)
	}
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		return fmt.Errorf("GOOGLE_CLOUD_PROJECT environment variable not set")
	}
	location := os.Getenv("GOOGLE_CLOUD_LOCATION")
	if location == "" {
		return fmt.Errorf("GOOGLE_CLOUD_LOCATION environment variable not set")
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  projectID,
		Location: location,
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return fmt.Errorf("failed to create Vertex AI client: %w", err)
	}
	slog.Info("generating CSV from KBA image", "imagePath", imagePath)
	resp, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", []*genai.Content{
		{
			Parts: []*genai.Part{
				{InlineData: &genai.Blob{Data: imageData, MIMEType: "image/png"}},
			},
			Role: genai.RoleUser,
		},
	}, &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: systemPrompt},
			},
			Role: genai.RoleUser,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to generate content from Vertex AI API: %w", err)
	}
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
		return fmt.Errorf("no content received from Vertex AI API")
	}
	var response strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		response.WriteString(part.Text)
	}
	csvPath := strings.TrimSuffix(imagePath, filepath.Ext(imagePath)) + ".csv"
	f, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	csvReader := csv.NewReader(strings.NewReader(response.String()))
	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		csvWriter.Write(record)
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("failed to write CSV: %w", err)
	}
	slog.Info("generated CSV successfully", "csvPath", csvPath)
	return nil
}
