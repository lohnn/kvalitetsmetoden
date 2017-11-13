import com.fasterxml.jackson.databind.JsonMappingException
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.SystemExitException
import java.io.OutputStreamWriter
import kotlin.system.exitProcess

fun main(args: Array<String>) {
    val parser = ArgParser(args)

    val jackson = jacksonObjectMapper()

    try {
        ParsedArgs(parser).run {
            parser.force()

            val inputList = jackson.readValue<InputList>(input)
            println(jackson.writeValueAsString(inputList.rank()))
        }
    } catch (e: SystemExitException) {
        e.printAndExit()
    } catch (e: JsonMappingException) {
        exitMalformedParameter()
    } catch (e: IllegalArgumentException) {
        exitMalformedParameter()
    }
}

private fun exitMalformedParameter() {
    val writer = OutputStreamWriter(System.err)
    writer.write("Input parameter malformed")
    writer.flush()
    exitProcess(3)
}