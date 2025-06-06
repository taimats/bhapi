## 概要
WEBアプリ「BOOK HISTORY」のバックエンドREST APIサーバー

## 使用技術
|名前|説明|
---|---
|**01. Golang 1.23.1**|主要言語|
|**02. Echo 4.13.3**|Go製の軽量Webフレームワーク|
|**03. Bun 1.2.11**|Go製のORM|
|**04. oapi-codegen 2.4.1**|OpenAPIスキーマのコード生成|
|**05. Atlas 0.31.1**|マイグレーションファイルのバージョン管理に使用|
|**06. migrate 4.18.1**|マイグレーションに使用|
|**07. PostgreSQL 16**|データベース|
|**08. AWS**|インフラを構築|
|**09. Terraform 1.11.4**|インフラをコードで管理|
|**10. GitHub Actions**|CI/CDを実現|

## サーバーアーキテクチャ
ヘキサゴナルアーキテクチャ（をイメージ）

<img src="https://github.com/user-attachments/assets/4951eb7e-d040-4942-8d04-a31bb8c88eb3" width="600">

## 機能一覧
|メソッド|エンドポイント|機能|認証の有無|
|-------|-------------|----|---------|
|GET|/health|サーバーの監視|無
|GET|/health/db|DBの監視|無
|POST|/auth/register|ユーザー登録|認証キー
|GET|/users/{id}|ユーザー情報を取得|認証キー
|DELETE|/users/{id}|ユーザー情報を削除|認証キー
|PUT|/users|ユーザー情報を更新|認証キー
|GET|/records/{id}|記録の取得|認証キー
|GET|/charts/{id}|図表の取得|認証キー
|GET|/shelf/{id}|本棚の取得|認証キー
|PUT|/shelf/{id}|本棚の更新|認証キー
|POST|/shelf/{id}|本棚に本を追加|認証キー
|DELETE|/shelf/{id}|本棚の本を削除|認証キー
|GET|/search|書籍の検索結果を取得|認証キー

## インフラアーキテクチャ
Terraformを通じてAWSで構築

<img src="https://github.com/user-attachments/assets/b5360d49-6c84-4010-9e0a-91c54726c3a3" width="800">
