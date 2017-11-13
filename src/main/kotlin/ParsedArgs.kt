import com.xenomachina.argparser.ArgParser

class ParsedArgs(parser: ArgParser) {
    val input by parser.storing(names = "-i", help = "Input your votes here in a json format")
}