# Space

tags: default

## Space名を登録することができる
* "/v1/spaces"にボディ"setup/request/space_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"spaceName"である
* MongoDBの"SPACE"に登録されている値のJsonPath"$.name"の値が"spaceName"である