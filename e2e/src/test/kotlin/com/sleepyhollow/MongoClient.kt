package com.sleepyhollow

import com.mongodb.client.MongoCollection
import org.bson.Document
import org.litote.kmongo.KMongo



enum class MongoClient {

    CONTENT_MODEL {
        override fun getCollection(): MongoCollection<Document> =
            models.getCollection("content_model")
    };

    abstract fun getCollection(): MongoCollection<Document>

    companion object {
        private val client = KMongo.createClient(Configuration[content_management_api.mongo.endpoint])
        val models = client.getDatabase("models")

        fun drop() = models.drop()
    }
}

