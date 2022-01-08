# TerratestWithTerraform

Terratest is a Go library developed at Gruntwork, that makes it easier to write automated tests for our infrastructure code. It provides a variety of helper functions and patterns for common infrastructure testing tasks but here we will be discussing about Testing Terraform code.

# To Run this application
* git clone https://github.com/imnitin28/terra-form-test-techhub.git  <br />
* cd test  <br />
* go mod init "<MODULE_NAME>"  <br />
* **MODULE_NAME would be github.com/<YOUR_USERNAME>/<YOUR_REPO_NAME>**  <br />
* go mod init github.com/imnitin28/terra-form-test-techhub  <br />
* go run

--------------------------------------------------------------------------------------------------------------------------------------------------------------------

go mod init "<MODULE_NAME>" would create go.mod file into test folder.  <br />
* The go.mod file is the root of dependency management in GoLang. 
* All the modules which are needed or to be used in the project are maintained here in go.mod file.
* It creates entry for all the packages we are going to use/import in our project.
* It reduces effort for getting each dependencies manually.

On running **go test** for the first time you would get go.sum file created.
* go.sum file is created when **go test** or **go build** is executed for the first time.
* It installs all the packages with specific version(latest)
* we do not need to edit or modify this file.

