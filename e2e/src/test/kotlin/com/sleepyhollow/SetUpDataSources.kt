package com.sleepyhollow

import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSpec
import org.bson.Document
import org.bson.types.ObjectId

class SetUpDataSources {

    @BeforeScenario
    fun setUp() {
        setUpMongoDb()
    }

    @BeforeSpec(tags = ["default"])
    fun setUpDefault() {
        setUpMongoDb()
    }

    private fun setUpMongoDb() {
        MongoClient.drop()
        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b004"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name0",
                    "fields" to listOf(
                        mapOf("field_type" to "multiple-text", "required" to true, "name" to "name00"),
                        mapOf("field_type" to "text", "required" to false, "name" to "name01")))))

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b001"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "fields" to listOf(
                        mapOf("field_type" to "text", "required" to true, "name" to "name10"),
                        mapOf("field_type" to "text", "required" to false, "name" to "name11")))))
    }
}