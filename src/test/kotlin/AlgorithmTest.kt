import org.junit.Assert.assertEquals
import org.junit.Test
import java.util.*

class AlgorithmTest {
    companion object {
        val blue = Vote(UUID.randomUUID().toString(), "Blue")
        val red = Vote(UUID.randomUUID().toString(), "Red")
        val green = Vote(UUID.randomUUID().toString(), "Green")
        val brown = Vote(UUID.randomUUID().toString(), "Brown")
        val orange = Vote(UUID.randomUUID().toString(), "Orange")
        val banana = Vote(UUID.randomUUID().toString(), "Banana")
    }

    @Test(expected = MissingVotesException::class)
    fun testTooManyInOne() {
        val firstVoter = Voter(listOf(
                listOf(banana),
                listOf(green),
                listOf(blue, red, brown, orange)
        ))
        val secondVoter = Voter(listOf(
                listOf(green, red),
                listOf(blue),
                listOf(banana, orange)
        ))

        val inputList = InputList(listOf(
                firstVoter,
                secondVoter
        ))

        inputList.rank()
    }

    @Test
    fun test1() {
        val a = Vote(UUID.randomUUID().toString(), "A")
        val b = Vote(UUID.randomUUID().toString(), "B")
        val c = Vote(UUID.randomUUID().toString(), "C")
        val d = Vote(UUID.randomUUID().toString(), "D")


        val voters = listOf(
                Voter(listOf(
                        listOf(a),
                        listOf(b),
                        listOf(c),
                        listOf(d)
                )),
                Voter(listOf(
                        listOf(b),
                        listOf(c),
                        listOf(a),
                        listOf(d)
                )),
                Voter(listOf(
                        listOf(a),
                        listOf(d),
                        listOf(b),
                        listOf(c)
                )),
                Voter(listOf(
                        listOf(c),
                        listOf(a),
                        listOf(b),
                        listOf(d)
                )),
                Voter(listOf(
                        listOf(a),
                        listOf(c),
                        listOf(b),
                        listOf(d)
                )),
                Voter(listOf(
                        listOf(b),
                        listOf(c, d, a)
                )),
                Voter(listOf(
                        listOf(d),
                        listOf(b, c),
                        listOf(a)
                )))

        val expectedResult = Result(listOf(
                listOf(a),
                listOf(b),
                listOf(c),
                listOf(d)
        ))

        val calculated = InputList(voters).rank()

        assertEquals(expectedResult, calculated)
    }

    @Test
    fun test2() {
        val a = Vote(UUID.randomUUID().toString(), "A")
        val b = Vote(UUID.randomUUID().toString(), "B")
        val c = Vote(UUID.randomUUID().toString(), "C")
        val d = Vote(UUID.randomUUID().toString(), "D")


        val voters = listOf(
                Voter(listOf(
                        listOf(a),
                        listOf(b),
                        listOf(c),
                        listOf(d)
                )),
                Voter(listOf(
                        listOf(a),
                        listOf(b),
                        listOf(c),
                        listOf(d)
                )),
                Voter(listOf(
                        listOf(a),
                        listOf(c),
                        listOf(b),
                        listOf(d)
                )))

        val expectedResult = Result(listOf(
                listOf(a),
                listOf(b),
                listOf(c),
                listOf(d)
        ))

        val calculated = InputList(voters).rank()

        assertEquals(expectedResult, calculated)
    }
}