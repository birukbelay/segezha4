package main

import (
	// "bytes"
	// "fmt"
	"log"
	"os"
	"strconv"

	// "strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
		// ownerID   = os.Getenv("OWNER_ID")   // you must add it to your config vars
		// chatID    = os.Getenv("CHAT_ID")    // you must add it to your config vars
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
		log.Println("****")
		log.Println(q.Text)
		log.Println(q.ID)
		log.Println(q.From.ID)
		log.Println("****")
		if q.Text == "" {
			return
		}
		results := make(tb.Results, len(ArticleCases)) // []tb.Result
		for i, articleCase := range ArticleCases {
			result := &tb.ArticleResult{
				Title:       articleCase.name,
				Text:        "OK",
				HideURL:     true,
				URL:         articleCase.url,
				ThumbURL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/6/6a/External_link_font_awesome.svg/240px-External_link_font_awesome.svg.png",
				ThumbWidth:  24,
				ThumbHeight: 24,
			}
			result.SetResultID(strconv.Itoa(i))
			results[i] = result
		}
		err = b.Answer(q, &tb.QueryResponse{
			Results:           results,
			CacheTime:         60, // a minute
			SwitchPMText:      "SwitchPMText",
			SwitchPMParameter: "SwitchPMParameter",
		})
		if err != nil {
			log.Println(err)
		}
	})
	b.Handle(tb.OnChosenInlineResult, func(r *tb.ChosenInlineResult) {
		// incoming inline queries
		log.Println("====")
		log.Println(r.MessageID)
		log.Println(r.ResultID)
		log.Println(r.Query)
		log.Println(r.From.ID)
		log.Println("====")
		// ticketName := r.ResultID
		// to := tb.ChatID(parseInt(chatID))
		// commands := make([]string, 0)
		// for _, param := range strings.Split(r.Query, " ") {
		// 	if strings.HasPrefix(param, "#") || strings.HasPrefix(param, "$") {
		// 		continue
		// 	}
		// 	commands = append(commands, param)
		// }
		// if len(commands) == 0 || contains(commands, "finviz") {
		// 	screenshot := Screenshot(ticketName)
		// 	photo := &tb.Photo{
		// 		File: tb.FromReader(bytes.NewReader(screenshot)),
		// 		Caption: fmt.Sprintf(
		// 			"\\#%[1]s [finviz](https://finviz.com/quote.ashx?t=%[1]s)",
		// 			ticketName,
		// 		),
		// 	}
		// 	b.Send(
		// 		to,
		// 		photo,
		// 		&tb.SendOptions{
		// 			ParseMode: tb.ModeMarkdownV2,
		// 		},
		// 	)
		// }
		// if (len(commands) == 0 || contains(commands, "ark")) && contains(ARKTickets, ticketName) {
		// 	log.Println("OK")
		// 	log.Println("====")
		// 	b.Send(
		// 		to,
		// 		fmt.Sprintf(
		// 			"\\#%s [ARK](https://cathiesark.com/ark-combined-holdings-of-%s)",
		// 			ticketName,
		// 			strings.ToLower(ticketName),
		// 		),
		// 		&tb.SendOptions{
		// 			ParseMode: tb.ModeMarkdownV2,
		// 		},
		// 	)
		// }
	})
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
