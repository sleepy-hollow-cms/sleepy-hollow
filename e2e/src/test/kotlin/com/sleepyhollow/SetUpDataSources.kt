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

    @BeforeScenario(tags = ["Entryデータセット"])
    fun setUpEntryDataset() {
        
    }

    @Step("SETUP: User更新データ準備")
    fun setUpUserUpdate() {
        MongoClient.USER.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("8883114bd386d8fadbd6b002"),
                    "name" to "name0",
                )
            )
        )
    }

    @Step("SETUP: User削除データ準備")
    fun setUpUserDelete() {
        MongoClient.USER.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("8883114bd386d8fadbd6b001"),
                    "name" to "name0",
                )
            )
        )
    }

    @Step("SETUP: Entry削除データ準備")
    fun setUpEntryDelete() {
        MongoClient.ENTRY.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("9993114bd386d8fadbd6b009"),
                    "content_model_id" to ObjectId("9993114bd386d8fadbd6b004"),
                    "items" to listOf(
                        mapOf("value" to "for delete")
                    )
                )
            )
        )
    }
    
    @Step("SETUP: ContentModel削除データ準備")
    fun setUpContentModelDelete() {
        run {
            var contentModelId = "2263114bd386d8fadbd6b004"
            MongoClient.CONTENT_MODEL.getCollection().insertOne(
                Document(
                    mapOf(
                        "_id" to ObjectId(contentModelId),
                        "space_id" to "5063114bd386d8fadbd6b007",
                        "name" to "name0",
                        "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                        "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 47),
                        "fields" to listOf(
                            mapOf("field_type" to "text", "required" to true, "name" to "name01")
                        )
                    )
                )
            )

            MongoClient.ENTRY.getCollection().insertOne(
                Document(
                    mapOf(
                        "_id" to ObjectId("1063114bd386d8fadbd6b009"),
                        "content_model_id" to ObjectId(contentModelId),
                        "items" to listOf(
                            mapOf("value" to "for delete")
                        )
                    )
                )
            )
        }
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

    @BeforeScenario(tags = ["Space更新用データの設定"])
    fun setupSpaceForUpdate() {
        MongoClient.SPACE.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b010"),
                    "name" to "spaceName_before",
                    "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
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
                    "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                )
            )
        )
        MongoClient.SPACE.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b001"),
                    "name" to "spaceName2",
                    "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                )
            )
        )
        MongoClient.SPACE.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("5063114bd386d8fadbd6b007"),
                    "name" to "spaceName",
                    "created_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
                    "updated_at" to LocalDateTime.of(2021, 8, 2, 19, 46),
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
                    "_id" to ObjectId("1063114bd386d8fadbd6b004"),
                    "content_model_id" to ObjectId("5063114bd386d8fadbd6b002"),
                    "items" to listOf(
                        mapOf("value" to "タイトル")))))

        MongoClient.ENTRY.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b006"),
                    "content_model_id" to ObjectId("5063114bd386d8fadbd6b002"),
                    "items" to listOf(
                        mapOf("value" to "2021-08-02T19:46:00Z")))))

        MongoClient.ENTRY.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b007"),
                    "content_model_id" to ObjectId("5063114bd386d8fadbd6b003"),
                    "items" to listOf(
                        mapOf("value" to "2021-08-02T19:46:00Z")))))

        MongoClient.ENTRY.getCollection().insertOne(
            Document(
                mapOf(
                    "_id" to ObjectId("1063114bd386d8fadbd6b008"),
                    "content_model_id" to ObjectId("5063114bd386d8fadbd6b003"),
                    "items" to listOf(
                        mapOf("value" to "2021-08-02T19:46:00Z")))))
    }
}