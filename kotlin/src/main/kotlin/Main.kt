import com.fasterxml.jackson.databind.JsonMappingException
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.SystemExitException
import java.io.File
import java.io.OutputStreamWriter
import java.util.*
import kotlin.system.exitProcess

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

fun main(args: Array<String>) {
//    val random = voteRandom(createRandomCandidates(349), 4_000_000)
//    val inputList = InputList(random)
//    val bytes = jacksonObjectMapper().writeValueAsBytes(inputList)
//    val file = File("test9_in.json")
//    if(file.exists()) {
//        file.delete()
//    }
//    file.writeBytes(bytes)
//    exitProcess(0)

    val parser = ArgParser(args)
    val jackson = jacksonObjectMapper()

    try {
        ParsedArgs(parser).run {
            parser.force()

            if (validate) {
                val inputList = jackson.readValue<InputList>(getText(input, inputFile))
                val outputRank = jackson.readValue<Result>(getText(outputExpected, outputFile))

                if (inputList.rank() == outputRank) {
                    println("True")
                    exitProcess(0)
                } else {
                    println("False")
                    exitProcess(1)
                }
            } else {
                kotlin.system.measureTimeMillis {
                    if (input != null && inputFile != null) {
                        println("Decide whether to use input or input file...")
                        exitProcess(1)
                    }
                    if (input == null && inputFile == null) {
                        println("You have to have one of input or input file...")
                        exitProcess(1)
                    }

                    getText(input, inputFile).let {
                        val inputList = jackson.readValue<InputList>(it)
                        output(outputFile) { jackson.writeValueAsString(inputList.rank()) }
                    }
                }.also {
                    println("Opeartion took ${it}ms")
                }
            }
        }
    } catch (e: SystemExitException) {
        e.printAndExit()
    } catch (e: JsonMappingException) {
        exitMalformedParameter()
    } catch (e: IllegalArgumentException) {
        exitMalformedParameter()
    }
}

fun output(outputFile: File?, function: () -> String) {
    outputFile?.printWriter()?.use {
        it.println(function())
    } ?: run { println(function()) }
}

fun getText(text: String?, file: File?): String {
    return text ?: file?.readText() ?: ""
}

private fun exitMalformedParameter() {
    val writer = OutputStreamWriter(System.err)
    writer.write("Input parameter malformed")
    writer.flush()
    exitProcess(3)
}