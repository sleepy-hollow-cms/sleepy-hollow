# Entry

tags: default

## Entryを保存することができる
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"5063114bd386d8fadbd6b004"である

## ContentModel IDが登録されてないものであった場合失敗エラーを返す
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_create_with_dummy_model_id.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる

## Entryのフィールドにテキストを保存することができる
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_text.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"5063114bd386d8fadbd6b001"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].name"の値が"title"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value"の値が"タイトル"である

## Entryのフィールドにマルチプルテキストを保存することができる
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_multipletext.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"5063114bd386d8fadbd6b001"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].type"の値が"multiple-text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].name"の値が"キーの名前"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0]"の値が"A"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1]"の値が"B"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[2]"の値が"C"である

## Entryのフィールドにマッチしないタイプと値を保存しようとすると400エラーを返す
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_unmatch_typevalue_text.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_unmatch_typevalue_multipletext.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる

## IDを指定してEntryを取得できる
* "/v1/spaces/spaceId/entries/1063114bd386d8fadbd6b004"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.content-model-id"の値が"5063114bd386d8fadbd6b002"である
* レスポンスボディのJsonPath"$.items[0].contentType"の値が"text"である
* レスポンスボディのJsonPath"$.items[0].name"の値が"title"である
* レスポンスボディのJsonPath"$.items[0].value"の値が"タイトル"である