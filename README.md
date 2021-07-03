# segezha4

Telegram-bot - учебная граната на GoLang

$GOOGL #GOOGL

Q: Бот может таргет от инвест домов дергать?

A: Я не придумал, как этим пользоваться для принятия решений о торговом моменте. Но есть другая идея. На примере Кати. Если повторять покупки вслед за ней, то повышается вероятность зацепить краткосрочный рост. Хочется такого же эффекта по этим данным - собирать последние рекомендации лучших аналитиков https://www.tipranks.com/analysts/top

1. Сбор скриншотов по тикеру
2. Дайджест с других каналов по тикеру https://github.com/keselekpermen69/Telegram_Forwarder/blob/adc0ffd3aa/forwarder/modules/auto_forward.py
3. Копирование ФА в портфель по тикеру https://googlesheets.medium.com/bot-happens-telegram-bot-google-sheets-on-webhooks-e415509a6213
4. парсер, чтобы собирать последние рекомендации лучших аналитиков https://www.tipranks.com/analysts/top
5. с определённым интервалом постить информер https://finviz.com/map.ashx?t=sec
6. наблюдать за https://t.me/FTD_ALGO и добавлять графики к трём зелёным кружочкам
7. кнопки с заявками как @pantini_rats
8. пересылка из Твиттера https://twitter.com/eWhispers/status/1383376573240274952

Как бы добавить в бота российские тикеры и ETF (https://etfdb.com/screener/)

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

Не знаю, как лучше. И надо ли оно вообще.
/index Почему-то $INX $DOWI показывают объемы, а $NASX - нет. Отключил объемы совсем.
/volume \#SPY \#QQQ \#DOW - там другая цена, чем по индексам, но видно объемы при той же динамике цены. Дима тоже на них смотрит.

## BUGS

- /info marketwatch.com bidy crsp pypl - повесился после "#BIDY not found"
- "Bad Request: can't parse entities: Can't find end of Italic entity at byte offset 70 (400)"
- /info marketbeat.com M - Error R14 (Memory quota exceeded)
- параллельно обрабатывать запросы на несколько бумажек или несколько на marketbeat
- go backgroundTask() не работает на heroku

## CHANGELOG

- /info finviz.com ATV - бот повесился
- /info finviz.com #TCEHY - повис
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

## Docker's steps (draft)

```
sudo groupadd docker
sudo usermod -aG docker $(whoami)
su -s $(whoami)
chmod 777 /var/run/docker.sock

docker build -t go-docker-image .
docker run -v ~/segezha4:/app -p 8080:8080 go-docker-image
docker image list
docker container ls
docker rmi -f $(docker images -f "dangling=true" -q)
docker rm -vf $(docker ps -a -q)
docker-compose up

// https://onedev.net/post/578
wget -qO- https://get.docker.com/ | sh
```

# .env

```
SEGEZHA4_CHAT_ID=-87654321
SEGEZHA4_SECRET=1234567890:XXXXYYYYZZZZ-A1B2C3B412345678901234
SEGEZHA4_ADMIN_USER_IDS=12345678
SEGEZHA4_THREADS=2
SEGEZHA4_TIMEOUT_FACTOR=125 // %
```

---

> Классный бот!

Спасибо! Визуализация данных - наше всё. Кейсы применения: обмен идеями по торговым моментам, сравнение бумажек по одинаковым информерам, принятие решения о сделке по срезу всех информеров на одной бумажке, дополненная реальность для торговых сигналов, периодичная публикация информеров о состоянии индексов, динамика бумажек в портфеле. И тд и тп. Один админ тут удалял мои информеры - не проникся, видимо. 😹 И некоторые "пользователи" моей экспериментальной группы @teslaholics стонут от засилия информеров... Оно не всем заходит. Но недавно ко мне обратился ваш подписчик с идеей добавлять картинку в одно сообщение для публикации новостей. Вы сейчас это делаете руками, видимо? Я записал себе в блокнотик, может сделаю попозже, если оно Вам надо. 😊

---

многие спрашивают как парсить с инвестинга фьючи

на истории. открываем техникал чарт

в отладке ищем такую лабуду

https://tvc4.investing.com/153f74cc92560d8c9781ea0dae109c19/1625243498/1/1/8/history?symbol=8839&resolution=15&from=1623947504&to=1625243564

вот наш линк symbol= тут айди, from= to= временной отрезок в юникстайм

https://api.investing.com/api/financialdata/table/list/8839%2C8874%2C8873%2C8864%2C8849%2C1035793%2C8884%2C956228%2C14966%2C8984?fieldmap=general.slim

те же самые номера символов, но поддерживается групповой запрос с разделителем %2C

инвестинг очень капризный, лучше группировать все, что можно сгруппировать

---
