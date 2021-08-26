# Entry get

tags: default

## IDを指定してEntryを取得できる
* "/v1/spaces/spaceId/entries/1063114bd386d8fadbd6b004"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.content-model-id"の値が"5063114bd386d8fadbd6b002"である
* レスポンスボディのJsonPath"$.items[0].value"の値が"タイトル"である

## IDを指定してEntryを取得出来なかった場合は404を返す
* "/v1/spaces/spaceId/entries/notfound"にGETリクエストを送る
* "404"ステータスコードが返ってくる