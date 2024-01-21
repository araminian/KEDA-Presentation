package de.cloudarmin.keda.kcurd

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class KcurdApplication

fun main(args: Array<String>) {
	runApplication<KcurdApplication>(*args)
}
