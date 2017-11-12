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
    votes.compareAllAgainstEachOther { me, enemy ->
        val value = me.victories.getOrDefault(enemy, 0)
        me.victories[enemy] = value.inc()
    }
    return Result(votes.flatMap { it }.distinct().resolve())
}

val alreadyResolved = mutableListOf<List<Vote>>()

fun List<Vote>.resolve(): List<MutableList<Vote>> {
    val sortedVotes = sortedByDescending {
        it.victoryAgainst(this)
    }

    //Create a two dimensional array
    val foldedVotes: List<MutableList<Vote>> = sortedVotes.fold(mutableListOf(), { list, vote ->
        if (list.lastOrNull()?.lastOrNull()?.victoryAgainst(this) == vote.victoryAgainst(this)) {
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
            }
            results[0].realVictoriesAgainst(this).size != foldedVotes.subList(i, foldedVotes.size).flatMap { it }.size - 1 -> {
                //The vote has not won over all later votes
                toResolve.addAll(results)
            }
            else -> {
                //Let's now try to resolve the votes
                victories.addAll(toResolve.resolve())
                toResolve.clear()
                victories.add(results)
            }
        }
    }
    if (toResolve.isNotEmpty()) {
        victories.addAll(toResolve.resolve())
    }
    return victories
}

//TODO Om vinnare inte har alla vinster så sitter hen i ett cirkelberoende
// Detta gäller då även alla nästkommande, måste ha vunnit
// Om någon saknar något poäng för att ligga på rätt plats:
// Jämför alla nästkommande och se var cirkelberoendet slutar (slår eller har lika med den felande)
// Detta görs sedan inom gruppen för att se till att inbördes ordning stämmer, om inte, gör proceduren igen
// Därefter fortsätter man kontrollen efter gruppen

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