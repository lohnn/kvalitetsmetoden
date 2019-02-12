import com.fasterxml.jackson.annotation.JsonIgnore

data class Vote(val uuid: String, val name: String) {
    @get:JsonIgnore
    val victories by lazy { mutableMapOf<Vote, Int>() }

    fun realVictoriesAgainst(currentCompetitors: List<Vote>): Int {
        return victories.filter { currentCompetitors.contains(it.key) }
                .map { it.key }
                .filter { other ->
                    victories[other] ?: 0 > other.victories[this] ?: 0
                }.size
    }
}

fun <T> printTimeAndReturn(message: String? = "", function: () -> T): T {
    measureTimeMillis { }
    val (toReturn, time) = measureTimeMillis { function() }
    println("$message ${time}ms")
    return toReturn
}

private inline fun <T> measureNanoTime(block: () -> T): Pair<T, Long> {
    val start = System.nanoTime()
    return block() to System.nanoTime() - start
}

private inline fun <T> measureTimeMillis(block: () -> T): Pair<T, Long> {
    val start = System.currentTimeMillis()
    return block() to System.currentTimeMillis() - start
}

