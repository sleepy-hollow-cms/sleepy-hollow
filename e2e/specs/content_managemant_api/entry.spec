# Entry

## Entryを保存することができる
* "/v1/specs/spaceId/entries"にボディ"setup/request/entry_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"5063114bd386d8fadbd6b004"である

## ContentModel IDが登録されてないものであった場合失敗エラーを返す
tags: unimplemented
* "/v1/specs/spaceId/contentModels/dummy_id/entry"にボディ"setup/request/entry_create.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる
