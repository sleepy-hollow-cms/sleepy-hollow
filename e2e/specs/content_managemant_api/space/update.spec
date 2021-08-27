# Space Register

tags: Space更新用データの設定

## Spaceを更新することができる
* "/v1/spaces/1063114bd386d8fadbd6b010"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"spaceName_before"である
* "/v1/spaces/1063114bd386d8fadbd6b010"にボディ"setup/request/space_update.json"でPUTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"spaceName_update"である
* レスポンスボディのJsonPath"$.updatedAt"の日付がISO 8601形式でUTCである
* MongoDBの"SPACE"に登録されている値のJsonPath"$.updated_at"に作成日時が保存されている