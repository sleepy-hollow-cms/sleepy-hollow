package com.sleepyhollow

import com.natpryce.konfig.*

object Configuration {
    private val entries =
        EnvironmentVariables overriding
                ConfigurationProperties.fromResource("e2e.properties")

    operator fun <T> get(key: Key<T>): T = entries[key]
}

object contnt_managemant_api : PropertyGroup() {
    val endpoint by stringType
}
