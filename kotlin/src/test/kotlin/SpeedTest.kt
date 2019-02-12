import org.junit.FixMethodOrder
import org.junit.Test
import org.junit.runners.MethodSorters
import java.util.*

@FixMethodOrder(MethodSorters.NAME_ASCENDING)
class SpeedTest {
    companion object {
        private val candidates100 = createRandomCandidates(100)
        private val voters100x100 = voteRandom(candidates100, 100)

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

    //2:762
    //2:831
    //c: 1:169
    @Test
    fun test100x100() {
        InputList(voters100x100).rank()
    }

    @Test
    fun test50x4() {
        val voters50x40000 = voteRandom(createRandomCandidates(50), 4)
        InputList(voters50x40000).rank()
    }

    @Test
    fun test50x40() {
        val voters50x40000 = voteRandom(createRandomCandidates(50), 40)
        InputList(voters50x40000).rank()
    }

    @Test
    fun test50x400() {
        val voters50x40000 = voteRandom(createRandomCandidates(50), 400)
        InputList(voters50x40000).rank()
    }

    @Test
    fun test50x4000() {
        val voters50x40000 = voteRandom(createRandomCandidates(50), 4_000)
        InputList(voters50x40000).rank()
    }

    @Test
    fun test50x40000() {
        val voters50x40000 = voteRandom(createRandomCandidates(50), 4_0000)
        InputList(voters50x40000).rank()
    }

    //2:20:157
    //2:14:783
    //c: 51:342
//    @Test
//    fun test100x1000() {
//        InputList(voters100x1000).rank()
//    }
//
//    @Test
//    fun test1000x100() {
//        InputList(voters100x1000).rank()
//    }

//    @Test(timeout = 1000 * 10 * 60)
//    fun test1000x1000() {
//        InputList(voters1000x1000).rank()
//    }
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
