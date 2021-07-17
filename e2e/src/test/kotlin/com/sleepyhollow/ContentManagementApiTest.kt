package com.sleepyhollow

import com.jayway.jsonpath.JsonPath
import com.thoughtworks.gauge.Step
import com.thoughtworks.gauge.datastore.SpecDataStore
import org.amshove.kluent.shouldBeEqualTo
import org.amshove.kluent.shouldNotBe
import org.bson.Document
import org.litote.kmongo.findOne

class ContentManagementApiTest : TestBase {

    @Step("<path>にアクセスするとボディ<body>、ステータスコード<code>のレスポンスを返す")
    fun systemLivenessCheck(path: String, body: String, code: Int) {
        val (statusCode, resBody) = HttpClient.getRequest("${Configuration[content_management_api.endpoint]}$path")
        statusCode shouldBeEqualTo code
        resBody.trimEnd() shouldBeEqualTo body
    }

    @Step("<path>にボディ<filePath>でリクエストを送る")
    fun requestCreateContentModel(path: String, filePath: String) {
        val (statusCode, body, _) = HttpClient.putRequest(
            "${Configuration[content_management_api.endpoint]}$path",
            readFromFile(filePath),
            listOf(Pair("Content-Type", "application/json"))
        )

        SpecDataStore.put("statusCode", statusCode)
        SpecDataStore.put("body", body)
    }

    @Step("<statusCode>ステータスコードが返ってくる")
    fun verifyStatusCode(statusCode: Int) =
        SpecDataStore.get("statusCode") shouldBeEqualTo statusCode


    @Step("レスポンスボディのJsonPath<jsonPath>の値が<value>である")
    fun verifyBody(jsonPath: String, value: String) {
        val body = SpecDataStore.get("body") as String
        JsonPath.read<String>(body, jsonPath) shouldBeEqualTo value
    }

    data class Hello(val name: String = "henoheno")

    @Step("MonogoDBに<key>の値が<value>で登録されている")
    fun verifyMongoDBC(key: String, value: String) =
        MongoClient.CONTENT_MODEL.getCollection()
            .findOne(Document(mapOf(key to value)))
            ?.toJson() shouldNotBe null
}