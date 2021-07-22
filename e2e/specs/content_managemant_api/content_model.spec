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

## ContentModelのフィールドに名前をつけて保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_fieldname.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.fields[0].type"の値が"text"である
* レスポンスボディのJsonPath"$.fields[0].required"の真偽値が"true"である
* レスポンスボディのJsonPath"$.fields[0].name"の値が"fieldName"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].field_type"の値が"text"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].required"の真偽値が"true"である
* MongoDBの"CONTENT_MODEL"に登録されている値のJsonPath"$.fields[0].name"の値が"fieldName"である

## ID指定でContentModelを取得できる
* "/v1/spaces/spaceId/contentModels/5063114bd386d8fadbd6b004"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.fields[0].type"の値が"text"である
* レスポンスボディのJsonPath"$.fields[0].name"の値が"name0"である
* レスポンスボディのJsonPath"$.fields[0].required"の真偽値が"true"である
* レスポンスボディのJsonPath"$.fields[1].type"の値が"text"である
* レスポンスボディのJsonPath"$.fields[1].name"の値が"name1"である
* レスポンスボディのJsonPath"$.fields[1].required"の真偽値が"fase"である

## ID指定でContentModelを削除できる
* "/v1/spaces/spaceId/contentModels/5063114bd386d8fadbd6b001"にDELETEリクエストを送る
* "204"ステータスコードが返ってくる
* MongoDBの"CONTENT_MODEL"に"5063114bd386d8fadbd6b001"のIDでデータが登録されていない
