# ContentModelを取得できる

## ContentModelを保存できる
* "/v1/spaces/space1/contentModels"にボディ"setup/request/content_model_create.json"でリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.name"の値が"name"である
* MonogoDBに"name"の値が"name"で登録されている
