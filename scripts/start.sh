#!/usr/bin/env bash

# 以下の処理の完了後に行う
# - SSHキーをGitHubに登録
# - リポジトリのクローン
# - 本番環境用に環境変数(.env)ファイルの配置
# - マイグレーション

source ./.env

# go関連の処理
cd ${SERVER_DIR}
go mod download
go mod tidy

# アプリ起動
go clean -cache && go build -o ${EXECUTE_PATH} -trimpath -ldflags "-w -s"
sudo chmod +x ${EXECUTE_PATH}
nohup ${EXECUTE_PATH} &
echo "サーバーが稼働中です!!"
