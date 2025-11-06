# Go GitHub Actions CI

This project demonstrates a Continuous Integration (CI) pipeline for a Go application using GitHub Actions. It showcases a standard workflow that includes installing dependencies, running linters, executing tests, and building the application.

## About The Project

This repository contains a simple REST API built with the Gin framework in Go. The primary goal is to provide a clear example of how to configure a CI workflow for a Go project.

## CI Workflow (`.github/workflows/ci.yml`)

The CI pipeline is configured to run on every `push` and `pull_request` to the `main` branch. It consists of a single job, `build`, which performs the following steps:

1.  **Checkout Repository**: The first step checks out the repository's code so the workflow can access it.
    ```yaml
    - name: Checkout repository
      uses: actions/checkout@v4
    ```

2.  **Set up Go**: It installs and configures the specified version of the Go programming language.
    ```yaml
    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.24.x'
    ```

3.  **Install Dependencies**: This step runs `go mod tidy` to ensure the `go.mod` and `go.sum` files are consistent and downloads the required dependencies.
    ```yaml
    - name: Install dependencies
      run: go mod tidy
    ```

4.  **Run Linter**: The workflow is set up to use `golangci-lint` to analyze the source code for style issues and potential bugs. This helps maintain code quality. (Note: This step is present but commented out in the current `ci.yml`).
    ```yaml
    # - name: Run linter
    #   uses: golangci/golangci-lint-action@v6
    #   with:
    #     version: v1.59
    ```

5.  **Run Tests**: It executes the project's unit tests using `go test`. The `-v ./...` flags ensure that tests in all subdirectories are run with verbose output.
    ```yaml
    - name: Run tests
      run: go test -v ./...
    ```

6.  **Build Application**: Finally, the workflow builds the application to verify that it compiles successfully. The `-v` flag provides verbose output.
    ```yaml
    - name: Build application
      run: go build -v -o go-github-actions-ci .
    ```