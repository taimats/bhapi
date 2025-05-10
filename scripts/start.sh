#!/usr/bin/env bash

# 以下の処理の完了後に行う
# 　SSHキーをGitHubに登録
# 　リポジトリのクローン
# 　本番環境用に環境変数(.env)ファイルの配置

source ./.env
cd ${SERVER_DIR}

# go関連の処理
go mod download
go mod tidy

# マイグレーション
chmod +x ${SERVER_DIR}/scripts/updmig.sh
${SERVER_DIR}/scripts/updmig.sh

# アプリ起動
if [ $? -eq 0 ]; then
  go clean -cache && go build -o ${HOME}/bin/bhapi -trimpath -ldflags "-w -s"
  chmod +x ${HOME}/bin/bhapi
  ${HOME}/bin/bhapi
fi
