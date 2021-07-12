package com.sleepyhollow

import org.http4k.client.ApacheClient
import org.http4k.core.Method
import org.http4k.core.Request

class HttpClient {
    companion object {
        val client = ApacheClient()

        fun getRequest(url: String): Pair<Int, String> {
            val request = Request(Method.GET, url)
            val response = client(request)
            return Pair(response.status.code, response.bodyString())
        }
    }

}