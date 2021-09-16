# Entry published / archived

tags: default

## EntryのPublishをすることができる
tags: unimplemented
* "/v1/spaces/spaceId/entries/1063114bd386d8fadbd6b004/published"にPUTリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.id"の値が"1063114bd386d8fadbd6b004"である
* レスポンスボディのJsonPath"$.publication.isPublished"の真偽値が"true"である

## EntryのUnPublishをすることができる
tags: unimplemented
* "/v1/spaces/spaceId/entries/1063114bd386d8fadbd6b004/published"にDELETEリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.publication.isPublished"の真偽値が"false"である

## 存在しないEntryにはPublish/UnPublishをすることができない
tags: unimplemented
* "/v1/spaces/spaceId/entries/9/published"にPUTリクエストを送る
* "404"ステータスコードが返ってくる
* "/v1/spaces/spaceId/entries/9/published"にDELETEリクエストを送る
* "404"ステータスコードが返ってくる

## EntryのArchiveをすることができる
tags: unimplemented
* "/v1/spaces/spaceId/entries/1063114bd386d8fadbd6b004/archived"にPUTリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.id"の値が"1063114bd386d8fadbd6b004"である
* レスポンスボディのJsonPath"$.archive.isArchived"の真偽値が"true"である

## EntryのUnArchivedをすることができる
tags: unimplemented
* "/v1/spaces/spaceId/entries/1063114bd386d8fadbd6b004/archived"にDELETEリクエストを送る
* "200"ステータスコードが返ってくる
* レスポンスボディのJsonPath"$.archive.isArchived"の真偽値が"false"である

## 存在しないEntryにはArchive/UnArchiveをすることができない
tags: unimplemented
* "/v1/spaces/spaceId/entries/9/archived"にPUTリクエストを送る
* "404"ステータスコードが返ってくる
* "/v1/spaces/spaceId/entries/9/archived"にDELETEリクエストを送る
* "404"ステータスコードが返ってくる
