# Entry

## Entryを保存することができる
* "/v1/specs/spaceId/entries"にボディ"setup/request/entry_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"5063114bd386d8fadbd6b004"である

## ContentModel IDが登録されてないものであった場合失敗エラーを返す
* "/v1/specs/spaceId/entries"にボディ"setup/request/entry_create_with_dummy_model_id.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる

## Entryのフィールドにテキストを保存することができる
tags: unimplemented
* "/v1/specs/spaceId/entries"にボディ"setup/request/entry_with_text.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"1063114bd386d8fadbd6b001"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0:1].contentType"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0:1].name"の値が"title"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0:1].value"の値が"タイトル"である