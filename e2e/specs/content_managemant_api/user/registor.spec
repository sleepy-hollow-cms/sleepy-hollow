# User


## ユーザを登録することができる

* "/v1/user"にボディ"setup/request/user_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"SleepyHollow太郎"である
* レスポンスボディのJsonPath"$.id"がnullでないこと
* MongoDBの"USER"に登録されている値のJsonPath"$.name"の値が"SleepyHollow太郎"である