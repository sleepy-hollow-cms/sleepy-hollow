package com.sleepyhollow

import com.mongodb.client.MongoCollection
import org.bson.Document
import org.litote.kmongo.KMongo


enum class MongoClient {

    SPACE {
        override fun getCollection(): MongoCollection<Document> =
            space.getCollection("space")
    },

    CONTENT_MODEL {
        override fun getCollection(): MongoCollection<Document> =
            models.getCollection("content_model")
    },

    ENTRY {
        override fun getCollection(): MongoCollection<Document> =
            models.getCollection("entry")
    },

    USER {
        override fun getCollection(): MongoCollection<Document> =
            user.getCollection("user")
    };

    abstract fun getCollection(): MongoCollection<Document>

    companion object {
        private val client = KMongo.createClient(Configuration[content_management_api.mongo.endpoint])
        val models = client.getDatabase("models")
        val space = client.getDatabase("space")
        val user = client.getDatabase("user")

        fun drop() {
            models.drop()
            space.drop()
            user.drop()
        }
    }
}

