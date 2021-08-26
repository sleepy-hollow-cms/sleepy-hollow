# Entry Failer

## ContentModel IDが登録されてないものであった場合失敗エラーを返す
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_create_with_dummy_model_id.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる

## EntryがContentModelの形に沿っていない場合は400エラーを返す
* "/v1/spaces/spaceId/entries"にボディ"setup/request/entry_unmatch_to_content_model.json"でPOSTリクエストを送る
* "400"ステータスコードが返ってくる