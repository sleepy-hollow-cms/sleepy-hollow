package com.sleepyhollow

interface TestBase {

    fun readFromFile(filePath:String) =
        this.javaClass
            .classLoader
            .getResourceAsStream(filePath)
            ?.bufferedReader()
            ?.use { it.readText() }
}