package com.sleepyhollow

import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSpec
import com.thoughtworks.gauge.Step
import org.bson.Document
import org.bson.types.ObjectId
import java.time.LocalDateTime

class SetUpDataSources {

    @BeforeScenario
    fun setUp() {
        setUpMongoDb()
    }

    @BeforeSpec(tags = ["default"])
    fun setUpDefault() {
        setUpMongoDb()
    }
    
    @Step("ContentModelにID<id>、field_typeが<fieldType>のデータを投入する")
    fun setupMongoData(id: String, fieldType: String) {
        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId(id),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "content-model-name",
                    "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to fieldType, "required" to true, "name" to "name00"),
                    )
                )
            )
        )
    }

    @BeforeScenario(tags = ["Spaceデータの設定"])
    fun setupSpace() {
        MongoClient.SPACE.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b000"),
                    "name" to "spaceName1",
                )
            )
        )
        MongoClient.SPACE.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b001"),
                    "name" to "spaceName2",
                )
            )
        )
        MongoClient.SPACE.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b007"),
                    "name" to "spaceName",
                )
            )
        )
    }

    private fun setUpMongoDb() {
        MongoClient.drop()
        
        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b004"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name0",
                    "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "multiple-text", "required" to true, "name" to "name00"),
                        mapOf("field_type" to "text", "required" to false, "name" to "name01")))))

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b001"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "multiple-text", "required" to true, "name" to "name10")
                    )
                )
            )
        )

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b002"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "text", "required" to true, "name" to "value")))))

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b003"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "text", "required" to true, "name" to "value"),
                        mapOf("field_type" to "text", "required" to true, "name" to "value"),
                        mapOf("field_type" to "multiple-text", "required" to true, "name" to "value")))))
        
        MongoClient.ENTRY.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b004"),
                    "content_model_id" to ObjectId("5063114bd386d8fadbd6b002"),
                    "items" to listOf(
                        mapOf("value" to "タイトル")))))
        
        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b005"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "text", "required" to true, "name" to "value")))))

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b006"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "date", "required" to true, "name" to "value")))))
        
        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("3063114bd386d8fadbd6b007"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "number", "required" to true, "name" to "value")))))

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b008"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "date", "required" to true, "name" to "value")))))

        MongoClient.CONTENT_MODEL.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b009"),
                    "space_id" to "5063114bd386d8fadbd6b007",
                    "name" to "name1",
                    "created_at" to LocalDateTime.of(2021, 8, 3, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 3, 19, 47),
                    "fields" to listOf(
                        mapOf("field_type" to "bool", "required" to true, "name" to "value1"),
                        mapOf("field_type" to "bool", "required" to true, "name" to "value2")
                    ))))
        
        MongoClient.ENTRY.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b006"),
                    "content_model_id" to ObjectId("5063114bd386d8fadbd6b002"),
                    "items" to listOf(
                        mapOf("value" to "2021-08-02T19:46:00Z")))))
    }
}