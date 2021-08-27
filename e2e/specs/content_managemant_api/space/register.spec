# Space Register

tags: default

## Space名を登録することができる
* "/v1/spaces"にボディ"setup/request/space_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"spaceName"である
* MongoDBの"SPACE"に登録されている値のJsonPath"$.name"の値が"spaceName"である

## Space登録時に作成日時が保存されている
* "/v1/spaces"にボディ"setup/request/space_create_with_timestamp.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.createdAt"の日付がISO 8601形式でUTCである
* MongoDBの"SPACE"に登録されている値のJsonPath"$.created_at"に作成日時が保存されている

## Space登録時に更新日時が保存されている
* "/v1/spaces"にボディ"setup/request/space_create_with_timestamp.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.updatedAt"の日付がISO 8601形式でUTCである
* MongoDBの"SPACE"に登録されている値のJsonPath"$.updated_at"に作成日時が保存されている