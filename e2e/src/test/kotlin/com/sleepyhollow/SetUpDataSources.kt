package com.sleepyhollow

import com.thoughtworks.gauge.BeforeScenario

class SetUpDataSources {

    @BeforeScenario
    fun setUp() {
        MongoClient.drop()
    }
}