data class Vote(val uuid: String, val name: String) {
    val victories = mutableMapOf<Vote, Int>()
}