package repositories

import (
	"aioperator/src/api/config"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func ConnectRetrievalCache(e *embeddings.EmbedderImpl) *redisvector.Store {
	store, err := redisvector.New(context.Background(),
		redisvector.WithConnectionURL("redis://"+config.RedisUser+":"+config.RedisPassword+"@"+config.RedisHost+":"+strconv.Itoa(config.RedisPort)),
		redisvector.WithIndexName(config.RedisDB, true),
		redisvector.WithEmbedder(e),
	)
	if err != nil {
		log.Fatal(err)
	}
	Initialization(context.Background(), store)
	return store
}

type RetrievalCache struct {
	store *redisvector.Store
}

func NewRetrievalCache(store *redisvector.Store) RetrievalCache {
	return RetrievalCache{
		store: store,
	}
}

func (r *RetrievalCache) Search(ctx context.Context, query string) ([]schema.Document, error) {
	docs, err := r.store.SimilaritySearch(ctx, query, 1,
		vectorstores.WithScoreThreshold(0.5),
	)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func Initialization(ctx context.Context, store *redisvector.Store) error {
	fmt.Println("Initializing retrieval cache")
	data := []schema.Document{
		{PageContent: "Tokyo", Metadata: map[string]any{"population": 13.96, "area": 2194, "country": "Japan", "founded": 1457, "timezone": "JST (UTC+9)", "gdp": 2.0, "elevation": 40, "language": "Japanese", "currency": "Yen (JPY)", "landmarks": "Tokyo Tower, Shibuya Crossing, Senso-ji Temple", "best_travel_period": "Late March to April (cherry blossoms) and October to November (mild weather, fall foliage)"}},
		{PageContent: "Kyoto", Metadata: map[string]any{"population": 1.46, "area": 828, "country": "Japan", "founded": 794, "timezone": "JST (UTC+9)", "gdp": 0.1, "elevation": 41, "language": "Japanese", "currency": "Yen (JPY)", "landmarks": "Kinkaku-ji, Fushimi Inari Shrine, Gion District", "best_travel_period": "March to May (spring) and October to November (fall colors)"}},
		{PageContent: "New York", Metadata: map[string]any{"population": 8.8, "area": 783.8, "country": "USA", "founded": 1624, "timezone": "EST (UTC-5)", "gdp": 1.8, "elevation": 10, "language": "English", "currency": "US Dollar (USD)", "landmarks": "Statue of Liberty, Times Square, Central Park", "best_travel_period": "April to June and September to early November (mild weather, fewer crowds)"}},
		{PageContent: "London", Metadata: map[string]any{"population": 8.98, "area": 1572, "country": "UK", "founded": 47, "timezone": "GMT (UTC+0)", "gdp": 0.8, "elevation": 35, "language": "English", "currency": "Pound Sterling (GBP)", "landmarks": "Big Ben, Tower of London, London Eye", "best_travel_period": "May to September (warmer weather) and December (Christmas markets)"}},
		{PageContent: "Paris", Metadata: map[string]any{"population": 2.16, "area": 105, "country": "France", "founded": 52, "timezone": "CET (UTC+1)", "gdp": 0.9, "elevation": 35, "language": "French", "currency": "Euro (EUR)", "landmarks": "Eiffel Tower, Louvre Museum, Notre-Dame Cathedral", "best_travel_period": "April to June and October to early November (mild weather, fewer tourists)"}},
		{PageContent: "Shanghai", Metadata: map[string]any{"population": 26.32, "area": 6340, "country": "China", "founded": 1291, "timezone": "CST (UTC+8)", "gdp": 0.7, "elevation": 4, "language": "Mandarin", "currency": "Yuan (CNY)", "landmarks": "The Bund, Shanghai Tower, Yu Garden", "best_travel_period": "March to May and September to November (comfortable temperatures)"}},
		{PageContent: "Mumbai", Metadata: map[string]any{"population": 20.96, "area": 603, "country": "India", "founded": 1507, "timezone": "IST (UTC+5:30)", "gdp": 0.3, "elevation": 14, "language": "Marathi, Hindi", "currency": "Indian Rupee (INR)", "landmarks": "Gateway of India, Chhatrapati Shivaji Terminus, Marine Drive", "best_travel_period": "November to February (cooler, dry season)"}},
		{PageContent: "Santiago", Metadata: map[string]any{"population": 6.9, "area": 641, "country": "Chile", "founded": 1541, "timezone": "CLT (UTC-4)", "gdp": 0.12, "elevation": 570, "language": "Spanish", "currency": "Chilean Peso (CLP)", "landmarks": "Cerro San Cristóbal, La Moneda Palace, Santa Lucía Hill", "best_travel_period": "September to November (spring) and March to May (fall)"}},
		{PageContent: "Buenos Aires", Metadata: map[string]any{"population": 15.37, "area": 203, "country": "Argentina", "founded": 1536, "timezone": "ART (UTC-3)", "gdp": 0.36, "elevation": 25, "language": "Spanish", "currency": "Argentine Peso (ARS)", "landmarks": "Obelisco, La Boca, Recoleta Cemetery", "best_travel_period": "March to May and September to December (mild weather)"}},
		{PageContent: "Rio de Janeiro", Metadata: map[string]any{"population": 13.63, "area": 1200, "country": "Brazil", "founded": 1565, "timezone": "BRT (UTC-3)", "gdp": 0.18, "elevation": 31, "language": "Portuguese", "currency": "Brazilian Real (BRL)", "landmarks": "Christ the Redeemer, Sugarloaf Mountain, Copacabana Beach", "best_travel_period": "April to June and September to November (shoulder seasons with good weather)"}},
		{PageContent: "Sao Paulo", Metadata: map[string]any{"population": 22.43, "area": 1521, "country": "Brazil", "founded": 1554, "timezone": "BRT (UTC-3)", "gdp": 0.2, "elevation": 760, "language": "Portuguese", "currency": "Brazilian Real (BRL)", "landmarks": "Paulista Avenue, Ibirapuera Park, São Paulo Cathedral", "best_travel_period": "March to May and October to November (mild temperatures, less rain)"}},
		{PageContent: "Singapore", Metadata: map[string]any{"population": 5.7, "area": 728, "country": "Singapore", "founded": 1819, "timezone": "SGT (UTC+8)", "gdp": 0.4, "elevation": 15, "language": "English, Malay, Mandarin, Tamil", "currency": "Singapore Dollar (SGD)", "landmarks": "Marina Bay Sands, Gardens by the Bay, Merlion Park", "best_travel_period": "February to April (drier months, though warm year-round)"}},
		{PageContent: "Sydney", Metadata: map[string]any{"population": 5.31, "area": 12367, "country": "Australia", "founded": 1788, "timezone": "AEST (UTC+10)", "gdp": 0.5, "elevation": 58, "language": "English", "currency": "Australian Dollar (AUD)", "landmarks": "Sydney Opera House, Harbour Bridge, Bondi Beach", "best_travel_period": "September to November and March to May (mild temperatures)"}},
		{PageContent: "Cape Town", Metadata: map[string]any{"population": 4.71, "area": 2461, "country": "South Africa", "founded": 1652, "timezone": "SAST (UTC+2)", "gdp": 0.12, "elevation": 5, "language": "Afrikaans, English, Xhosa", "currency": "South African Rand (ZAR)", "landmarks": "Table Mountain, Robben Island, Victoria & Alfred Waterfront", "best_travel_period": "March to May and September to November (shoulder seasons with good weather)"}},
		{PageContent: "Dubai", Metadata: map[string]any{"population": 3.4, "area": 4114, "country": "UAE", "founded": 1833, "timezone": "GST (UTC+4)", "gdp": 0.18, "elevation": 0, "language": "Arabic", "currency": "UAE Dirham (AED)", "landmarks": "Burj Khalifa, Palm Jumeirah, Dubai Mall", "best_travel_period": "November to March (cooler temperatures, outdoor activities possible)"}},
		{PageContent: "Toronto", Metadata: map[string]any{"population": 2.93, "area": 630, "country": "Canada", "founded": 1793, "timezone": "EST (UTC-5)", "gdp": 0.3, "elevation": 76, "language": "English, French", "currency": "Canadian Dollar (CAD)", "landmarks": "CN Tower, Royal Ontario Museum, Ripley's Aquarium", "best_travel_period": "May to September (warm weather, outdoor festivals)"}},
		{PageContent: "Bangkok", Metadata: map[string]any{"population": 10.54, "area": 1569, "country": "Thailand", "founded": 1782, "timezone": "ICT (UTC+7)", "gdp": 0.15, "elevation": 2, "language": "Thai", "currency": "Baht (THB)", "landmarks": "Grand Palace, Wat Arun, Khao San Road", "best_travel_period": "November to February (cooler, dry season)"}},
		{PageContent: "Berlin", Metadata: map[string]any{"population": 3.77, "area": 892, "country": "Germany", "founded": 1237, "timezone": "CET (UTC+1)", "gdp": 0.2, "elevation": 34, "language": "German", "currency": "Euro (EUR)", "landmarks": "Brandenburg Gate, Berlin Wall, Reichstag Building", "best_travel_period": "May to September (warmer weather, outdoor events)"}},
		{PageContent: "Moscow", Metadata: map[string]any{"population": 12.65, "area": 2511, "country": "Russia", "founded": 1147, "timezone": "MSK (UTC+3)", "gdp": 0.25, "elevation": 156, "language": "Russian", "currency": "Russian Ruble (RUB)", "landmarks": "Red Square, Kremlin, St. Basil's Cathedral", "best_travel_period": "May to September (mild weather, white nights in June)"}},
	}
	_, err := store.AddDocuments(ctx, data)
	if err != nil {
		fmt.Println("Error adding documents to cache", err)
		return err
	}
	fmt.Println("Documents added to cache")
	return nil
}
