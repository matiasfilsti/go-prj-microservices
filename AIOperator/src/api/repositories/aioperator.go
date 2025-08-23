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
		ollama.WithModel("llama3.2:1b"),
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

// func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document) (string, error) {
// 	promptTemplate := prompts.NewPromptTemplate(`teniendo la siguiente informacion de una ciudad, responde la siguiente pregunta: cual es su población?
// 		- nombre: {{.name}}
// 		- poblacion: {{.population}}
// 		- area: {{.area}}
// 		- pais: {{.country}}
// 		- fundacion: {{.founded}}
// 		- timezone: {{.timezone}}
// 		- gdp: {{.gdp}}
// 		- elevation: {{.elevation}}
// 		- language: {{.language}}
// 		- currency: {{.currency}}
// 		- landmarks: {{.landmarks}}
// 		- best_travel_period: {{.best_travel_period}},
// 		`, []string{"name", "population", "area", "country", "founded", "timezone", "gdp", "elevation", "language", "currency", "landmarks", "best_travel_period", "question"})

// 	//llmChain := chains.NewLLMChain(r.llm, prompt) {{.question}}

// 	input := map[string]any{
// 		"name":               docs[0].PageContent,
// 		"population":         docs[0].Metadata["population"],
// 		"area":               docs[0].Metadata["area"],
// 		"country":            docs[0].Metadata["country"],
// 		"founded":            docs[0].Metadata["founded"],
// 		"timezone":           docs[0].Metadata["timezone"],
// 		"gdp":                docs[0].Metadata["gdp"],
// 		"elevation":          docs[0].Metadata["elevation"],
// 		"language":           docs[0].Metadata["language"],
// 		"currency":           docs[0].Metadata["currency"],
// 		"landmarks":          docs[0].Metadata["landmarks"],
// 		"best_travel_period": docs[0].Metadata["best_travel_period"],
// 		"question":           docs[0].PageContent,
// 	}
// 	fmt.Println(input["population"])
// 	llmChain := chains.NewLLMChain(r.llm, promptTemplate)
// 	llmChain.OutputKey = "respuesta"
// 	// prompt, err := promptTemplate.Format(input)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	simpleSeqChain, err := chains.NewSequentialChain([]chains.Chain{llmChain}, []string{"name", "population", "area", "country", "founded", "timezone", "gdp", "elevation", "language", "currency", "landmarks", "best_travel_period", "question"}, []string{"respuesta"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	result, err := chains.Call(ctx, simpleSeqChain, input)
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println("info response result", result)
// 	fmt.Println("info response err", err)
// 	return "hola", nil

// }

// func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document) (string, error) {
// 	// First chain: Process the city information
// 	summaryTemplate := `Eres un asistente que resume información sobre ciudades. Basándote en los siguientes datos,
// crea un resumen conciso sobre la ciudad:

// Nombre: {{.name}}
// Población: {{.population}} millones
// País: {{.country}}
// Idioma: {{.language}}
// Moneda: {{.currency}}

// Resumen conciso:`

// 	summaryChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(summaryTemplate,
// 		[]string{"name", "population", "country", "language", "currency"}))
// 	summaryChain.OutputKey = "resumen"

// 	// Second chain: Answer the specific question
// 	answerTemplate := `Eres un asistente que responde preguntas sobre ciudades.
// Basándote en la siguiente información:

// {{.resumen}}

// Responde de manera concisa a la pregunta: ¿Cuál es la población de {{.name}}?`

// 	answerChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(answerTemplate,
// 		[]string{"resumen", "name"}))
// 	answerChain.OutputKey = "respuesta"

// 	// Create sequential chain
// 	sequentialChain, err := chains.NewSequentialChain(
// 		[]chains.Chain{summaryChain, answerChain},
// 		[]string{"name", "population", "country", "language", "currency"},
// 		[]string{"respuesta"},
// 	)
// 	if err != nil {
// 		return "", fmt.Errorf("error creating sequential chain: %v", err)
// 	}

// 	// Prepare input
// 	input := map[string]any{
// 		"name":       docs[0].PageContent,
// 		"population": docs[0].Metadata["population"],
// 		"country":    docs[0].Metadata["country"],
// 		"language":   docs[0].Metadata["language"],
// 		"currency":   docs[0].Metadata["currency"],
// 	}

// 	// Execute the chain
// 	result, err := chains.Call(ctx, sequentialChain, input)
// 	if err != nil {
// 		return "", fmt.Errorf("error executing chain: %v", err)
// 	}

// 	// Extract and return the final answer
// 	if respuesta, ok := result["respuesta"].(string); ok {
// 		return respuesta, nil
// 	}

// 	return "No se pudo generar una respuesta.", nil
// }

// func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document) (string, error) {
// 	// Debug: Print the document structure
// 	fmt.Printf("Document content: %+v\n", docs[0].PageContent)
// 	fmt.Printf("Document metadata: %+v\n", docs[0].Metadata)

// 	// First chain: Process the city information
// 	summaryTemplate := `Eres un asistente que resume información sobre ciudades. Basándote en los siguientes datos,
// crea un resumen conciso sobre la ciudad:

// Nombre: {{.name}}
// Población: {{.population}} millones
// País: {{.country}}
// Idioma: {{.language}}
// Moneda: {{.currency}}

// Resumen conciso:`

// 	summaryChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(summaryTemplate,
// 		[]string{"name", "population", "country", "language", "currency"}))
// 	summaryChain.OutputKey = "resumen"

