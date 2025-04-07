# Go Task API

タスク管理のためのシンプルでモダンなREST APIです。Goによって実装され、クリーンアーキテクチャを採用しています。

## 機能

- ユーザー認証 (JWT)
- タスク管理
- クリーンアーキテクチャに基づく設計

## 技術スタック

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - 高性能なWebフレームワーク
- [GORM](https://gorm.io/) - Go用のORMライブラリ
- [PostgreSQL](https://www.postgresql.org/) - データベース
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/) - コンテナ化と開発環境

## プロジェクト構造

```
go-task-api/
├── cmd/                  # エントリーポイント
│   └── main.go
├── internal/             # 非公開パッケージ
│   ├── domain/           # ドメインモデルとリポジトリインターフェース
│   ├── usecase/          # ユースケース（アプリケーションロジック）
│   ├── interface/        # インターフェースレイヤ（コントローラ）
│   └── infrastructure/   # データベース、外部API連携など
├── docker-compose.yaml   # Docker Compose設定
└── Dockerfile            # Dockerコンテナ設定
```

## セットアップと実行

### 必要条件

- Go 1.20以上
- Docker と Docker Compose

### ローカル環境での実行

1. リポジトリをクローン

```bash
git clone https://github.com/tryuuu/go-task-api.git
cd go-task-api
```

2. 環境変数ファイルの設定

```bash
cp .env.example .env
# 必要に応じて.envファイルを編集
```

3. Docker Composeで起動

```bash
docker-compose up
```

サーバーは http://localhost:8080 で起動します。

## APIエンドポイント

### 認証

- `POST /signup` - ユーザー登録
- `POST /login` - ログイン

## 開発

- [Air](https://github.com/cosmtrek/air)を使用したホットリロード
- クリーンアーキテクチャに準拠した構造
- Dockerによる開発環境の統一

## ライセンス

MIT
