# Entry

## Entryを保存することができる
* "/v1/specs/spaceId/contentModels/modelId/entry"にボディ"setup/request/entry_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBのEntryに保存されているのデータをSpecDataStoreにストアする
* MongoDBにIDが保存されている