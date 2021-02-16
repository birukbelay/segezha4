package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}
	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle(tb.OnQuery, func(q *tb.Query) {
		re := regexp.MustCompile("[^A-Za-z]")
		symbol := re.ReplaceAllString(q.Text, "")
		ticker := GetExactTicker(symbol)
		if ticker == nil {
			return
		}
		results := make(tb.Results, 1+len(ArticleCases)) // []tb.Result
		linkURL := fmt.Sprintf("https://ru.tradingview.com/symbols/%s", ticker.symbol)
		result := &tb.ArticleResult{
			Title:       ticker.symbol,
			Description: ticker.description,
			HideURL:     true,
			URL:         linkURL,
			ThumbURL:    fmt.Sprintf("https://storage.googleapis.com/iexcloud-hl37opg/api/logos/%s.png", ticker.symbol), // from stockanalysis.com
		}
		result.SetContent(&tb.InputTextMessageContent{
			Text: fmt.Sprintf(`\#%s \- [%s](%s)`,
				ticker.symbol,
				escape(ticker.description),
				linkURL,
			),
			ParseMode:      tb.ModeMarkdownV2,
			DisablePreview: true,
		})
		result.SetResultID(ticker.symbol)
		results[0] = result
		for i, articleCase := range ArticleCases {
			linkURL := fmt.Sprintf(articleCase.linkURL, ticker.symbol)
			title := articleCase.name
			if articleCase.hasGift {
				title += " 🎁"
			}
			result := &tb.ArticleResult{
				Title:       title,
				Description: ticker.symbol,
				HideURL:     true,
				URL:         linkURL,
			}
			result.SetContent(&tb.InputTextMessageContent{
				Text: fmt.Sprintf("/info %s %s",
					articleCase.name,
					ticker.symbol,
				),
				DisablePreview: true,
			})
			result.SetResultID(ticker.symbol + "=" + articleCase.name)
			results[i+1] = result
		}
		err = b.Answer(q, &tb.QueryResponse{
			Results:   results,
			CacheTime: 60,
		})
		if err != nil {
			log.Println(err)
		}
	})
	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Println("****")
		log.Println(m.Text)
		log.Println(b.Me)
		log.Println(m.Via)
		log.Println(m.Chat)
		log.Println(m.Sender)
		log.Println("****")
		// mode := ""
		// if m.Via != nil && m.Via.ID == b.Me.ID {
		// 	mode = "inline mode"
		// } else if m.Chat.Type == tb.ChatPrivate {
		// 	mode = "command mode"
		// }
		// if mode == "" {
		// 	return
		// }
		// log.Println(mode)
		if strings.HasPrefix(m.Text, "/info ") {
			re := regexp.MustCompile(",")
			payload := re.ReplaceAllString(m.Payload, " ")
			arguments := strings.Split(payload, " ")
			symbols := arguments[1:]
			if len(symbols) == 0 {
				log.Println("Empty symbols")
				return
			}
			log.Println(symbols)
			articleCaseName := arguments[0]
			articleCase := GetExactArticleCase(articleCaseName)
			if articleCase == nil {
				log.Println("Invalid command")
				return
			}
			for _, symbol := range symbols {
				if strings.HasPrefix(symbol, "#") || strings.HasPrefix(symbol, "$") {
					symbol = symbol[1:]
				}
				ticker := GetExactTicker(symbol)
				if ticker == nil {
					continue
				}
				if articleCaseName == "finviz.com" {
					linkURL := fmt.Sprintf(articleCase.linkURL, ticker.symbol)
					screenshot := Screenshot(linkURL)
					photo := &tb.Photo{
						File: tb.FromReader(bytes.NewReader(screenshot)),
						Caption: fmt.Sprintf(
							`\#%s [%s](%s)`,
							ticker.symbol,
							escape(articleCase.name),
							linkURL,
						),
					}
					sendInformer(b, m, photo)
				}
				if articleCase.imageURL != "" {
					imageURL := fmt.Sprintf(articleCase.imageURL, ticker.symbol)
					linkURL := fmt.Sprintf(articleCase.linkURL, ticker.symbol)
					photo := &tb.Photo{
						File: tb.FromURL(imageURL),
						Caption: fmt.Sprintf(
							`\#%s [%s](%s)`,
							ticker.symbol,
							escape(articleCase.name),
							linkURL,
						),
					}
					sendInformer(b, m, photo)
				}
			}
			deleteCommand(b, m)
		}
	})
	// b.Handle(tb.OnChosenInlineResult, func(r *tb.ChosenInlineResult) {
	// 	// incoming inline queries
	// 	log.Println("====")
	// 	log.Println(r.MessageID)
	// 	log.Println(r.ResultID)
	// 	log.Println(r.Query)
	// 	log.Println(r.From.ID)
	// 	log.Println("====")
	// 	resultID := strings.Split(r.ResultID, "=")
	// 	if len(resultID) != 2 {
	// 		// TODO: empty message
	// 		return
	// 	}
	// 	symbol := resultID[0]
	// 	articleCaseName := resultID[1]
	// 	log.Println(articleCaseName)
	// 	log.Println(symbol)
	// 	// ticketName := r.ResultID
	// 	// TODO: to https://core.telegram.org/bots#deep-linking-example
	// 	to := tb.ChatID(parseInt(chatID))
	// 	// commands := make([]string, 0)
	// 	// for _, param := range strings.Split(r.Query, " ") {
	// 	// 	if strings.HasPrefix(param, "#") || strings.HasPrefix(param, "$") {
	// 	// 		continue
	// 	// 	}
	// 	// 	commands = append(commands, param)
	// 	// }
	// 	articleCase := GetExactArticleCase(articleCaseName)
	// 	if articleCaseName == "finviz.com" {
	// 		linkURL := fmt.Sprintf(articleCase.linkURL, symbol)
	// 		screenshot := Screenshot(linkURL)
	// 		photo := &tb.Photo{
	// 			File: tb.FromReader(bytes.NewReader(screenshot)),
	// 			Caption: fmt.Sprintf(
	// 				`\#%s [%s](%s)`,
	// 				symbol,
	// 				escape(articleCase.name),
	// 				linkURL,
	// 			),
	// 		}
	// 		b.Send(
	// 			to,
	// 			photo,
	// 			&tb.SendOptions{
	// 				ParseMode: tb.ModeMarkdownV2,
	// 			},
	// 		)
	// 	}
	// 	if articleCase.imageURL != "" {
	// 		imageURL := fmt.Sprintf(articleCase.imageURL, symbol)
	// 		linkURL := fmt.Sprintf(articleCase.linkURL, symbol)
	// 		photo := &tb.Photo{
	// 			File: tb.FromURL(imageURL),
	// 			Caption: fmt.Sprintf(
	// 				`\#%s [%s](%s)`,
	// 				symbol,
	// 				escape(articleCase.name),
	// 				linkURL,
	// 			),
	// 		}
	// 		b.Send(
	// 			to,
	// 			photo,
	// 			&tb.SendOptions{
	// 				ParseMode: tb.ModeMarkdownV2,
	// 			},
	// 		)
	// 	}
	// 	// if (len(commands) == 0 || contains(commands, "ark")) && contains(ARKTickets, ticketName) {
	// 	// 	log.Println("OK")
	// 	// 	log.Println("====")
	// 	// 	b.Send(
	// 	// 		to,
	// 	// 		fmt.Sprintf(
	// 	// 			"\\#%s [ARK](https://cathiesark.com/ark-combined-holdings-of-%s)",
	// 	// 			ticketName,
	// 	// 			strings.ToLower(ticketName),
	// 	// 		),
	// 	// 		&tb.SendOptions{
	// 	// 			ParseMode: tb.ModeMarkdownV2,
	// 	// 		},
	// 	// 	)
	// 	// }
	// })
	b.Start()
}

func contains(slice []string, search string) bool {
	for _, element := range slice {
		if element == search {
			return true
		}
	}
	return false
}

func parseInt(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return result
}

func escape(s string) string {
	re := regexp.MustCompile("[.|-]")
	return re.ReplaceAllString(s, `\$0`)
}

func deleteCommand(b *tb.Bot, m *tb.Message) {
	err := b.Delete(
		&tb.StoredMessage{
			MessageID: strconv.Itoa(m.ID),
			ChatID:    m.Chat.ID,
		},
	)
	if err != nil {
		log.Println(err)
	}
}

func sendInformer(b *tb.Bot, m *tb.Message, photo *tb.Photo) {
	_, err := b.Send(
		tb.ChatID(m.Chat.ID),
		photo,
		&tb.SendOptions{
			ParseMode: tb.ModeMarkdownV2,
		},
	)
	if err != nil {
		log.Println(err)
	}
}
