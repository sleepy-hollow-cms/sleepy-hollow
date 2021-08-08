# ContentModel

## ContentModelを保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* ContentModelの作成でレスポンスボディが正しい形である
* ContentModelの作成でDBに登録されている値が正しい値である

## ContentModelに数値フィールドを含めて保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_with_number.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.fields[0].type"の値が"number"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].field_type"の値が"number"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].required"の真偽値が"true"である

## ContentModelの作成日時が保存されている
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_with_number.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.created-at"の日付がISO 8601形式でUTCである
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.created_at"に作成日時が保存されている

## ContentModelのフィールドに名前をつけて保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_fieldname.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.fields[0].type"の値が"text"である
* レスポンスボディのJsonPath"$.fields[0].required"の真偽値が"true"である
* レスポンスボディのJsonPath"$.fields[0].name"の値が"fieldName"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].field_type"の値が"text"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].required"の真偽値が"true"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].name"の値が"fieldName"である

## ContentModelを更新することができるすることができる
tags: default
* "/v1/spaces/space1/contentModels/5063114bd386d8fadbd6b004"にボディ"setup/request/content_model_update.json"でPUTリクエストを送る
* "200"ステータスコードが返ってくる
* ContentModelの更新でレスポンスボディが正しい形である
* ContentModelの更新でDBにID"5063114bd386d8fadbd6b004"で登録されている値が正しい値に変更されている

## ID指定でContentModelを取得できる
tags: default
* "/v1/spaces/spaceId/contentModels/5063114bd386d8fadbd6b004"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"name0"である
* レスポンスボディのJsonPath"$.created-at"の日付がISO 8601形式でUTCである
* レスポンスボディのJsonPath"$.fields[0].type"の値が"multiple-text"である
* レスポンスボディのJsonPath"$.fields[0].name"の値が"name00"である
* レスポンスボディのJsonPath"$.fields[0].required"の真偽値が"true"である
* レスポンスボディのJsonPath"$.fields[1].type"の値が"text"である
* レスポンスボディのJsonPath"$.fields[1].name"の値が"name01"である
* レスポンスボディのJsonPath"$.fields[1].required"の真偽値が"false"である

## Spaceに紐づくContentModelをリストで取得できる
tags: default
* "/v1/spaces/5063114bd386d8fadbd6b007/contentModels"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* ContentModelの一覧取得でレスポンスボディのContentModelリストが正しい形である

## ID指定でContentModelを削除できる
tags: default
* "/v1/spaces/spaceId/contentModels/5063114bd386d8fadbd6b001"にDELETEリクエストを送る
* "204"ステータスコードが返ってくる
* MongoDBの"CONTENT_MODEL"に"5063114bd386d8fadbd6b001"のIDでデータが登録されていない
