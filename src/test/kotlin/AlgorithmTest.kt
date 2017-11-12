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

    @Test
    fun test3() {
        val k1 = Vote(UUID.randomUUID().toString(), "K1")
        val k2 = Vote(UUID.randomUUID().toString(), "k2")
        val k3 = Vote(UUID.randomUUID().toString(), "k3")
        val k4 = Vote(UUID.randomUUID().toString(), "k4")
        val k5 = Vote(UUID.randomUUID().toString(), "k5")
        val k6 = Vote(UUID.randomUUID().toString(), "k6")
        val k7 = Vote(UUID.randomUUID().toString(), "k7")
        val k8 = Vote(UUID.randomUUID().toString(), "k8")
        val k9 = Vote(UUID.randomUUID().toString(), "k9")
        val k10 = Vote(UUID.randomUUID().toString(), "k10")
        val k11 = Vote(UUID.randomUUID().toString(), "k11")
        val k12 = Vote(UUID.randomUUID().toString(), "k12")
        val k13 = Vote(UUID.randomUUID().toString(), "k13")

        val allCandidates = listOf(
                k1,
                k2,
                k3,
                k4,
                k5,
                k6,
                k7,
                k8,
                k9,
                k10,
                k11,
                k12,
                k13
        )

        fun <T> List<List<T>>.addMissingToLast(itemsToAdd: List<T>): List<List<T>> {
            val flatMapped = this.flatMap { it }
            return this.toMutableList().apply {
                add(itemsToAdd.filterNot { flatMapped.contains(it) })
            }
        }

        val voters = listOf(
                //1
                Voter(listOf(
                        listOf(k4),
                        listOf(k7),
                        listOf(k3),
                        listOf(k9),
                        listOf(k2),
                        listOf(k10),
                        listOf(k8),
                        listOf(k6),
                        listOf(k1),
                        listOf(k5)
                ).addMissingToLast(allCandidates)),
                //2
                Voter(listOf(
                        listOf(k12),
                        listOf(k11),
                        listOf(k10),
                        listOf(k3),
                        listOf(k5),
                        listOf(k4),
                        listOf(k8)
                ).addMissingToLast(allCandidates)),
                //3
                Voter(listOf(
                        listOf(k9),
                        listOf(k3),
                        listOf(k4),
                        listOf(k7),
                        listOf(k8)
                ).addMissingToLast(allCandidates)),
                //4
                Voter(listOf(
                        listOf(k11),
                        listOf(k6),
                        listOf(k10),
                        listOf(k12),
                        listOf(k7),
                        listOf(k9),
                        listOf(k2)
                ).addMissingToLast(allCandidates)),
                //5
                Voter(listOf(
                        listOf(k12),
                        listOf(k3),
                        listOf(k7)
                ).addMissingToLast(allCandidates)),
                //6
                Voter(listOf(
                        listOf(k3),
                        listOf(k5),
                        listOf(k9),
                        listOf(k4),
                        listOf(k8),
                        listOf(k10),
                        listOf(k12)
                ).addMissingToLast(allCandidates)),
                //7
                Voter(listOf(
                        listOf(k1),
                        listOf(k4),
                        listOf(k6),
                        listOf(k3),
                        listOf(k8),
                        listOf(k7),
                        listOf(k9),
                        listOf(k5),
                        listOf(k10),
                        listOf(k11),
                        listOf(k2),
                        listOf(k12)
                ).addMissingToLast(allCandidates)),
                //8
                Voter(listOf(
                        listOf(k3),
                        listOf(k12),
                        listOf(k11),
                        listOf(k7)
                ).addMissingToLast(allCandidates)),
                //9
                Voter(listOf(
                        listOf(k3),
                        listOf(k5),
                        listOf(k4),
                        listOf(k9),
                        listOf(k8),
                        listOf(k10),
                        listOf(k12)
                ).addMissingToLast(allCandidates)),
                //10
                Voter(listOf(
                        listOf(k5),
                        listOf(k1),
                        listOf(k6)
                ).addMissingToLast(allCandidates)),
                //11
                Voter(listOf(
                        listOf(k4),
                        listOf(k7),
                        listOf(k3),
                        listOf(k5),
                        listOf(k8),
                        listOf(k12),
                        listOf(k11),
                        listOf(k9)
                ).addMissingToLast(allCandidates)),
                //12
                Voter(listOf(
                        listOf(k3),
                        listOf(k12),
                        listOf(k10)
                ).addMissingToLast(allCandidates)),
                //13
                Voter(listOf(
                        listOf(k4),
                        listOf(k3),
                        listOf(k9),
                        listOf(k8),
                        listOf(k7),
                        listOf(k10),
                        listOf(k5),
                        listOf(k11)
                ).addMissingToLast(allCandidates)),
                //14
                Voter(listOf(
                        listOf(k10),
                        listOf(k9)
                ).addMissingToLast(allCandidates)),
                //15
                Voter(listOf(
                        listOf(k4),
                        listOf(k6),
                        listOf(k9)
                ).addMissingToLast(allCandidates)),
                //16
                Voter(listOf(
                        listOf(k4),
                        listOf(k3),
                        listOf(k5),
                        listOf(k6),
                        listOf(k8),
                        listOf(k2),
                        listOf(k12)
                ).addMissingToLast(allCandidates)),
                //17
                Voter(listOf(
                        listOf(k9),
                        listOf(k3),
                        listOf(k4),
                        listOf(k11),
                        listOf(k8),
                        listOf(k5)
                ).addMissingToLast(allCandidates)),
                //18
                Voter(listOf(
                        listOf(k10),
                        listOf(k11),
                        listOf(k12),
                        listOf(k9),
                        listOf(k7),
                        listOf(k3),
                        listOf(k5),
                        listOf(k6),
                        listOf(k2),
                        listOf(k8),
                        listOf(k1)
                ).addMissingToLast(allCandidates)),
                //19
                Voter(listOf(
                        listOf(k2),
                        listOf(k5),
                        listOf(k9)
                ).addMissingToLast(allCandidates)),
                //20
                Voter(listOf(
                        listOf(k4),
                        listOf(k1),
                        listOf(k3),
                        listOf(k9),
                        listOf(k7),
                        listOf(k6),
                        listOf(k5),
                        listOf(k8),
                        listOf(k12),
                        listOf(k2),
                        listOf(k11),
                        listOf(k10)
                ).addMissingToLast(allCandidates)),
                //21
                Voter(listOf(
                        listOf(k8),
                        listOf(k11),
                        listOf(k1),
                        listOf(k2),
                        listOf(k9),
                        listOf(k3),
                        listOf(k4),
                        listOf(k12),
                        listOf(k6),
                        listOf(k5),
                        listOf(k7)
                ).addMissingToLast(allCandidates)),
                //22
                Voter(listOf(
                        listOf(k3),
                        listOf(k4),
                        listOf(k6),
                        listOf(k8),
                        listOf(k9),
                        listOf(k10),
                        listOf(k1),
                        listOf(k12),
                        listOf(k11),
                        listOf(k5),
                        listOf(k7),
                        listOf(k2)
                ).addMissingToLast(allCandidates)),
                //23
                Voter(listOf(
                        listOf(k3),
                        listOf(k13),
                        listOf(k9),
                        listOf(k10),
                        listOf(k5),
                        listOf(k7),
                        listOf(k4),
                        listOf(k1),
                        listOf(k11),
                        listOf(k12),
                        listOf(k6),
                        listOf(k8),
                        listOf(k2)
                ).addMissingToLast(allCandidates)),
                //24
                Voter(listOf(
                        listOf(k7),
                        listOf(k4),
                        listOf(k5),
                        listOf(k3),
                        listOf(k12),
                        listOf(k8),
                        listOf(k9),
                        listOf(k11),
                        listOf(k6),
                        listOf(k1),
                        listOf(k10),
                        listOf(k2)
                ).addMissingToLast(allCandidates)),
                //25
                Voter(listOf(
                        listOf(k5),
                        listOf(k12),
                        listOf(k11),
                        listOf(k7),
                        listOf(k1),
                        listOf(k3),
                        listOf(k4),
                        listOf(k9),
                        listOf(k8),
                        listOf(k2),
                        listOf(k6),
                        listOf(k10)
                ).addMissingToLast(allCandidates)),
                //26
                Voter(listOf(
                        listOf(k3),
                        listOf(k7),
                        listOf(k9),
                        listOf(k4),
                        listOf(k5),
                        listOf(k10),
                        listOf(k6),
                        listOf(k1),
                        listOf(k11),
                        listOf(k8),
                        listOf(k2),
                        listOf(k12)
                ).addMissingToLast(allCandidates)),
                //27
                Voter(listOf(
                        listOf(k4)
                ).addMissingToLast(allCandidates)),
                //28
                Voter(listOf(
                        listOf(k3),
                        listOf(k4),
                        listOf(k5),
                        listOf(k8),
                        listOf(k10),
                        listOf(k6)
                ).addMissingToLast(allCandidates)),
                //29
                Voter(listOf(
                        listOf(k4)
                ).addMissingToLast(allCandidates)))

        val expectedResult = Result(listOf(
                listOf(k3),
                listOf(k4),
                listOf(k9),
                listOf(k5, k7),
                listOf(k8),
                listOf(k10),
                listOf(k12),
                listOf(k11),
                listOf(k6),
                listOf(k1),
                listOf(k2),
                listOf(k13)
        ))

        val calculated = InputList(voters).rank()

        assertEquals(expectedResult, calculated)
    }
}