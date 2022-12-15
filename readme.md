
## Hello, **Username** !

Here you can observe a first stable version of my CSV-parser which receives 2 arguments:

**I)** A path to your csv-file
**II)** `all` to parse and render the whole file 
         or
        `x` to parse and render the necessary string,
        where `x` is the integer number of this particular string

I nested this build with two real csv-files for testing.
This Github repository also includes `releases` menu, where I have been adding binaries of the app.

##Preparing to launch

You can build bin-file by yourself. Start the terminal in this folder
and type `go build csvParser.go` to get the bin-file.

Or you can get the binary file from `releases` menu.

##Running the app

For example, to run the app on an one of the included csv-files type 
`./csvParser ./theCSV.csv all` to parse the whole file.

##Feedback
Please enjoy and send your commentaries to `almanaxstories@gmail.com`