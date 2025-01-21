FROM golang:1.23-alpine

WORKDIR /app

# 依存関係を事前ダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードコピー
COPY . .

# ビルド
RUN go build -o main ./cmd

EXPOSE 8081

CMD ["./main"]
