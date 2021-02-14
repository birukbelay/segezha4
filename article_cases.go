package main

import (
	"strings"
)

// ArticleCase struct
type ArticleCase struct {
	name     string
	linkURL  string
	imageURL string
}

// TODO: чаще используемые перемещать наверх
// TODO: конфигурировать через интерфейс

// TODO: map https://habr.com/ru/post/457728/
// m := make(map[key_type]value_type)
// m := new(map[key_type]value_type)
// var m map[key_type]value_type
// m := map[key_type]value_type{key1: val1, key2: val2}

// ArticleCases slice
var ArticleCases = []ArticleCase{
	{name: "finviz.com", linkURL: "https://finviz.com/quote.ashx?t=%s"},
	{name: "cathiesark.com", linkURL: "https://cathiesark.com/ark-combined-holdings-of-%s"},
	{name: "marketwatch.com", linkURL: "https://marketwatch.com/investing/stock/%s"},
	{name: "stockscores.com", linkURL: "https://stockscores.com/charts/charts/?ticker=%s",
		imageURL: "https://stockscores.com/chart.asp?TickerSymbol=%s&TimeRange=100&Interval=d&Volume=1&ChartType=CandleStick&Stockscores=None&ChartWidth=860&ChartHeight=480&LogScale=1&Band=None&avgType1=EMA&movAvg1=20&avgType2=EMA&movAvg2=50&Indicator1=RSI&Indicator2=MACD&Indicator3=AccDist&CompareWith=&entryPrice=&stopLossPrice=&candles=redgreen",
		// linkURL: "https://www.stockscores.com/chart.asp?TickerSymbol=%s&TimeRange=100&Interval=d&Volume=1&ChartType=CandleStick&Stockscores=None&ChartWidth=860&ChartHeight=480&LogScale=1&Band=None&avgType1=EMA&movAvg1=20&avgType2=EMA&movAvg2=50&Indicator1=Momentum&Indicator2=RSI&Indicator3=MACD&Indicator4=AccDist&CompareWith=&entryPrice=&stopLossPrice=&candles=redgreen",
		// linkURL: "https://stockscores.com/chart.asp?TickerSymbol=%s&TimeRange=365&Interval=d&Volume=1&ChartType=CandleStick&Stockscores=None&ChartWidth=1920&ChartHeight=430&LogScale=1&Band=None&avgType1=EMA&movAvg1=20&avgType2=EMA&movAvg2=50&Indicator1=Momentum&Indicator2=RSI&Indicator3=MACD&Indicator4=AccDist&CompareWith=&entryPrice=&stopLossPrice=&candles=redgreen",
	},
	{name: "shortvolume.com", linkURL: "https://shortvolume.com/?t=%s",
		imageURL: "https://shortvolume.com/chart_engine/draw_chart.php?Symbol=%s&TimeRange=100"},
	{name: "marketbeat.com", linkURL: "https://marketbeat.com/stocks/%s"},
	{name: "earningswhispers.com", linkURL: "https://earningswhispers.com/stocks/%s"},
	// {name: "tipranks.com", linkURL: "https://tipranks.com/stocks/%s/forecast"},
	{name: "barchart.com", linkURL: "https://barchart.com/stocks/quotes/%s/overview"},
	{name: "gurufocus.com", linkURL: "https://gurufocus.com/stock/%s/summary"},
	{name: "stockrow.com", linkURL: "https://stockrow.com/%s"},
	{name: "stockanalysis.com", linkURL: "https://stockanalysis.com/stocks/%s/"},
	{name: "finasquare.com", linkURL: "https://www.finasquare.com/stocks/%s/company-info/overview"},
}

// GetExactArticleCase function
func GetExactArticleCase(search string) *ArticleCase {
	var result *ArticleCase
	if len(search) > 0 {
		search = strings.ToUpper(search)
		for _, articleCase := range ArticleCases {
			if articleCase.name == search {
				result = &articleCase
				break
			}
		}
	}
	return result
}
