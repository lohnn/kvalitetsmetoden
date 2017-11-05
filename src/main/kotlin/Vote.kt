data class Vote(val uuid: String, val name: String) {
    val victories = mutableMapOf<Vote, Int>()
    val victorySum: Int by lazy {
        victories.map { it.value }.reduce { sum, element -> sum + element }
    }
}