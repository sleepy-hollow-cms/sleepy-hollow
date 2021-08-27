# Space Get

tags: Spaceデータの設定

## Spaceの名前をID指定で取得することができる
* "/v1/spaces/5063114bd386d8fadbd6b007"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"spaceName"である

## Spaceの作成日付をID指定で取得することができる
* "/v1/spaces/5063114bd386d8fadbd6b007"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.createdAt"の日付がISO 8601形式でUTCである

## 全てのSpaceを取得することができる
* "/v1/spaces"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.[0].name"の値が"spaceName1"である
* レスポンスボディのJsonPath"$.[0].id"の値が"1063114bd386d8fadbd6b000"である
* レスポンスボディのJsonPath"$.[0].createdAt"の日付がISO 8601形式でUTCである
* レスポンスボディのJsonPath"$.[1].name"の値が"spaceName2"である
* レスポンスボディのJsonPath"$.[1].id"の値が"1063114bd386d8fadbd6b001"である
* レスポンスボディのJsonPath"$.[1].createdAt"の日付がISO 8601形式でUTCである
* レスポンスボディのJsonPath"$.[2].name"の値が"spaceName"である
* レスポンスボディのJsonPath"$.[2].id"の値が"5063114bd386d8fadbd6b007"である
* レスポンスボディのJsonPath"$.[2].createdAt"の日付がISO 8601形式でUTCである