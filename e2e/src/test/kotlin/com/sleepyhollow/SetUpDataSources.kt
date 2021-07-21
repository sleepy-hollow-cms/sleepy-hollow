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
            Document(
                mapOf(
                    "_id" to "5063114bd386d8fadbd6b004",
                    "fields" to listOf(
                        mapOf("field_type" to "text", "required" to true, "name" to "name0"),
                        mapOf("field_type" to "text", "required" to false, "name" to "name1")))))
    }
}