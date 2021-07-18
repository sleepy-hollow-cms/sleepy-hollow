# 死活監視

## HTTP Callでの死活管理ができる
* "/v1/systems/ping"にGETリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.ping"の値が"pong"である