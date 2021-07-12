package com.sleepyhollow

import com.thoughtworks.gauge.Step
import org.amshove.kluent.shouldBeEqualTo

class ContentManagementApiTest {

    @Step("<path>にアクセスするとボディ<body>、ステータスコード<code>のレスポンスを返す")
    fun systemLivenessCheck(path: String, body: String, code: Int) {
        val (statusCode, resBody) = HttpClient.getRequest("${Configuration[contnt_managemant_api.endpoint]}$path")
        statusCode shouldBeEqualTo code
        resBody.trimEnd() shouldBeEqualTo body
    }
}