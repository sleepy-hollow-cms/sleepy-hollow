# Entry delete

tags: default

## IDを指定してEntryを削除できる
* SETUP: Entry削除データ準備
* "/v1/spaces/spaceId/entries/9993114bd386d8fadbd6b009"にDELETEリクエストを送る
* "204"ステータスコードが返ってくる
* MongoDBの"ENTRY"に"9993114bd386d8fadbd6b009"のIDでデータが登録されていない
