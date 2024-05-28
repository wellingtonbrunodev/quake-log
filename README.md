# quake-log

This repo will store the code solution for the Quake Log File challenge in the Cloudwalk hiring process.


Built on top of Go 1.22.3, to run it you need to have go installed on your machine.


To run this program just run the following:

```
cd cmd

go run cmd/main.go
```


The program will execute, process the file containing the log and print the result to the terminal.

For convenience this program also saves the content above to a file inside the folder pkg/output_files/output.json

If you want to exchange the input log file for a new one just create it in the following path (replacing the old one): pkg/input_files/qgames.log
