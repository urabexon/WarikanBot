# WarikanBot 💸

Slack上での割り勘を支援するBotです。<br>
イベントごとに支払者・金額・重み（%）を登録しておけば、誰が誰にいくら支払うべきかを自動で計算・投稿してくれます。

## 使い方 🧑‍💻

各イベントごとにSlackチャンネルを作成し、その中で `/warikan` コマンドを使って記録・清算を行います。

### 立替えを登録する 💰

```
/warikan register <イベント名> <支払者> <金額> [重み%]
```

- 例: `/warikan register 飲み会 山田 3000 100`
- 重みを省略した場合、デフォルトは100%になります。


### 支払い参加者の登録（重み付き）🤝

```
/warikan register <イベント名> <支払者> <金額> <重み%>
```

- 参加者の「支払い能力」「取り分」に応じて重みを調整できます。

### 清算 📊

```
/warikan settle <イベント名>
```
割り勘ボットが各参加者の支払額と重みに基づいて清算金額を計算し、「誰が誰にいくら払えばよいか」をSlackに投稿してくれます。

### リセット（やり直し） 🔄

```
/warikan reset <イベント名>
```

指定したイベントの登録情報（立替え・参加者）をすべて削除します。

## 起動方法 🚀

### 1. クローンと環境構築

```bash
git clone https://github.com/urabexon/WarikanBot.git
cd WarikanBot
go mod tidy
```

### 2. .env ファイルの作成

.env.example を元に .env を作成します。

```
SLACK_BOT_TOKEN=your-slack-bot-token
SLACK_SIGNING_SECRET=your-slack-signing-secret
```

### 3. アプリの起動

```bash
go run main.go
```

ローカルサーバーが http://localhost:5272 で起動します。

## テスト実行 🧪

```bash
go test ./...
```

## Dockerでの起動 🐳

```bash
docker build -t WarikanBot .
docker run --env-file .env -p 5272:5272 WarikanBot
```

## Slack連携手順 🔐

1. Slack APIでアプリ作成

2. Slash Command /warikan を作成（URL例: https://yourdomain.com/slack/command）

3. Event Subscriptions のURL設定＋必要なイベント（messageなど）を追加

4. Bot Token Scopes に以下を追加：
  - commands
  - chat:write
  - users:read

5. .env に Bot Token と Signing Secret を記入

6. go run main.go で起動し、Slackと連携