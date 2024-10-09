## これはなに

Real World HTTP 第３版を読みながら手を動かしたもの

## 内容

### 2.1.1 テストエコーサーバーの実行

practice-2-1-1に記載。
実行する場合、下記のようにする。

```
go run practice-2-1-1/main.go
```


### 4.4 GETメソッドの送信と、コンテンツ、ステータス、フィールドの受信

practice-02に記載。

初回のみ、下記でgo.modファイルを作る

```
go mod init practice-4-4
```

#### スクリプトのように実行する場合
下記を実行。

```
go run main.go
```

#### buildする場合

ビルドする

```
go build
```

実行する

```
./practice-4-4
```

