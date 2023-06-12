# General Anhgelus contribution guidelines

Thank you for investing your time in contributing to our project!

First of all, don't forget to read our [Code of Conduct](./CODE_OF_CONDUCT.md).

In this guide you will get an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

## Issues

In this section, we'll talk about the "issue" part.

### Create a new issue

If you spot a problem, [search if an issue is already open](https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests#search-by-the-title-body-or-comments).
If a related issue doesn't exist, you can open a new issue using the "Issues" tab.

When you create a new issue, use our issue template and fill every required informations.

### Solve an issue

If you want to solve an issue, search an existing issue and make a PR to fix it.

## Pull Requests (PR)

In this section, we'll talk about the "PR" part.

### Disclaimer

In this section, we'll not talk on "how to use git or github", we'll just talk about how we works with PR in this project.

### Create a PR

Before creating a PR, search if a related PR is already open.

When you create a new PR, use our PR template and fill every required informations.

### Editing our code

When you are editing our code to make a new PR, you must follow our standards.

#### Commit

One commit is one changes!

Commit namings:
- Your name is composed in 3 parts: `prefix(section): details`
- `prefix` describe the type of changes.
- `section` describe the part affected by your commit (flatcase).
- `details` describe your changes (flatcase).

Prefixes:
> `build` - affects the build part (gradle, npm, etc)
>
> `ci` - affects the continuous integration (git, github actions, etc)
>
> `feat` - add a new feature
>
> `fix` - fix a bug
>
> `perf` - improves performance
>
> `refactor` - changes that brings no new functionality or performance improvements
>
> `style` - changes that do not alter function or semantics (indentation, formatting, etc.)
>
> `docs` - changes related to the documentations
>
> `test` - changes related to the test

The details part must not exceed 70 characters in length.

When your commit contains a breaking change, use this format: `prefix(part)!: details`

#### Branch

Understand the branch system is very important.

The default branch is `main` and no one can push changes directly to this branch.
We must create a new branch and merge it through a PR if we want to modify the `main` branch.

Branch namings:
- Your name is composed in 2 parts: `prefix/details`
- `prefix` describe the type of changes. These are the same as the commits.
- `details` describe your changes (kebab case).

#### Tag

Basically, you are not autorized to add tag to your changes.

We follow the [semver](https://semver.org/).
So a tag is composed by: `M.m.p` (e.g. 1.2.3).
- `M` is the major version.
- `m` is the minor version.
- `p` is the patch version.

Major is increased when a new breaking changes is introduced.
Minor is increased when a new features is introduced.
Patch is increased when other changes is introduced.

### Reviewing PR

When you make your PR finished, we'll review your PR.

If everything is good, we'll merge it and close the PR.

If something is not good, we'll start a conversation and you must fix this before merging.
