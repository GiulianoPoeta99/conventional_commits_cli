# Conventional Commits CLI

Conventional Commits CLI is an interactive command-line tool designed to help you create standardized commit messages in accordance with the Conventional Commits specification. This tool streamlines the commit process by guiding you through all the necessary elements—from type and scope to detailed descriptions and metadata like reviewers and issue references.

## Features

- **Interactive Prompts:** Step-by-step guidance using interactive prompts to select commit type, add a scope, and enter detailed descriptions.
- **Emoji Integration:** Enhance your commit messages with GitHub emojis. Recommended emojis are suggested based on the type of change.
- **Validation:** Built-in validations ensure that commit descriptions and other inputs meet the necessary criteria.
- **Advanced Metadata:** Optionally add commit body, denote breaking changes, reference issues, and list reviewers.
- **Extensible and Configurable:** Easily extend or configure settings (e.g., storing commonly used scopes) to streamline your commit process.

## Requirements

- Go (version 1.16+ recommended)
- Git, properly installed and configured
- [PromptUI](https://github.com/manifoldco/promptui) (included as a dependency)

## Installation

You can install the CLI tool using Go:

```bash
go install github.com/GiulianoPoeta99/conventional_commits_cli@latest
```

Alternatively, clone the repository and build the project locally:

```bash
git clone https://github.com/GiulianoPoeta99/conventional_commits_cli.git
cd conventional_commits_cli
go build
```

## Usage

Run the CLI in your repository to launch the commit assistant:

```bash
./conventional_commits_cli
```

Follow the interactive prompts to create a commit message that adheres to the Conventional Commits standard.

## Roadmap / TODO

- Full Emoji Integration:
  - [List of GitHub Emojis](https://gist.github.com/rxaviers/7360908)
  - Integrate the complete list of available GitHub emojis to expand your emoji options.
- Emoji Search Functionality:
- Add a search feature within the emoji selection prompt to quickly find the desired emoji.
-Scope Persistence:
  Implement a method for saving and reusing scopes, potentially by storing them in a dedicated configuration folder (e.g., under .config).

## Contributing

Contributions are very welcome! If you have ideas, improvements, or bugs to report, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](./LICENSE).

## Additional Resources

- [Conventional Commits Specification](https://www.conventionalcommits.org/en/v1.0.0/)
