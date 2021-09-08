package com.sleepyhollow

import com.jayway.jsonpath.JsonPath
import com.thoughtworks.gauge.Step
import com.thoughtworks.gauge.datastore.SpecDataStore
import org.amshove.kluent.shouldBeEqualTo
import org.amshove.kluent.shouldNotBe
import org.amshove.kluent.shouldNotBeEqualTo
import org.bson.types.ObjectId
import org.litote.kmongo.findOneById

class MongoAssertSteps {

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の値が<value>である")
    fun assertMongoDBValue(collection: String, jsonPath: String, value: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findJsonFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の数値が<value>である")
    fun assertMongoDBValueNumber(collection: String, jsonPath: String, value: Int) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findJsonFromMongo(collection, id)
        val read: Int = JsonPath.read(data, jsonPath)
        read shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>の真偽値が<value>である")
    fun assertMongoDBBooleanValue(collection: String, jsonPath: String, value: Boolean) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findJsonFromMongo(collection, id)
        val read: Boolean = JsonPath.read(data, jsonPath)
        read shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>にID<id>で登録されている値のJsonPath<jsonPath>の値が<value>である")
    fun verifyMongoDBByID(collection: String, id: String, jsonPath: String, value: String) {
        val data = findJsonFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>にID<id>で登録されている<jsonPath>の日時が<value>である")
    fun verifyMongoDBByIDEqualTo(collection: String, id: String, jsonPath: String, value: String) {
        val data = findJsonFromMongo(collection, id)
        val date = JsonPath.read<HashMap<String, String>>(data, jsonPath)
        date["\$date"].toString() shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>にID<id>で登録されている<jsonPath>の日時が<value>でない")
    fun verifyMongoDBByIDNotEqualTo(collection: String, id: String, jsonPath: String, value: String) {
        val data = findJsonFromMongo(collection, id)
        val date = JsonPath.read<HashMap<String, String>>(data, jsonPath)
        date["\$date"].toString() shouldNotBe value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath1>とJsonPathの<jsonPath2>が同じ値である")
    fun verifyCreateDateEqualToUpdateDate(collection: String, jsonPath1: String, jsonPath2: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findJsonFromMongo(collection, id)
        val value1 = JsonPath.read<HashMap<String, String>>(data, jsonPath1)
        val value2 = JsonPath.read<HashMap<String, String>>(data, jsonPath2)
        value1 shouldNotBe value2
    }

    @Step("MongoDBの<collection>にID<id>で登録されている値のJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyMongoDBBooleanByID(collection: String, id: String, jsonPath: String, value: Boolean) {
        val data = findJsonFromMongo(collection, id)
        JsonPath.read<String>(data, jsonPath) shouldBeEqualTo value
    }

    @Step("MongoDBの<collection>に登録されている値のJsonPath<jsonPath>に作成日時が保存されている")
    fun verifyMongoDBDateFormat(collection: String, jsonPath: String) {
        val body = SpecDataStore.get("body") as String
        val id = JsonPath.read<String>(body, "$.id")
        val data = findJsonFromMongo(collection, id)
        val date = JsonPath.read<HashMap<String, String>>(data, jsonPath)
        date["\$date"].toString() shouldNotBe null
    }

    @Step("MongoDBの<collection>に<id>のIDでデータが登録されていない")
    fun verifyMongoDBNotExist(collection: String, id: String) =
        findJsonFromMongo(collection, id) shouldBeEqualTo null

    @Step("MongoDBの<collection>に<id>のIDでデータが登録されている")
    fun verifyMongoDBExist(collection: String, id: String) =
        findJsonFromMongo(collection, id) shouldNotBeEqualTo null

    private fun findJsonFromMongo(collection: String, id: String) =
        MongoClient.valueOf(collection)
            .getCollection()
            .findOneById(ObjectId(id))
            ?.toJson()

}