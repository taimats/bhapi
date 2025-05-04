## 概要
WEBアプリ「BOOK HISTORY」のバックエンドREST APIサーバー

## 使用技術
|名前|説明|
---|---
|**1. Golang 1.23.1**|主要言語|
|**2. Echo 4.13.3**|Go製の軽量Webフレームワーク|
|**3. Bun 1.2.11**|Go製のORM|
|**4. oapi-codegen 2.4.1**|OpenAPIスキーマのコード生成|
|**5. Atlas 0.31.1**|マイグレーションファイルのバージョン管理に使用|
|**6. migrate 4.18.1**|マイグレーションに使用|
|**7. PostgreSQL 16**|データベース|
|**8. Render**|サーバー、データベースともにデプロイ先として利用|

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
<img src="https://github.com/user-attachments/assets/346d8ee7-391e-4190-ade3-d0763b4805ea" width="600">
