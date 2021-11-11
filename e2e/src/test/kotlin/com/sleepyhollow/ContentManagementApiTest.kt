package com.sleepyhollow

import com.thoughtworks.gauge.Step
import com.thoughtworks.gauge.datastore.SpecDataStore

class ContentManagementApiTest : TestBase {
    
    @Step("<path>にボディ<filePath>でPOSTリクエストを送る")
    fun requestPost(path: String, filePath: String) {
        val (statusCode, body, _) = HttpClient.postRequest(
            "${Configuration[content_management_api.endpoint]}$path",
            readFromFile(filePath),
            listOf(Pair(HttpClient.CONTENT_TYPE, HttpClient.APPLICATION_JSON))
        )
        SpecDataStore.put("statusCode", statusCode)
        SpecDataStore.put("body", body)
    }

    @Step("<path>にPUTリクエストを送る")
    fun requestPut(path: String) {
        val (statusCode, body, _) = HttpClient.putRequest(
            "${Configuration[content_management_api.endpoint]}$path",
            null,
            listOf(Pair(HttpClient.CONTENT_TYPE, HttpClient.APPLICATION_JSON))
        )
        SpecDataStore.put("statusCode", statusCode)
        SpecDataStore.put("body", body)
    }
    
    @Step("<path>にボディ<filePath>でPUTリクエストを送る")
    fun requestPutWithBody(path: String, filePath: String) {
        val (statusCode, body, _) = HttpClient.putRequest(
            "${Configuration[content_management_api.endpoint]}$path",
            readFromFile(filePath),
            listOf(Pair(HttpClient.CONTENT_TYPE, HttpClient.APPLICATION_JSON))
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
}