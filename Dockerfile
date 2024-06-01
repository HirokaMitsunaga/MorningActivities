# ビルドステージ
FROM golang:1.21 as builder
WORKDIR /app

# 依存関係のファイルをコピーし、ダウンロード
COPY go.mod go.sum ./
COPY .env ./

RUN go mod download

# ソースコードをコピーし、ビルド
COPY . .
RUN CGO_ENABLED=0 go build -o main .

# 実行ステージ
FROM alpine:latest
WORKDIR /app

# ビルドステージから実行可能ファイルをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# アプリケーションの実行
CMD ["./main"]