# Entry リッチテキスト

## Entryのフィールドに見出し1~5のテキストを保存することができる
* ContentModelにID"2063114bd386d8fadbd6b000"、field_typeが"rich-text"のデータを投入する
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_richtext_header1-5.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"2063114bd386d8fadbd6b000"である
見出し1
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].type"の値が"header1"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value"の値が"見出し1"である
見出し2
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].type"の値が"header2"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value"の値が"見出し2"である
見出し3
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[2].type"の値が"header3"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[2].value"の値が"見出し3"である
見出し4
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[3].type"の値が"header4"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[3].value"の値が"見出し4"である
見出し5
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[4].type"の値が"header5"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[4].value"の値が"見出し5"である