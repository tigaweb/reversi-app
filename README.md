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

```bash
% docker compose up --build -d
```

backendコンテナにアタッチして以下を実行(dbのmigrate)

```bash
% go run migrate/migrate.go
```

アプリケーションへの[アクセス](http://localhost:3000/)

`http://localhost:3000/`

## 本番環境

## CI/CD

## バックエンドアーキテクチャ
