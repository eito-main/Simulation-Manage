# Go言語のイメージをベースにする
FROM golang:1.23.6-alpine3.20

# 必要なツールのインストール
# RUN apt-get update && apt-get install -y \
#     bash \
#     git \
#     vim \
#     && rm -rf /var/lib/apt/lists/*

# 作業ディレクトリの設定
WORKDIR /app

# COPY ./app .
COPY . .

# コンテナ起動時に実行されるコマンド
CMD ["sh", "-c", "ls && pwd && go run ./app/main.go"]