# Entry Rich-Text

## Entryのリッチテキストフィールド内の見出し1~5テキストを保存することができる
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

## Entryのリッチテキストフィールド内の段落に文字装飾されたテキストを保存することができる
* ContentModelにID"2063114bd386d8fadbd6b000"、field_typeが"rich-text"のデータを投入する
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_richtext_paragraph_decoration.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"2063114bd386d8fadbd6b000"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].type"の値が"paragraph"である
通常段落
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].value"の値が"段落テキスト"である
太字
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[1].type"の値が"bold"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[1].value"の値が"太字"である
下線
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[2].type"の値が"underline"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[2].value"の値が"下線"である
打ち消し線
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[3].type"の値が"strike"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[3].value"の値が"打ち消し線"である
斜体
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[4].type"の値が"italics"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[4].value"の値が"斜体"である
コード
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[5].type"の値が"code"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[5].value"の値が"コード"である
引用
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[6].type"の値が"blockquote"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[6].value"の値が"引用"である

## Entryのリッチテキストフィールド内の複数段落に文字装飾されたテキストを保存することができる
* ContentModelにID"2063114bd386d8fadbd6b000"、field_typeが"rich-text"のデータを投入する
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_richtext_multi-paragraph.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"2063114bd386d8fadbd6b000"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].type"の値が"paragraph"である
1段落目
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].value"の値が"1段落目"である
2段落目
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[0].type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[0].value"の値が"2段落目"である

## Entryのリッチテキストフィールド内の段落にハイパーリンクを含んだテキストを保存することができる
* ContentModelにID"2063114bd386d8fadbd6b000"、field_typeが"rich-text"のデータを投入する
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_richtext_paragraph_hyperlink.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"2063114bd386d8fadbd6b000"である
1段落目
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].type"の値が"paragraph"である
リンク文字
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].type"の値が"hyperlink"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].value.link-target"の値が"https://sleepy-hollow.io/"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].value.link-text"の値が"リンク文字"である
2段落目
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].type"の値が"paragraph"である
左辺
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[0].type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[0].value"の値が"left-text"である
リンク文字
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[1].type"の値が"hyperlink"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[1].value.link-target"の値が"https://sleepy-hollow.io/"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[1].value.link-text"の値が"リンク文字"である
右辺
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[2].type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[1].value[2].value"の値が"right-text"である

## Entryのリッチテキストフィールド内にリストを含んだテキストを保存することができる
* ContentModelにID"2063114bd386d8fadbd6b000"、field_typeが"rich-text"のデータを投入する
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_with_richtext_list.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.content_model_id"の値が"2063114bd386d8fadbd6b000"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].type"の値が"list"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].type"の値が"list-item"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].value.type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[0].value.value"の値が"リスト1"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[1].type"の値が"list-item"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[1].value.type"の値が"text"である
* MongoDBの"ENTRY"に登録されている値のJsonPath"$.items[0].value[0].value[1].value.value"の値が"リスト2"である