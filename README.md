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