# Userの更新

* SETUP: User更新データ準備

## ユーザを更新することができる
* "/v1/user/8883114bd386d8fadbd6b002"にボディ"setup/request/user_update.json"でPUTリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"SleepyHollow太郎変更"である
* レスポンスボディのJsonPath"$.id"がnullでないこと
* MongoDBの"USER"に登録されている値のJsonPath"$.name"の値が"SleepyHollow太郎変更"である