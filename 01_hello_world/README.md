## MEMO

- 実行ファイルの作成

  - go build

  - go build -o <ファイル名>

  - 各 OS ごと(mac==darwin, linux と mac は拡張子いらない)

    - GOOS=windows GOARCH=amd64 go build -o winapp.exe

    - GOOS=linux GOARCH=amd64 go build -o linuxapp

    - GOOS=darwin GOARCH=amd64 go build -o macapp
