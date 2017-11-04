import java.util.*

fun InputList.rank(): Result {
    //Check that we have more than one voters
//    if (this.voters.size <= 1) {
//        return Result(this.voters.first().votes)
//    }

    //Check if any voters has double votes
    this.voters.forEach {
        if (it.hasDoubles()) {
            throw DoubleVotesException("You cannot vote for the same alternative more than once.")
        }
    }

    //Check if any vote alternatives are missing from any voter
    for (i in 1 until this.voters.size) {
        val first = this.voters[i - 1]
        val second = this.voters[i]

        if (first.missesVotes(second)) {
            throw MissingVotesException("Someone forgot to send me all votes in all voters.\nfirst: $first\nsecond: $second")
        }
    }

    this.voters.forEach { voter ->
        voter.votes.compareAllAgainstEachOther { me, enemy ->
            val value = me.victories.getOrDefault(enemy, 0)
            me.victories[enemy] = value.inc()
        }
    }

    val temp = this.voters.flatMap { it.votes.flatMap { it } }.distinct().sortedByDescending {
        it.victories.map { it.value }.reduce { sum, element -> sum + element }
    }

    return Result(listOf())
}

private fun <T> List<List<T>>.compareAllAgainstEachOther(methodToRun: (T, T) -> Unit) {
    for (index in 0 until this.size - 1) {
        this[index].forEach { vote ->
            for (index2 in index + 1 until this.size) {
                val me = vote
                this[index2].forEach { enemy ->
                    methodToRun(me, enemy)
                }
            }
        }
    }
}


class DoubleVotesException(message: String) : Exception(message)
class MissingVotesException(message: String) : Exception(message)

fun <T> Iterable<T>.shuffle(seed: Long? = null): List<T> {
    val list = this.toMutableList()
    val random = if (seed != null) Random(seed) else Random()
    Collections.shuffle(list, random)
    return list
}