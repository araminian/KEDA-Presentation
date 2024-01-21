package de.cloudarmin.keda.kcurd

import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController
import io.micrometer.core.instrument.MeterRegistry
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import java.util.concurrent.atomic.AtomicInteger


data class Message(val message: Int)

@RestController
class KedaController {

    private lateinit var lastMessageLength: AtomicInteger
    private var messageNumber = 0

    @GetMapping("/message")
    fun get(): String {
        return messageNumber.toString()
    }

    @Autowired
    fun setMetrics(metricRegistry: MeterRegistry) {

        //gauges -> shows the current value of a meter.
        lastMessageLength = metricRegistry.gauge("keda.message.number", AtomicInteger())!!
    }

    @PostMapping("/message")
    fun setMessage(@RequestBody message: Message): Int {
        messageNumber = message.message
        lastMessageLength.set(messageNumber)
        return messageNumber
    }
    // curl -X POST -H "Content-Type: application/json" -d '{"message": 10}' http://localhost:8081/message
    // keda_message_number
}