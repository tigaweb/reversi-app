# reversi-app

## 利用技術

### フロントエンド

Typescript
Node.js v20.14.0
React v18.2.0

### バックエンド

Go  v1.22.3
Echo
Gorm

### データベース

MySQL 8.0.29

## 開発環境構成

docker-compose.yml

### 初期構築

Docker Desktopを利用できること

#### コンテナのビルド&起動

```bash
% docker compose up --build -d
```

#### DBの初期化

(初回のみ)スクリプトの権限付与

```bash
% ch mod +x bin/init.sh
```

初期化スクリプトの実行

```bash
% ./bin/init.sh
```

アプリケーションへの[アクセス](http://localhost:3000/)

`http://localhost:3000/`

#### Goのdebug

vscodeのリモートコンテナでbackendコンテナに接続して開発する際、vscode標準のdebugを使えるように修正した

**要件**

1 リモートコンテナ内の/app/.vscode/launch.jsonに以下追記

```json
{
  "version": "0.2.0",
  "configurations": [
        {
            "name": "Connect to server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "/app",
            "port": 2345,
            "host": "127.0.0.1",
        }
    ]
}
```

※詳細な設定は各自の設定による

2 backend/.air.tomlの下記のコメントアウトを解除

```bash
# full_bin = "dlv exec ./tmp/main --headless --listen=:2345 --api-version=2 --log"
```

3 コンテナが立ち上がっている場合は再起動

4 vscode(リモートコンテナ)でdebugを開始

注意点として、2の設定を有効にしている場合、debugを開始しないとサーバが起動しない

## 本番環境

## CI/CD

## バックエンドアーキテクチャ
