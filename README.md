Prerequisites:

  * [tfenv](https://github.com/tfutils/tfenv.git )
  * [tgenv](https://github.com/taosmountain/tgenv.git)
    * NOTE: you cannot use the original, unmaintained implementation because it adds JSON incompatible lines to `terragrunt output` calls
Based on:

https://github.com/cloudposse/terraform-example-module

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
1. Add buildspec, .terraform-version
1. Create Codebuild project (manually for now)
    1. Using GitHub personal access token to authenticate
    1. Using Ubuntu 5.0 for golang 1.15 support (as of 9/23/21)
    1. Autocreated Codebuild service role, attached IAM policy to allow it to create/delete the s3 bucket the example module manages
    1. Set a max of 1 concurrent build at a time, though that may be unnecessary
    1. Builds triggered manually for now
