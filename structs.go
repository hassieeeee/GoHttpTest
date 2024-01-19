package main

//ChatGPT用リクエスト用
type OpenaiRequest struct {
	Password string `json:"password"`
	Body     Body   `json:"body"`
}

type Body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenaiResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usages  Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Messages     Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

//NEWSAPI用
type NewsAPIja_data struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		UrlToImage  string `json:"urlToImage"`
		PublishedAt string `json:"publishedAt"`
		Content     string `json:"content"`
	}
}

type CurrentWeather struct {
	Time          string  `json:"time"`
	Temperature2m float64 `json:"temperature_2m"`
	WindSpeed10m  float64 `json:"wind_speed_10m"`
}

type HourlyWeather struct {
	Time               []string  `json:"time"`
	WindSpeed10m       []float64 `json:"wind_speed_10m"`
	Temperature2m      []float64 `json:"temperature_2m"`
	RelativeHumidity2m []int     `json:"relative_humidity_2m"`
}

type WeatherData struct {
	Current CurrentWeather `json:"current"`
	Hourly  HourlyWeather  `json:"hourly"`
}
