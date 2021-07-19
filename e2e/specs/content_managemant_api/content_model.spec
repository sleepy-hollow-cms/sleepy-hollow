# ContentModelを取得できる

## ContentModelを保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_create.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* ContentModelの作成でレスポンスボディが正しい形である
* ContentModelの作成でDBに登録されている値が正しい値である

## ContentModelに数値フィールドを含めて保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_with_number.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.fields[0].type"の値が"number"である
* レスポンスボディのJsonPath"$.fields[0].required"の真偽値が"true"である
* MongoDBに登録されている値のJsonPath"$.fields[0].field_type"の値が"number"である
* MongoDBに登録されている値のJsonPath"$.fields[0].required"の真偽値が"true"である

## ContentModelのフィールドに名前をつけて保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_fieldname.json"でPOSTリクエストを送る
* "201"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.fields[0].type"の値が"text"である
* レスポンスボディのJsonPath"$.fields[0].required"の真偽値が"true"である
* レスポンスボディのJsonPath"$.fields[0].name"の値が"fieldName"である
* MongoDBに登録されている値のJsonPath"$.fields[0].field_type"の値が"text"である
* MongoDBに登録されている値のJsonPath"$.fields[0].required"の真偽値が"true"である
* MongoDBに登録されている値のJsonPath"$.fields[0].name"の値が"fieldName"である