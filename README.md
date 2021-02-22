# segezha4

Telegram-bot - учебная граната на GoLang

$GOOGL #GOOGL

1. Сбор скриншотов по тикеру
2. Дайджест с других каналов по тикеру https://github.com/keselekpermen69/Telegram_Forwarder/blob/adc0ffd3aa/forwarder/modules/auto_forward.py
3. Копирование ФА в портфель по тикету https://googlesheets.medium.com/bot-happens-telegram-bot-google-sheets-on-webhooks-e415509a6213
4. парсер, чтобы собирать последние рекомендации лучших аналитиков https://www.tipranks.com/analysts/top

[Telegram Bot - how to get a group chat id?](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)

https://github.com/heroku/heroku-buildpack-google-chrome

И если стесняетесь общаться с ботом публично, то можно перейти к нему в приватный чат 😊

Если кликнуть по цветной иконке, то откроется сайт на нужном тикере. А если кликнуть на название сайта, то бот ответит сообщением с этой ссылкой. Где обозначен подарочек - к сообщению прикрепится скриншот.
Первый пункт меню (там где логотип тикера) переходит или возвращает ссылку на TradingView.com

Смысл в том, что можно тыкаться на телефоне, а не набирать требуемые сайты руками. 😊
Мне очень часто хочется поделиться скриншотом, забодался от ручного труда 😊

Бот делает скриншоты и ярлыки требуемых тикеров. Для трейдинга с телефона на диване 😊

Очень сложно сделать просто, как известно. Но я смог! 😊 Это работает в любом сообщении (если бот добавлен в группу, как админ), или в приватном чате с ботом. Добавьте к хештегам бумажек "!", или "?", или "?!". И бот Вам ответит, например: #TSLA! #TSLA? #TSLA?!

## BUGS

- "Bad Request: can't parse entities: Can't find end of Italic entity at byte offset 70 (400)"
- /info marketbeat.com M - Error R14 (Memory quota exceeded)
- параллельно обрабатывать запросы на несколько бумажек или несколько на marketbeat

## CHANGELOG

- FIXED \#AYX? не отдаёт скриншоты
- marketbeat #BABA - только Institutional Ownership > один график
- при отсутствии тикера выдавать текстовое сообщение
- marketbeat #BABA - только Institutional Ownership > надо подписывать графики и total

## QUESTIONS

- Зачем фигурные скобки? Видимо для параллельного кода

```go
{ // show time to resize
tp := time.Now()
// perform resizing
res = scaleTo(src, dr, sc.Scaler)
// report time to scaling to console
log.Printf("scaling using %q takes %v time",
sc.Name, time.Now().Sub(tp))
}
```
