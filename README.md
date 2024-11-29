# GO-LOCATE

If you have worked with linux, chances are you used locate command. Well i find it slow \s. This is an attempt to use go and redis (you can probably see where iam going with this) to make faster locate tool

## PLAN

well its simple, breaking down it into tasks

- Figure out the architecture a little
- code ofc, in go

### Arch
> Well running redis locally in a server might not be a bad idea, but are we gonna do it when we know how to use docker ? no ? redis in docker container it is

> And there is no shortcut to go code for the application / tool itself.

> maybe add a Makefile to deploy a docker container for redis with required port and mount

### USAGE

```bash
# to update the db for particular path
sudo go run main.go --dev testfolder 
```
Example
```bash
> sudo go run main.go 4.txt
testfolder/testfolder2/4.txt
```





## TODO 

well, after little performance test, this does not seem that fast, well atleast we will fix the downsides of the mlocate

1. No Support for Metadata-Based Search

	•	mlocate is limited to searching by file names and does not index file metadata such as file size, modification time, or ownership.
	•	Problem: If your use case requires searching based on metadata (e.g., finding all files larger than a certain size or modified within the last day), mlocate will not provide that functionality.

2. Lack of Granular Search Criteria

	•	mlocate’s search criteria are limited to simple string matches of filenames.
	•	Problem: More complex search needs, like case-insensitive matching, regular expressions, or other advanced search criteria, require additional processing with tools like grep or awk.
