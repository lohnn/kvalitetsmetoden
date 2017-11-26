import org.junit.Test
import java.util.*

class SpeedTest {
    companion object {
        private val candidates100 = createRandomCandidates(100)
        private val voters100x100 = voteRandom(candidates100, 100)

        private val candidates1000 = createRandomCandidates(1000)
        private val voters100x1000 = voteRandom(candidates100, 1000)

        private fun voteRandom(candidates: List<List<Vote>>, amount: Int): List<Voter> {
            return (0 until amount)
                    .map {
                        Voter(candidates.randomOrder())
                    }
        }

        private fun createRandomCandidates(amount: Int): List<List<Vote>> {
            return (0 until amount)
                    .map { listOf(Vote(UUID.randomUUID().toString(), it.toString())) }
        }
    }

    @Test
    fun test100x100() {
        InputList(voters100x100).rank()
    }

    @Test
    fun test100x1000() {
        InputList(voters100x1000).rank()
    }
}

val rg: Random = Random()
private fun <E> List<E>.randomOrder(): List<E> {
    val items = this.toMutableList()
    for (i in 0 until items.size) {
        val randomPosition = rg.nextInt(items.size)
        val tmp: E = items[i]
        items[i] = items[randomPosition]
        items[randomPosition] = tmp
    }
    return items
}
