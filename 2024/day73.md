# Day 73 - Introducing the Terraform Test Framework
[![Watch the video](thumbnails/day73.png)](https://www.youtube.com/watch?v=ksYiVW6fEeA)

 In this explanation, you have described a Terraform test setup for deploying an infrastructure and checking its availability. Here's a summary of the steps you have outlined:

1. Create a random number using a custom module to append to the website name for unique deployment.
2. Set up the integration tests in the following order:
   - First run (not really a test): Creates the random integer value, sourced from a module within the `tests/setup` subdirectory.
   - Second run (execute): Deploys the actual configuration using the generated random number as part of the website name.
   - Third run (check site): Checks that the deployed website is up and running by using another custom module to get the response from the website URL and asserting that the status code equals 200.
3. When executing Terraform tests, remember the order of runs matters, regardless of their location in separate files, and lexically when multiple files are involved.
4. Reference the outputs of a module run in the next run by using the `run.<name of the run>.<name of the output>` syntax.
5. Use the implicit `apply` command in Terraform tests or make it explicit for clarity, but remember that the implicit command is assumed if not stated.
6. Provide feedback on the Terraform testing framework to help it improve further.

For more details, you can find the example code in the link provided or in your GitHub repository (terraform-tuesdays). Happy learning and testing!
A presentation on Terraform testing! Let me summarize the key points for you:

**Identity**: The speaker, Ned Bellance, introduces himself as an expert content summarizer and explains that this is not a test, but rather a presentation on Terraform testing.

**Purpose**: The purpose of the presentation is to demonstrate how to use Terraform testing to write unit and integration tests for your Terraform code. The speaker will be walking through an example of setting up infrastructure using Terraform and running tests against it.

**Example**: The speaker shows how to create a simple module that uses a random integer resource to create a number between 101 and 999, and then passes that integer value out as an output. This output can be used in subsequent test runs.

The speaker then demonstrates how to write integration tests using the Terraform testing framework. He creates a run block to deploy the actual configuration, and then uses another run block to check that the website is up and running by asserting that the status code of the data source equals 200.

**Terraform Testing**: The speaker notes that the order of runs matters in Terraform testing, with all runs of the same command type (e.g., "apply") being executed in the order they appear. He also explains how to reference outputs from a module run in the next run using the syntax `run <name> <output>`.

**Conclusion**: The speaker concludes by highlighting the benefits of Terraform testing, including the ability to write tests in the same language as your infrastructure code (Terraform) and the simplicity of setting up tests. He encourages attendees to take Terraform testing for a test drive and provide feedback to the maintainers of Terraform and Open TDD.
