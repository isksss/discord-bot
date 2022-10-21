# discord-bot

## 実行方法
1. ファイル名の変更  
`config.yml_sample` -> `config.yml`
1. config.yml内の`url`にDiscordのwebhookURLを追記する。
1. `go run main.go`で実行

## config.ymlについて
|項目|説明|
|:-:|:--|
|url|DiscordのWebHookURL|
|name|WebHookの名前|
|message|送信するメッセージs|
