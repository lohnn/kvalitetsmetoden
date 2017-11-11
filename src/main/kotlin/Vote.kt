data class Vote(val uuid: String, val name: String) {
    val victories = mutableMapOf<Vote, Int>()
    val victorySum: Int by lazy {
        victories.map { it.value }.fold(0, { sum, element -> sum + element })
    }

    val realVictories: List<Vote> by lazy {
        return@lazy victories.map { it.key }
                .filter { other ->
                    victories[other] ?: 0 > other.victories[this] ?: 0
                }
    }
}