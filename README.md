# golang docker environment

## 初回環境構築

### イメージを作成

```
docker-compose build
```

### イメージファイル確認

```
docker-compose images
```

## 開発時

### （もしコンテナが立ち上がってなかったら）コンテナを起動

```
docker-compose up -d
```

### コンテナ状態確認

```
docker-compose ps
```

### （もしコンテナが立ち上がっててコンテナを終了したいなら）コンテナを終了

```
docker-compose down
```

### コンテナの環境に対してコマンドを叩く方法

```
docker-compose exec goapp go run main.go
# docker-compose exec サービス名 コマンド
```
