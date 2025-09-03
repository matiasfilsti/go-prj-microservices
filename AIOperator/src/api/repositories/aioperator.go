package repositories

import (
	"aioperator/src/api/config"
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func NewAiOperator() (llms.Model, *embeddings.EmbedderImpl) {
	llm, err := ollama.New(
		ollama.WithModel("llama3.2"),
		ollama.WithFormat("json"),
		ollama.WithServerURL(config.LlmUrl),
	)
	if err != nil {
		log.Fatal(err)
	}
	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}
	return llms.Model(llm), e
}

type AiOperatorRepository struct {
	llm   llms.Model
	store *redisvector.Store
}

func NewAiOperatorRepository(llm llms.Model, store *redisvector.Store) *AiOperatorRepository {
	return &AiOperatorRepository{
		llm:   llm,
		store: store,
	}
}

type response struct {
	Id         string `json:"id"`
	Version    string `json:"version"`
	Type       string `json:"type"`
	Language   string `json:"language"`
	SourceType string `json:"sourceType"`
	Answer     string `json:"answer"`
}

// func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document, question string) (string, error) {
// 	// Debug: Print the document structure
// 	fmt.Printf("Document content: %+v\n", docs[0].PageContent)
// 	fmt.Printf("Document metadata: %+v\n", docs[0].Metadata)
// 	fmt.Printf("Question: %s\n", question)

// 	summaryTemplate := `Eres un asistente que envia a otra funcion los datos de la ciudad.

// 	Nombre: {{.name}}
// 	Población: {{.population}} millones
// 	País: {{.country}}
// 	Área o superficie: {{.area}} km²
// 	Fundada: {{.founded}}
// 	Zona horaria: {{.timezone}}
// 	GDP: {{.gdp}}
// 	Elevación: {{.elevation}} metros
// 	Idioma: {{.language}}
// 	Moneda: {{.currency}}
// 	Puntos turísticos: {{.landmarks}}
// 	Mejor período para visitar: {{.best_travel_period}}
// 	`

// 	summaryChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(summaryTemplate,
// 		[]string{"name", "population", "country", "language", "currency", "area", "founded", "timezone", "gdp", "elevation", "landmarks", "best_travel_period"}))
// 	summaryChain.OutputKey = "resumen"

// 	answerTemplate := `Eres un asistente que recibe informacion de una funcion y responde preguntas sobre ciudades.
// 		Basándote en la siguiente información:

// 		{{.resumen}}

// 		Instrucciones MUY IMPORTANTES:
// 		1. Responde ÚNICAMENTE con este formato exacto:
// 		"Hola, me llamo MMlA04, tu asistente virtual, aquí está tu respuesta: [respuesta]"
// 		2. No incluyas nada más en tu respuesta
// 		3. La pregunta es: ` + question + `
// 		4. No repitas la pregunta
// 		5. Responde en español
// 		6. Responde en texto plano
// 		`

// 	answerChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(answerTemplate,
// 		[]string{"resumen"}))
// 	answerChain.OutputKey = "respuesta"

// 	sequentialChain, err := chains.NewSequentialChain(
// 		[]chains.Chain{summaryChain, answerChain},
// 		[]string{"name", "population", "country", "language", "currency", "area", "founded", "timezone", "gdp", "elevation", "landmarks", "best_travel_period"},
// 		[]string{"respuesta"},
// 	)
// 	if err != nil {
// 		return "", fmt.Errorf("error creating sequential chain: %v", err)
// 	}

// 	// Prepare input - ensure all required fields exist
// 	input := map[string]any{
// 		"name":               docs[0].PageContent, // Use PageContent as the name
// 		"country":            docs[0].Metadata["country"],
// 		"population":         docs[0].Metadata["population"],
// 		"language":           docs[0].Metadata["language"],
// 		"currency":           docs[0].Metadata["currency"],
// 		"area":               docs[0].Metadata["area"],
// 		"founded":            docs[0].Metadata["founded"],
// 		"timezone":           docs[0].Metadata["timezone"],
// 		"gdp":                docs[0].Metadata["gdp"],
// 		"elevation":          docs[0].Metadata["elevation"],
// 		"landmarks":          docs[0].Metadata["landmarks"],
// 		"best_travel_period": docs[0].Metadata["best_travel_period"],
// 	}