// 	// Second chain: Answer the specific question
// 	answerTemplate := `Eres un asistente que responde preguntas sobre ciudades.
// Basándote en la siguiente información:

// {{.resumen}}

// Responde de manera concisa a la pregunta: ¿Cuál es la población de {{.name}}?`

// 	answerChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(answerTemplate,
// 		[]string{"resumen", "name"}))
// 	answerChain.OutputKey = "respuesta"

// 	// Create sequential chain
// 	sequentialChain, err := chains.NewSequentialChain(
// 		[]chains.Chain{summaryChain, answerChain},
// 		[]string{"name", "population", "country", "language", "currency"},
// 		[]string{"respuesta"},
// 	)
// 	if err != nil {
// 		return "", fmt.Errorf("error creating sequential chain: %v", err)
// 	}

// 	// Prepare input - ensure all required fields exist
// 	input := map[string]any{
// 		"name":     docs[0].PageContent, // Use PageContent as the name
// 		"country":  "",
// 		"language": "",
// 		"currency": "",
// 	}

// 	// Copy metadata values if they exist
// 	if meta := docs[0].Metadata; meta != nil {
// 		if pop, ok := meta["population"]; ok {
// 			input["population"] = pop
// 		}
// 		if country, ok := meta["country"]; ok {
// 			input["country"] = country
// 		}
// 		if lang, ok := meta["language"]; ok {
// 			input["language"] = lang
// 		}
// 		if curr, ok := meta["currency"]; ok {
// 			input["currency"] = curr
// 		}
// 	}

// 	// Debug: Print the final input
// 	fmt.Printf("Final input to chain: %+v\n", input)

// 	// Execute the chain
// 	result, err := chains.Call(ctx, sequentialChain, input)
// 	if err != nil {
// 		return "", fmt.Errorf("error executing chain: %v", err)
// 	}

// 	// Debug: Print the raw result
// 	fmt.Printf("Raw result from chain: %+v\n", result)

// 	// Extract and return the final answer
// 	if respuesta, ok := result["respuesta"].(string); ok {
// 		return respuesta, nil
// 	}

// 	return "No se pudo generar una respuesta.", nil
// }

func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document) (string, error) {
	// Debug: Print the document structure
	fmt.Printf("Document content: %+v\n", docs[0].PageContent)
	fmt.Printf("Document metadata: %+v\n", docs[0].Metadata)

	// First chain: Process the city information
	summaryTemplate := `Eres un asistente que resume información sobre ciudades. Basándote en los siguientes datos, 
	crea un resumen conciso sobre la ciudad:

	Nombre: {{.name}}
	Población: {{.population}} millones
	País: {{.country}}
	Idioma: {{.language}}
	Moneda: {{.currency}}
	`

	summaryChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(summaryTemplate,
		[]string{"name", "population", "country", "language", "currency"}))
	summaryChain.OutputKey = "resumen"

	// Second chain: Answer the specific question
	// answerTemplate := `Eres un asistente que responde preguntas sobre ciudades.
	// 	Basándote en la siguiente información:

	// 	{{.resumen}}

	// 	Deci: Hola, me llamo MMlA04
	// 	luego:
	// 	Menciona el nombre de la ciudad y deci su capital
	// 	`

	// Second chain: Answer the specific question
	answerTemplate := `Eres un asistente que responde preguntas sobre ciudades. 
Basándote en la siguiente información:

{{.resumen}}

Instrucciones IMPORTANTES:
1. Saluda diciendo: "Hola, me llamo MMlA04"
2. Luego menciona el nombre de la ciudad y su capital
3. Usa un tono amigable y natural
4. NO devuelvas JSON, solo texto plano
5. Responde en español
6. Responde la cantidad de habitantes de la ciudad

Respuesta:`

	answerChain := chains.NewLLMChain(r.llm, prompts.NewPromptTemplate(answerTemplate,
		[]string{"resumen"}))
	answerChain.OutputKey = "respuesta"

	// Create sequential chain
	sequentialChain, err := chains.NewSequentialChain(
		[]chains.Chain{summaryChain, answerChain},
		[]string{"name", "population", "country", "language", "currency"},
		[]string{"respuesta"},
	)
	if err != nil {
		return "", fmt.Errorf("error creating sequential chain: %v", err)
	}

	// Prepare input - ensure all required fields exist
	input := map[string]any{
		"name":     docs[0].PageContent, // Use PageContent as the name
		"country":  "",
		"language": "",
		"currency": "",
	}

	// Copy metadata values if they exist
	if meta := docs[0].Metadata; meta != nil {
		if pop, ok := meta["population"]; ok {
			input["population"] = pop
		}
		if country, ok := meta["country"]; ok {
			input["country"] = country
		}
		if lang, ok := meta["language"]; ok {
			input["language"] = lang
		}
		if curr, ok := meta["currency"]; ok {
			input["currency"] = curr
		}
	}

	// Debug: Print the final input
	fmt.Printf("Final input to chain: %+v\n", input)

	// Execute the chain
	result, err := chains.Call(ctx, sequentialChain, input)
	if err != nil {
		return "", fmt.Errorf("error executing chain: %v", err)
	}

	// Debug: Print the raw result
	fmt.Printf("Raw result from chain: %+v\n", result)

	// Extract and return the final answer
	if respuesta, ok := result["respuesta"].(string); ok {
		return respuesta, nil
	}

	return "No se pudo generar una respuesta.", nil
}
