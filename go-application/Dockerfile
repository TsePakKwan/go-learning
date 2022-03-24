FROM golang:1.15

## 作業ディレクトリ
WORKDIR /service

# モジュール管理のファイルをコピー
COPY service/go.mod .
COPY service/go.sum .

# 外部パッケージのダウンロード
RUN go mod download

EXPOSE 9000
