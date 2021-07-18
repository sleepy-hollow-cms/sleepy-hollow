package com.sleepyhollow

import com.jayway.jsonpath.JsonPath
import com.thoughtworks.gauge.Step
import com.thoughtworks.gauge.datastore.SpecDataStore
import org.amshove.kluent.shouldBeEqualTo
import org.bson.types.ObjectId
import org.litote.kmongo.findOneById

class ContentManagementApiTest : TestBase {

    @Step("<path>にボディ<filePath>でPOSTリクエストを送る")
    fun requestPost(path: String, filePath: String) {
        val (statusCode, body, _) = HttpClient.postRequest(
            "${Configuration[content_management_api.endpoint]}$path",
            readFromFile(filePath),
            listOf(Pair("Content-Type", "application/json"))
        )

        val id = JsonPath.read<String>(body, "$.id")
        val storedData = MongoClient.CONTENT_MODEL.getCollection()
            .findOneById(ObjectId(id))
            ?.toJson()

        SpecDataStore.put("statusCode", statusCode)
        SpecDataStore.put("body", body)
        SpecDataStore.put("storedData", storedData)
    }

    @Step("<path>にGETリクエストを送る")
    fun requestGet(path: String) {
        val (statusCode, body, _) = HttpClient.getRequest(
            "${Configuration[content_management_api.endpoint]}$path"
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

    @Step("レスポンスボディのJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyBodyBoolean(jsonPath: String, value: Boolean) {
        val body = SpecDataStore.get("body") as String
        JsonPath.read<String>(body, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBに登録されている値のJsonPath<jsonPath>の値が<value>である")
    fun verifyMongoDB(jsonPath: String, value: String) {
        val body = SpecDataStore.get("storedData") as String
        JsonPath.read<String>(body, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBに登録されている値のJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyMongoDBBoolean(jsonPath: String, value: Boolean) {
        val body = SpecDataStore.get("storedData") as String
        JsonPath.read<String>(body, jsonPath) shouldBeEqualTo value
    }
}