package com.sleepyhollow

import com.natpryce.konfig.*

object Configuration {
    private val entries =
        EnvironmentVariables overriding
                ConfigurationProperties.fromResource("e2e.properties")

    operator fun <T> get(key: Key<T>): T = entries[key]
}

object content_management_api : PropertyGroup() {
    val endpoint by stringType

    object mongo: PropertyGroup() {
        val endpoint by stringType
    }
}
