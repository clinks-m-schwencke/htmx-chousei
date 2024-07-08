[![Frontend Masters](https://static.frontendmasters.com/assets/brand/logos/full.png)](https://frontendmasters.com)

This is a companion repo for the [HTMX & Go with ThePrimeagen](https://frontendmasters.com/courses/htmx) course on [Frontend Masters](https://frontendmasters.com).

# 概要

HTMXとGoを使った、チュートリアルの実装

## インストール

- Go言語: https://go.dev/doc/install
- air: 
    - ターミナルで実行: `go install github.com/cosmtrek/air@latest`
    - .bashrc / .zshrcに追加: `alias air='$(go env GOPATH)/bin/air'

## 実行
ターミナルで`air`を実行して、ブラウザで`localhost:8888`でアクセス

## 対応内容
コンタクトリストのシンプルアプリです。DBなどがないため、サーバーをリロードしますと、コンタクトがリセットされます。

コンタクトの下記の動作ができます。
 - 追加
 - 削除

エラーハンドリングはサーバー側で確認しています。
ブラウザ側のエラーハンドリングはHTMLのハンドリングのみ。

削除の際、インジケータが表示され、削除完了したらCSSのtransitionが追加しています。
