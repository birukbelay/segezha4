package main

import (
	"strings"
)

// ScreenshotMode description
type ScreenshotMode = string

// ScreenshotMode variants
const (
	ScreenshotModePage       ScreenshotMode = "Page"
	ScreenshotModeImage      ScreenshotMode = "Image"
	ScreenshotModeMarketBeat ScreenshotMode = "MarketBeat"
)

// ArticleCase struct
type ArticleCase struct {
	name           string
	linkURL        string
	imageURL       string
	screenshotMode ScreenshotMode
	x              float64
	y              float64
	width          float64
	height         float64
}

// TODO: чаще используемые перемещать наверх
// TODO: конфигурировать через интерфейс

// ArticleCases slice
var ArticleCases = []ArticleCase{
	{name: "tradingview.com", linkURL: "https://ru.tradingview.com/symbols/%s"},
	{name: "finviz.com", linkURL: "https://finviz.com/quote.ashx?t=%s",
		screenshotMode: ScreenshotModePage,
		y:              215,
		height:         845 - 91, /* (banner) */
	},
	{name: "tipranks.com", linkURL: "https://tipranks.com/stocks/%s/forecast",
		screenshotMode: ScreenshotModePage,
		x:              64,
		y:              170,
		width:          800 - 64,
		height:         913,
	},
	{name: "stockscores.com", linkURL: "https://stockscores.com/charts/charts/?ticker=%s",
		screenshotMode: ScreenshotModeImage,
		imageURL:       "https://www.stockscores.com/chart.asp?TickerSymbol=%s&TimeRange=180&Interval=d&Volume=1&ChartType=CandleStick&Stockscores=None&ChartWidth=1180&ChartHeight=590&LogScale=None&Band=None&avgType1=EMA&movAvg1=20&avgType2=EMA&movAvg2=100&Indicator1=RSI&Indicator2=None&Indicator3=MACD&Indicator4=AccDist&CompareWith=&entryPrice=&stopLossPrice=&candles=redgreen&noCache=%d",
		// width:          1200,
		// height:         1000,
		// imageURL:       "https://www.stockscores.com/chart.asp?TickerSymbol=%s&TimeRange=120&Interval=d&Volume=1&ChartType=CandleStick&Stockscores=None&ChartWidth=1200&ChartHeight=525&LogScale=None&Band=None&avgType1=EMA&movAvg1=20&avgType2=EMA&movAvg2=100&Indicator1=RSI&Indicator2=None&Indicator3=MACD&Indicator4=AccDist&CompareWith=&entryPrice=&stopLossPrice=&candles=redgreen&noCache=%d",
	},
	{name: "shortvolume.com", linkURL: "https://shortvolume.com/?t=%s",
		screenshotMode: ScreenshotModeImage,
		// width:          800,
		// height:         600,
		imageURL: "https://shortvolume.com/chart_engine/draw_chart.php?Symbol=%s&TimeRange=100&noCache=%d"},
	{name: "cathiesark.com", linkURL: "https://cathiesark.com/ark-combined-holdings-of-%s", screenshotMode: ScreenshotModePage},
	{name: "marketwatch.com", linkURL: "https://marketwatch.com/investing/stock/%s", screenshotMode: ScreenshotModePage}, // y: 345,
	// height: 565,
	{name: "marketbeat.com", linkURL: "https://marketbeat.com/stocks/%s", screenshotMode: ScreenshotModeMarketBeat},
	{name: "barchart.com", linkURL: "https://barchart.com/stocks/quotes/%s/overview", screenshotMode: ScreenshotModePage},
	{name: "gurufocus.com", linkURL: "https://gurufocus.com/stock/%s/summary", screenshotMode: ScreenshotModePage},
	{name: "finasquare.com", linkURL: "https://www.finasquare.com/stocks/%s/company-info/overview", screenshotMode: ScreenshotModePage},
	{name: "stockrow.com", linkURL: "https://stockrow.com/%s", screenshotMode: ScreenshotModePage},
	{name: "stockanalysis.com", linkURL: "https://stockanalysis.com/stocks/%s/", screenshotMode: ScreenshotModePage},
	{name: "earningswhispers.com", linkURL: "https://earningswhispers.com/stocks/%s", screenshotMode: ScreenshotModePage},
}

// GetExactArticleCase function
func GetExactArticleCase(search string) *ArticleCase {
	var result *ArticleCase
	if len(search) > 0 {
		search = strings.ToUpper(search)
		for _, articleCase := range ArticleCases {
			if strings.ToUpper(articleCase.name) == search {
				result = &articleCase
				break
			}
		}
	}
	return result
}
