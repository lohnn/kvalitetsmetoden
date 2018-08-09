import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.default
import java.io.File

class ParsedArgs(parser: ArgParser) {
    val validate by parser.flagging(
            "-v", "--validate",
            help = "Validate input and output")

    val inputFile by parser.storing("-s", "--source",
            help = "Source file") { File(this) }.default<File?>(null)

    val input by parser.storing(names = "-i", help = "Input your votes here in a json format").default<String?>(null)

    val outputFile by parser.storing("-d", "--destination",
            help = "Destination file") { File(this) }.default<File?>(null)

    val outputExpected by parser.storing("-e", "--expected",
            help = "Expected value").default<String?>(null)
}