// 	fmt.Printf("Final input to chain: %+v\n", input)

// 	// Execute the chain
// 	result, err := chains.Call(ctx, sequentialChain, input)
// 	if err != nil {
// 		return "", fmt.Errorf("error executing chain: %v", err)
// 	}

// 	fmt.Printf("Raw result from chain: %+v\n", result)

// 	// var response response
// 	// if err := json.Unmarshal([]byte(result["respuesta"].(string)), &response); err != nil {
// 	// 	return "", fmt.Errorf("error unmarshalling response: %v", err)
// 	// }
// 	// Extract and return the final answer

// 	// return response.Answer, nil
// 	resp, ok := result["respuesta"].(string)
// 	if !ok {
// 		return "", fmt.Errorf("invalid response format")
// 	}
// 	return resp, nil
// }

func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document, question string) (string, error) {
	// Debug: Print the document structure
	fmt.Printf("Document content: %+v\n", docs[0].PageContent)
	fmt.Printf("Document metadata: %+v\n", docs[0].Metadata)
	fmt.Printf("Question: %s\n", question)

	summaryTemplate := `Eres un asistente que responde preguntas sobre ciudades.
	Basándote en la siguiente información:
	Nombre: {{.name}}
	Población: {{.population}} millones
	País: {{.country}}
	Área o superficie: {{.area}} km²
	Fundada: {{.founded}}
	Zona horaria: {{.timezone}}
	GDP: {{.gdp}}
	Elevación: {{.elevation}} metros
	Idioma: {{.language}}
	Moneda: {{.currency}}
	Puntos turísticos: {{.landmarks}}
	Mejor período para visitar: {{.best_travel_period}}
    
	responde la siguiente pregunta: ` + question + `

	Instrucciones:
	1. Responde ÚNICAMENTE con este formato exacto:
	"Hola, me llamo MMlA04, tu asistente virtual, aquí está tu respuesta: [respuesta]"
	2. No repitas la pregunta
	3. Responde en español
	4. Responde en texto plano
	`

	summaryChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(summaryTemplate,
		[]string{"name", "population", "country", "language", "currency", "area", "founded", "timezone", "gdp", "elevation", "landmarks", "best_travel_period"}))
	summaryChain.OutputKey = "respuesta"

	sequentialChain, err := chains.NewSequentialChain(
		[]chains.Chain{summaryChain},
		[]string{"name", "population", "country", "language", "currency", "area", "founded", "timezone", "gdp", "elevation", "landmarks", "best_travel_period"},
		[]string{"respuesta"},
	)
	if err != nil {
		return "", fmt.Errorf("error creating sequential chain: %v", err)
	}

	// Prepare input - ensure all required fields exist
	input := map[string]any{
		"name":               docs[0].PageContent, // Use PageContent as the name
		"country":            docs[0].Metadata["country"],
		"population":         docs[0].Metadata["population"],
		"language":           docs[0].Metadata["language"],
		"currency":           docs[0].Metadata["currency"],
		"area":               docs[0].Metadata["area"],
		"founded":            docs[0].Metadata["founded"],
		"timezone":           docs[0].Metadata["timezone"],
		"gdp":                docs[0].Metadata["gdp"],
		"elevation":          docs[0].Metadata["elevation"],
		"landmarks":          docs[0].Metadata["landmarks"],
		"best_travel_period": docs[0].Metadata["best_travel_period"],
	}

	fmt.Printf("Final input to chain: %+v\n", input)

	// Execute the chain
	result, err := chains.Call(ctx, sequentialChain, input)
	if err != nil {
		return "", fmt.Errorf("error executing chain: %v", err)
	}

	fmt.Printf("Raw result from chain: %+v\n", result)

	// var response response
	// if err := json.Unmarshal([]byte(result["respuesta"].(string)), &response); err != nil {
	// 	return "", fmt.Errorf("error unmarshalling response: %v", err)
	// }
	// Extract and return the final answer

	// return response.Answer, nil
	resp, ok := result["respuesta"].(string)
	if !ok {
		return "", fmt.Errorf("invalid response format")
	}
	return resp, nil
}
