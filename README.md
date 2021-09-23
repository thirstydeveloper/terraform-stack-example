Created this by:

1. Installed go, set GOPATH to ${HOME}/go
1. Create new git repo, cloning
1. Create main .gitignore using gitignore.io for terraform and terragrunt
1. Create directory structure `mkdir -p examples/complete src/example test/src  
1. Create basic module with just outputs.tf and a single output under src/example
1. Create basic module under `examples/complete` that:
    1. main.tf instantiates a module for `../../src/example`
    1. outputs.tf passes through the output from `src/example/outputs.tf`
1. `cd test/src && go mod init github.com/thirstydeveloper/terraform-example-module`
1. `touch Makefile examples_complete_test.go`
1. Add makefile and terratest code
