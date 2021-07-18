package com.sleepyhollow

import org.http4k.client.ApacheClient
import org.http4k.core.Body
import org.http4k.core.Method
import org.http4k.core.Request
import org.http4k.core.Response

class HttpClient {
    companion object {
        private val client = ApacheClient()

        fun getRequest(url: String): Triple<Int, String, List<Pair<String, String?>>> =
            Request(Method.GET, url)
                .let(client)
                .let { it.toTriple() }

        fun postRequest(
            url: String,
            body: String?,
            headers: List<Pair<String, String?>> = emptyList()
        ): Triple<Int, String, List<Pair<String, String?>>> =
            Request(Method.POST, url)
                .headers(headers)
                .let { if (body != null) it.body(body) else it.body(Body.EMPTY) }
                .let(client)
                .let { it.toTriple() }

        private fun Response.toTriple() =
            Triple(this.status.code, this.bodyString(), this.headers)
    }
}