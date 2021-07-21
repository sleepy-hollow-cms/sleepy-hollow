package com.sleepyhollow

import com.thoughtworks.gauge.BeforeScenario
import org.bson.Document

class SetUpDataSources {

    @BeforeScenario
    fun setUp() {
        setUpMongoDb()
    }

    private fun setUpMongoDb() {
        MongoClient.drop()
        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(mapOf(
                    "_id" to "ObjectId(\"5063114bd386d8fadbd6b004\")",
                    "fields" to listOf(mapOf("type" to "text", "required" to true)))))

    }
}