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

        SpecDataStore.put("statusCode", statusCode)
        SpecDataStore.put("body", body)
    }

    @Step("<path>にGETリクエストを送る")
    fun requestGet(path: String) {
        val (statusCode, body, _) = HttpClient.getRequest(
            "${Configuration[content_management_api.endpoint]}$path"
        )

        val id = JsonPath.read<String>(body, "$.id")
        SpecDataStore.put("statusCode", statusCode)
        SpecDataStore.put("body", body)
    }

    @Step("<path>にDELETEリクエストを送る")
    fun requestDelete(path: String) {
        val (statusCode, _, _) = HttpClient.deleteRequest(
            "${Configuration[content_management_api.endpoint]}$path"
        )

        SpecDataStore.put("statusCode", statusCode)
    }

    @Step("死活監視用GETリクエストを送る")
    fun requestGetLiveness() {
        val (statusCode, body, _) = HttpClient.getRequest(
            "${Configuration[content_management_api.endpoint]}/v1/systems/ping"
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

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の値が<value>である")
    fun verifyMongoDB(collection: String, jsonPath: String, value: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = MongoClient.valueOf(collection)
            .getCollection()
            .findOneById(ObjectId(id))
            ?.toJson()
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyMongoDBBoolean(collection: String, jsonPath: String, value: Boolean) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = MongoClient.valueOf(collection)
            .getCollection()
            .findOneById(ObjectId(id))
            ?.toJson()
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value

    }

    @Step("MongoDBの<collection>に<id>のIDでデータが登録されていない")
    fun verifyMongoDBNotExist(collection: String, id: String) {
        val data = MongoClient.valueOf(collection)
            .getCollection()
            .findOneById(ObjectId(id))
            ?.toJson()
        data shouldBeEqualTo null
    }
}