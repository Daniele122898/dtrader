package models

type CurrencyTicker struct {
	ID                string   `json:"id"`
	Currency          string   `json:"currency"`
	Symbol            string   `json:"symbol"`
	Name              string   `json:"name"`
	LogoURL           string   `json:"logo_url"`
	Status            string   `json:"status"`
	Price             string   `json:"price"`
	PriceDate         string   `json:"price_date"`
	PriceTimestamp    string   `json:"price_timestamp"`
	CirculatingSupply string   `json:"circulating_supply"`
	MaxSupply         string   `json:"max_supply"`
	MarketCap         string   `json:"market_cap"`
	NumExchanges      string   `json:"num_exchanges"`
	NumPairs          string   `json:"num_pairs"`
	NumPairsUnmapped  string   `json:"num_pairs_unmapped"`
	FirstCandle       string   `json:"first_candle"`
	FirstTrade        string   `json:"first_trade"`
	FirstOrderBook    string   `json:"first_order_book"`
	Rank              string   `json:"rank"`
	High              string   `json:"high"`
	HighTimestamp     string   `json:"high_timestamp"`
	I1H               Interval `json:"1h"`
	I1D               Interval `json:"1d"`
	I7D               Interval `json:"7d"`
	I30D              Interval `json:"30d"`
	I365D             Interval `json:"365d"`
	Ytd               Interval `json:"ytd"`
}

type Interval struct {
	Volume             string `json:"volume"`
	PriceChange        string `json:"price_change"`
	PriceChangePct     string `json:"price_change_pct"`
	VolumeChange       string `json:"volume_change"`
	VolumeChangePct    string `json:"volume_change_pct"`
	MarketCapChange    string `json:"market_cap_change"`
	MarketCapChangePct string `json:"market_cap_change_pct"`
}

type CurrencyMetadata struct {
	ID                  string `json:"id"`
	OriginalSymbol      string `json:"original_symbol"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	WebsiteURL          string `json:"website_url"`
	LogoURL             string `json:"logo_url"`
	BlogURL             string `json:"blog_url"`
	DiscordURL          string `json:"discord_url"`
	FacebookURL         string `json:"facebook_url"`
	GithubURL           string `json:"github_url"`
	MediumURL           string `json:"medium_url"`
	RedditURL           string `json:"reddit_url"`
	TelegramURL         string `json:"telegram_url"`
	TwitterURL          string `json:"twitter_url"`
	WhitepaperURL       string `json:"whitepaper_url"`
	YoutubeURL          string `json:"youtube_url"`
	LinkedinURL         string `json:"linkedin_url"`
	BitcointalkURL      string `json:"bitcointalk_url"`
	BlockExplorerURL    string `json:"block_explorer_url"`
	ReplacedBy          string `json:"replaced_by"`
	MarketsCount        string `json:"markets_count"`
	CryptocontrolCoinID string `json:"cryptocontrol_coin_id"`
}

type CurrencySparkline struct {
	Currency   string   `json:"currency"`
	Timestamps []string `json:"timestamps"`
	Prices     []string `json:"prices"`
}

