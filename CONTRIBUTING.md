# Contributing Guidelines

Thank you for your interest in contributing to our project. Whether it's a bug report, new feature, correction, or additional
documentation, we greatly value feedback and contributions from our community.

Please read through this document before submitting any issues or pull requests to ensure we have all the necessary
information to effectively respond to your bug report or contribution.

## Reporting Bugs, Features, and Enhancements

We welcome you to use the GitHub issue tracker to report bugs or suggest features and enhancements.

When filing an issue, please check existing open, or recently closed, issues to make sure someone else hasn't already
reported the issue.

Please try to include as much information as you can. Details like these are incredibly useful:

* A reproducible test case or series of steps.
* Any modifications you've made relevant to the bug.
* Anything unusual about your environment or deployment.

## Contributing via Pull Requests

Contributions via pull requests are appreciated. Before sending us a pull request, please ensure that:

1. You [open a discussion](https://github.com/MichaelCade/90DaysOfDevOps/discussions) to discuss any significant work with the maintainer(s).
2. You open an issue and link your pull request to the issue for context.
3. You are working against the latest source on the `main` branch.
4. You check existing open, and recently merged, pull requests to make sure someone else hasn't already addressed the problem.

To send us a pull request, please:

1. Fork the repository.
2. Modify the source; please focus on the **specific** change you are contributing.
3. Ensure local tests pass.
4. Updated the documentation, if required.
4. Commit to your fork [using a clear commit messages](http://chris.beams.io/posts/git-commit/). We ask you to please use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
5. Send us a pull request, answering any default questions in the pull request.
6. Pay attention to any automated failures reported in the pull request, and stay involved in the conversation.

GitHub provides additional document on [forking a repository](https://help.github.com/articles/fork-a-repo/) and
[creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

### Contributor Flow

This is a rough outline of what a contributor's workflow looks like:

- Create a topic branch from where you want to base your work.
- Make commits of logical units.
- Make sure your commit messages are [in the proper format](http://chris.beams.io/posts/git-commit/).
- Push your changes to a topic branch in your fork of the repository.
- Submit a pull request.

Example:

``` shell
git remote add upstream https://github.com/vmware-samples/packer-examples-for-vsphere.git
git checkout -b my-new-feature main
git commit -s -a
git push origin my-new-feature
```

### Staying In Sync With Upstream

When your branch gets out of sync with the  90DaysOfDevOps/main branch, use the following to update:

``` shell
git checkout my-new-feature
git fetch -a
git pull --rebase upstream main
git push --force-with-lease origin my-new-feature
```

### Updating Pull Requests

If your pull request fails to pass or needs changes based on code review, you'll most likely want to squash these changes into
existing commits.

If your pull request contains a single commit or your changes are related to the most recent commit, you can simply amend the commit.

``` shell
git add .
git commit --amend
git push --force-with-lease origin my-new-feature
```

If you need to squash changes into an earlier commit, you can use:

``` shell
git add .
git commit --fixup <commit>
git rebase -i --autosquash main
git push --force-with-lease origin my-new-feature
```

Be sure to add a comment to the pull request indicating your new changes are ready to review, as GitHub does not generate a notification when you `git push`.

### Formatting Commit Messages

We follow the conventions on [How to Write a Git Commit Message](http://chris.beams.io/posts/git-commit/).

Be sure to include any related GitHub issue references in the commit message.

See [GFM syntax](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown) for referencing issues and commits.

## Reporting Bugs and Creating Issues

When opening a new issue, try to roughly follow the commit message format conventions above.

## Finding Contributions to Work On

Looking at the existing issues is a great way to find something to contribute on. If you have an idea you'd like to discuss, [open a discussion](https://github.com/MichaelCade/90DaysOfDevOps/discussions).

## License

Shield: [![CC BY-NC-SA 4.0][cc-by-nc-sa-shield]][cc-by-nc-sa]

This work is licensed under a
[Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License][cc-by-nc-sa].

[![CC BY-NC-SA 4.0][cc-by-nc-sa-image]][cc-by-nc-sa]

[cc-by-nc-sa]: http://creativecommons.org/licenses/by-nc-sa/4.0/
[cc-by-nc-sa-image]: https://licensebuttons.net/l/by-nc-sa/4.0/88x31.png
[cc-by-nc-sa-shield]: https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-lightgrey.svg
