# API Polling Service

This repository contains a simple Go program that demonstrates the usage of the GitHub API to automate issue management. The program checks for open issues with a specific label in a given repository and automatically closes them.

## Overview

The program does the following:

1. Authenticates with the GitHub API using a provided access token.
2. Periodically polls the GitHub API to retrieve open issues in the specified repository.
3. Checks if the open issues have the desired label.
4. Closes the issues with the desired label.

## Showcase

This project serves as a showcase of how a service can be built to interact between two APIs, in this case, using GitHub's API to manage issues. GitHub issues were chosen as the example because they provide a simple and accessible way to demonstrate the concept.

By using this code as a base, you can expand upon it to create a more sophisticated service that manages issues or other resources across multiple APIs, or even create your own API to interact with GitHub or other services.

## Setup

To set up and run the program, follow these steps:

1. Install the Go programming language if you haven't already, by following the instructions here: https://golang.org/doc/install

2. Clone this repository:
``` 
git clone https://github.com/Rlyehan/api_polling_service.git
cd api_polling_service 
```

3. Install the required dependencies:
```
go get -u github.com/google/go-github/v38/github
go get -u golang.org/x/oauth2
```

4. Set up a GitHub personal access token with `repo` scope, by following the instructions here: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token

5. Export the GitHub access token as an environment variable:
```
export GITHUB_ACCESS_TOKEN=your_token_here
```

6. Update the `const` section in the code with your desired GitHub username, repository name, and label:

```go
const (
    owner           = "your_github_username"
    repo            = "your_repository_name"
    desiredLabel    = "your_desired_label"
    pollingInterval = 30 * time.Second
)
```

7. Run the Program:
```
go run main.go
```

## Testing the Program

To test if the program is working correctly, follow these steps:

1. Make sure you have set up the program as described in the [Setup](#setup) section.

2. Create an issue in the specified GitHub repository and add the desired label to it. You can do this from the GitHub web interface by navigating to the "Issues" tab in your repository and clicking "New Issue." After entering a title and description for the issue, click on the "Labels" dropdown on the right and select your desired label.

3. Run the program

4. Observe the program's output. You should see a message indicating that the issue with the desired label has been closed:
```
Closing issue #<issue_number> with label '<desired_label>'
```

5. Verify that the issue is closed by checking the "Issues" tab in your GitHub repository. The issue you created should now have a "Closed" status.

By following these steps, you can confirm that the program is correctly polling the GitHub API for open issues with the desired label and closing them.

