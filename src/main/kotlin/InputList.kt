data class InputList(val voters: List<Voter>)

data class Voter(val votes: List<List<Vote>>) {
    fun hasDoubles(): Boolean {
        val list = votes.flatMap { it }

        list.forEach { vote ->
            if (list.indexOfFirst { it == vote } != list.indexOfLast { it == vote }) {
                return true
            }
        }

        return false
    }

    fun missesVotes(second: Voter): Boolean {
        val myVotes = flatMapVotes()
        val otherVotes = second.flatMapVotes()

        if (myVotes.size != otherVotes.size) {
            return true
        }

        myVotes.forEach {
            if (!otherVotes.contains(it)) {
                return true
            }
        }

        return false
    }


    private fun flatMapVotes(): List<Vote> {
        return votes.flatMap { it }
    }
}