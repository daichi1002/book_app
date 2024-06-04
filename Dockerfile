# ベースイメージとして公式のGoイメージを使用
FROM golang:latest-alpine

# ワーキングディレクトリを設定
WORKDIR /app

# Goモジュールのキャッシュを効率的に利用するためにモジュールファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/server

# デフォルトのコマンドを設定
CMD ["./main"]