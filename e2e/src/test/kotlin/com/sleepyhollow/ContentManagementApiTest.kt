package com.sleepyhollow

import com.jayway.jsonpath.JsonPath
import com.thoughtworks.gauge.Step
import com.thoughtworks.gauge.datastore.SpecDataStore
import org.amshove.kluent.shouldBeEqualTo
import org.amshove.kluent.shouldNotBe
import org.bson.types.ObjectId
import org.litote.kmongo.findOneById

class ContentManagementApiTest : TestBase {

    private val regex = Regex("(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2})\\:(\\d{2})\\:(\\d{2})Z")

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

    @Step("<path>にボディ<filePath>でPUTリクエストを送る")
    fun requestPut(path: String, filePath: String) {
        val (statusCode, body, _) = HttpClient.putRequest(
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

    @Step("レスポンスボディのJsonPath<jsonPath>の日付がISO 8601形式でUTCである")
    fun verifyBodyDateFormat(jsonPath: String) {
        val body = SpecDataStore.get("body") as String
        val date = JsonPath.read<String>(body, jsonPath)
        regex.matches(date) shouldBeEqualTo true
    }

    //レプポンスボディのJsonPath"$.created-at"とJsonPath"$.updated-at"が同じ値である

    @Step("レプポンスボディのJsonPath<jsonPath1>とJsonPath<jsonPath2>が同じ値である")
    fun verifyBodyTwoValueEqual(jsonPath1: String, jsonPath2: String) {
        val body = SpecDataStore.get("body") as String
        val value1 = JsonPath.read<String>(body, jsonPath1)
        val value2 = JsonPath.read<String>(body, jsonPath2)
        value1 shouldBeEqualTo value2
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の値が<value>である")
    fun verifyMongoDB(collection: String, jsonPath: String, value: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findDattaFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }


    @Step("MongoDBの<collection>にID<id>で登録されている値のJsonPath<jsonPath>の値が<value>である")
    fun verifyMongoDBByID(collection: String, id: String, jsonPath: String, value: String) {
        val data = findDattaFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>にID<id>で登録されてい<jsonPath>の日時が<value>でない")
    fun verifyMongoDBByIDNotEqualTo(collection: String, id: String, jsonPath: String, value: String) {
        val data = findDattaFromMongo(collection, id)
        val date = JsonPath.read<HashMap<String, String>>(data, jsonPath)
        date["\$date"].toString() shouldNotBe value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath1>とJsonPathの<jsonPath2>が同じ値である")
    fun verifyCreateDateEqualToUpdateDate(collection: String, jsonPath1: String, jsonPath2: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findDattaFromMongo(collection, id)
        val value1 = JsonPath.read<HashMap<String, String>>(data, jsonPath1)
        val value2 = JsonPath.read<HashMap<String, String>>(data, jsonPath2)
        value1 shouldNotBe value2
    }


    @Step("MongoDBの<collection>にID<id>で登録されている値のJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyMongoDBBooleanByID(collection: String, id: String, jsonPath: String, value: Boolean) {
        val data = findDattaFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>に作成日時が保存されている")
    fun verifyMongoDBDateFormat(collection: String, jsonPath: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findDattaFromMongo(collection, id)
        val date = JsonPath.read<HashMap<String, String>>(data, jsonPath)
        date["\$date"].toString() shouldNotBe null
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyMongoDBBoolean(collection: String, jsonPath: String, value: Boolean) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findDattaFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value

    }

    @Step("MongoDBの<collection>に<id>のIDでデータが登録されていない")
    fun verifyMongoDBNotExist(collection: String, id: String) =
        findDattaFromMongo(collection, id) shouldBeEqualTo null

    private fun findDattaFromMongo(collection: String, id: String) =
        MongoClient.valueOf(collection)
            .getCollection()
            .findOneById(ObjectId(id))
            ?.toJson()
}