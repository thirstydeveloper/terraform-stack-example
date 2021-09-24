# terraform-example-module

Example of testing a terragrunt stack (root module). Derived from https://github.com/cloudposse/terraform-example-module

## Prerequisites

  * Go
  * [tfenv](https://github.com/tfutils/tfenv.git )
  * [tgenv](https://github.com/taosmountain/tgenv.git)
    * NOTE: you cannot use the original, unmaintained implementation because it adds JSON incompatible lines to `terragrunt output` calls

## Approach

Using terratest to test terragrunt root modules differs from traditional terratest testing.

With terratest you typically create a root module under `examples/` and instantiate the module under test there. When testing a
terragrunt stack, the module under test is itself a root module. You could still create a new root module and instantiate the
stack as a child module, but that has several disadvantages:

1. Your terragrunt stack almost certainly declares providers. Providers in child modules are not recommended and have limitations.
2. Your example is not indicative of how someone would actually consume your terragrunt stack.

Instead of creating a new root module, we should consume the stack the same way our users would: with a `terragrunt.hcl`.

It isn't quite that simple though. What if your stack requires a test fixture? Normally you'd put that in the examples root module,
but now we're not creating that.

Terragrunt is helpful here too with its -all commands. We can:

1. Create a separate fixture stack
1. Have our example terragrunt.hcl depend on the fixture stack
1. Use terratest's TgApplyAll and TgDestroyAll to create the fixture first, and our example module second

This works super well!

Next we have to perform assertions. How should we do that? We can use Golang libraries (e.g., AWS SDK). Another option though
is to use Terraform itself to interragate the results using data sources. This cuts down on the amount of golang you need to
write and may be more similar to how consuming stacks might integrate with this stack. If your team isn't strong in Golang,
this may be a more friendly approach to learn.

Since we're using TgApplyAll and TgDestroyAll, we can just add a `consumer` stack that depends on the example stack. The
consumer stack will use data sources to query whatever is needed and expose values using outputs. Our terratest code can then
access those outputs to run assertions.
