# infinite-monkey-theorem

### This infinite monkey theorem implementation randomly generates a byte array of the same length as a given input file until they both match.  A number of monkeys may be specified in the command line arguments.

How to use:
1. Create a text file with the text you want the monkeys to replicate.
2. Run `go build`
3. Run `infinite-monkey-theorem.exe [fileName] [number of monkeys]` or `./infinite-monkey-theorem [fileName] [number of monkeys]` (depending on your operating system).  For example:  `infinite-monkey-theorem.exe myFile.txt 3` will read from a file in the same directory called "myFile.txt" and will create three goroutines to generate random byte arrays until one of them match the original input.
*if no number of monkeys is specified, it defaults to 1.*
