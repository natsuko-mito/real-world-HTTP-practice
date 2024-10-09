## これはなに

Real World HTTP 第３版を読みながら手を動かしたもの

## 内容

### 2.1.1 テストエコーサーバーの実行

practice-01に記載。
practice-01を実行する場合、下記のようにする。

```
go run practice-01/main01.go
```


### 4.4 GETメソッドの送信と、コンテンツ、ステータス、フィールドの受信

practice-02に記載。

初回のみ、下記でgo.modファイルを作る

```
go mod init main02
```

#### スクリプトのように実行する場合
下記を実行。

```
go run main02.go
```

#### buildする場合

ビルドする

```
go build
```

実行する

```
./main02
```

