import com.fasterxml.jackson.databind.JsonMappingException
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.SystemExitException
import java.io.File
import java.io.OutputStreamWriter
import kotlin.system.exitProcess

fun main(args: Array<String>) {
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
                if (input != null && inputFile != null) {
                    println("Decide wether to use input or input file...")
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