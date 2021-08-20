package com.sleepyhollow

import com.jayway.jsonpath.JsonPath
import com.thoughtworks.gauge.Step
import com.thoughtworks.gauge.datastore.SpecDataStore
import org.amshove.kluent.shouldBeEqualTo
import org.amshove.kluent.shouldNotBe

class ResponseAssertSteps {
    @Step("<statusCode>ステータスコードが返ってくる")
    fun verifyStatusCode(statusCode: Int) {
        SpecDataStore.get("statusCode") shouldBeEqualTo statusCode
    }

    @Step("レスポンスボディのJsonPath<jsonPath>の値が<value>である")
    fun verifyBody(jsonPath: String, value: String) {
        val body = SpecDataStore.get("body") as String
        JsonPath.read<String>(body, jsonPath) shouldBeEqualTo value
    }

    @Step("レスポンスボディのJsonPath<jsonPath>の値が<value>でない")
    fun verifyBodyNotEqual(jsonPath: String, value: String) {
        val body = SpecDataStore.get("body") as String
        JsonPath.read<String>(body, jsonPath) shouldNotBe value
    }

    @Step("レスポンスボディのJsonPath<jsonPath>の真偽値が<value>である")
    fun verifyBodyBoolean(jsonPath: String, value: Boolean) {
        val body = SpecDataStore.get("body") as String
        JsonPath.read<String>(body, jsonPath) shouldBeEqualTo value
    }

    @Step("レスポンスボディのJsonPath<jsonPath>の日付がISO 8601形式でUTCである")
    fun verifyBodyDateFormat(jsonPath: String) {
        val regex = Regex("(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2})\\:(\\d{2})\\:(\\d{2})Z")
        val body = SpecDataStore.get("body") as String
        val date = JsonPath.read<String>(body, jsonPath)
        regex.matches(date) shouldBeEqualTo true
    }

    @Step("レスポンスボディのJsonPath<jsonPath1>とJsonPath<jsonPath2>が同じ値である")
    fun verifyBodyTwoValueEqual(jsonPath1: String, jsonPath2: String) {
        val body = SpecDataStore.get("body") as String
        val value1 = JsonPath.read<String>(body, jsonPath1)
        val value2 = JsonPath.read<String>(body, jsonPath2)
        value1 shouldBeEqualTo value2
    }
}