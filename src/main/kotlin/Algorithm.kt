import kotlinx.coroutines.experimental.launch
import kotlinx.coroutines.experimental.runBlocking

fun InputList.rank(): Result {
    //Check that we have more than one voters
    if (this.voters.size <= 1) {
        return Result(this.voters.first().votes)
    }

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

    val votes = voters.flatMap { it.votes }
    votes.compareAllAgainstEachOtherAsyncBlocking { me, enemy ->
        synchronized(me) {
            val value = me.victories.getOrDefault(enemy, 0)
            me.victories[enemy] = value.inc()
        }
    }
    return Result(votes.flatMap { it }.distinct().resolve())
}

val alreadyResolved = mutableListOf<List<Vote>>()

fun List<Vote>.resolve(): List<MutableList<Vote>> {
    runBlocking {
        val jobs = this@resolve.map {
            launch {
                it.realVictoriesAgainst(this@resolve)
            }
        }
        jobs.forEach { it.join() }
    }

    val sortedVotes = sortedByDescending {
        it.realVictoriesAgainst(this)
    }

    //Create a two dimensional array
    val foldedVotes: List<MutableList<Vote>> = sortedVotes.fold(mutableListOf(), { list, vote ->
        if (list.lastOrNull()?.lastOrNull()?.realVictoriesAgainst(this) == vote.realVictoriesAgainst(this)) {
            list.last().add(vote)
        } else {
            list.add(mutableListOf(vote))
        }
        return@fold list
    })

    //Check if we have already tried to resolve this, if have; if we have, it means that our votes ARE on the same place
    if (alreadyResolved.contains(this)) {
        return foldedVotes
    }
    alreadyResolved.add(this)

    val victories = mutableListOf<MutableList<Vote>>()
    val toResolve = mutableListOf<Vote>()
    foldedVotes.forEachIndexed { i, results ->
        when {
            results.size > 1 -> {
                //Two votes are on the same place
                toResolve.addAll(results)
                victories.addAll(toResolve.resolve())
                toResolve.clear()
            }
            results[0].realVictoriesAgainst(this) != foldedVotes.subList(i, foldedVotes.size).flatMap { it }.size - 1 -> {
                //The vote has not won over all later votes
                toResolve.addAll(results)
            }
            else -> {
                //Let's now try to resolve the votes
                if (toResolve.isNotEmpty()) {
                    victories.addAll(toResolve.resolve())
                    toResolve.clear()
                }
                victories.add(results)
            }
        }
    }
    if (toResolve.isNotEmpty()) {
        victories.addAll(toResolve.resolve())
    }
    return victories
}

private fun <T> List<List<T>>.compareAllAgainstEachOtherAsyncBlocking(methodToRun: (T, T) -> Unit) {
    val thisList = this
    runBlocking {
        val outerJobs = (0 until thisList.size - 1).map { index ->
            launch {
                val innerJobs = thisList[index].map { vote ->
                    launch {
                        (index + 1 until thisList.size).map { index2 ->
                            val me = vote
                            thisList[index2].forEach { enemy ->
                                methodToRun(me, enemy)
                            }
                        }
                    }
                }
                innerJobs.forEach { it.join() }
            }
        }
        outerJobs.forEach { it.join() }
    }
}

class DoubleVotesException(message: String) : Exception(message)
class MissingVotesException(message: String) : Exception(message)
