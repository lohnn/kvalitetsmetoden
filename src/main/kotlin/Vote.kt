data class Vote(val uuid: String, val name: String) {
    val victories = mutableMapOf<Vote, Int>()

    fun realVictoriesAgainst(currentCompetitors: List<Vote>): List<Vote> {
        return victories.filter { currentCompetitors.contains(it.key) }
                .map { it.key }
                .filter { other ->
                    victories[other] ?: 0 > other.victories[this] ?: 0
                }
    }
}