import com.google.gson.Gson
import java.util.*

val blue = Vote(UUID.randomUUID().toString(), "Blue")
val red = Vote(UUID.randomUUID().toString(), "Red")
val green = Vote(UUID.randomUUID().toString(), "Green")
val brown = Vote(UUID.randomUUID().toString(), "Brown")
val orange = Vote(UUID.randomUUID().toString(), "Orange")
val banana = Vote(UUID.randomUUID().toString(), "Banana")

fun main(args: Array<String>) {
    val firstVoter = Voter(listOf(
            listOf(banana),
            listOf(green),
            listOf(blue, red, brown, orange)
    ))
    val secondVoter = Voter(listOf(
            listOf(green, red),
            listOf(blue),
            listOf(banana, orange),
            listOf(brown)
    ))
    val thirdVoter = Voter(listOf(
            listOf(blue),
            listOf(red),
            listOf(green),
            listOf(brown),
            listOf(orange),
            listOf(banana)
    ))

    val inputList = InputList(listOf(
            firstVoter,
            secondVoter,
            thirdVoter
    ))

    println(Gson().toJson(inputList).toString())
    println(Gson().toJson(inputList.rank()).toString())
}