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

## 本番環境

## CI/CD

## バックエンドアーキテクチャ
