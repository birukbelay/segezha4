# segezha4

Telegram-bot - учебная граната на GoLang

$GOOGL #GOOGL

1. Сбор скриншотов по тикеру
2. Дайджест с других каналов по тикеру https://github.com/keselekpermen69/Telegram_Forwarder/blob/adc0ffd3aa/forwarder/modules/auto_forward.py
3. Копирование ФА в портфель по тикету https://googlesheets.medium.com/bot-happens-telegram-bot-google-sheets-on-webhooks-e415509a6213
4. парсер, чтобы собирать последние рекомендации лучших аналитиков https://www.tipranks.com/analysts/top
5. с определённым интервалом постить информер https://finviz.com/map.ashx?t=sec

[Telegram Bot - how to get a group chat id?](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)

https://github.com/heroku/heroku-buildpack-google-chrome

И если стесняетесь общаться с ботом публично, то можно перейти к нему в приватный чат 😊

Т.е. кейс немного другой. В привате спрашиваете бота. А уже отобранные идеи выносите на суд пересылкой.

Если кликнуть по цветной иконке, то откроется сайт на нужном тикере. А если кликнуть на название сайта, то бот ответит сообщением с этой ссылкой. Где обозначен подарочек - к сообщению прикрепится скриншот.
Первый пункт меню (там где логотип тикера) переходит или возвращает ссылку на TradingView.com

Смысл в том, что можно тыкаться на телефоне, а не набирать требуемые сайты руками. 😊
Мне очень часто хочется поделиться скриншотом, забодался от ручного труда 😊

Оно работает для обмена идеями о торговых моментах.

Бот делает скриншоты и ярлыки требуемых тикеров. Для трейдинга с телефона на диване 😊

Очень сложно сделать просто, как известно. Но я смог! 😊 Это работает в любом сообщении (если бот добавлен в группу, как админ), или в приватном чате с ботом. Добавьте к хештегам бумажек "!", или "?", или "?!". И бот Вам ответит, например: #TSLA! #TSLA? #TSLA?!

можно использовать короткие команды прямо в тексте сообщений: #TSLA! - finviz #TSLA? - stockscores #TSLA?! - marketwatch 🤓

Ещё есть инлайн-режим. Введите @TickerInfoBot и тикер (через пробел). Появится список вариантов. Если нажать на цветной квадратик в списке, то откроется ссылка по тикеру, а если на тексте в списке, то бот отправит сообщение или информер (там где подарочек) в ответ.

В командном режиме можно перечислять тикеры, например: "/info finviz.com tsla zm twtr tdoc". Или перечислять короткими запросами в одном сообщении: #tsla! #zm! #twtr! #tdoc! Для других таймфреймов планирую сделать модификаторы: #zm?5m #zm?4h и т.п.

А я сделал новый режим /map

Не знаю, как лучше. И надо ли оно вообще.
/index Почему-то $INX $DOWI показывают объемы, а $NASX - нет. Отключил объемы совсем.
/volume \#SPY \#QQQ \#DOW - там другая цена, чем по индексам, но видно объемы при той же динамике цены. Дима тоже на них смотрит.

## BUGS

- "Bad Request: can't parse entities: Can't find end of Italic entity at byte offset 70 (400)"
- /info marketbeat.com M - Error R14 (Memory quota exceeded)
- параллельно обрабатывать запросы на несколько бумажек или несколько на marketbeat
- go backgroundTask() не работает на heroku

## CHANGELOG

- /vix
- /map
- /info stockcharts.com ZM
- FIXED для ADR не отдаёт информер MarketWatch
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
