package com.sleepyhollow

import com.thoughtworks.gauge.Step

class HelloTests {

    @Step("hello")
    fun hello() {
        println("hello")
    }
}