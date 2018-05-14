# Contributing

We welcome contributions to Myapp of any kind.
By participating to this project, you agree to abide our [code of conduct](http://contributor-covenant.org/version/1/4/code-of-conduct.html).

## Prerequisite tools are

* [Git](http://git-scm.com/)
* [Go 1.10+](http://golang.org/doc/install)

## Fetch the sources from GitHub

Here's a simple walk-through for getting started:

1. Clone `myapp` from source into `$GOPATH`:

    ```bash
    go get -u -v -d github.com/deild/myapp
    cd $GOPATH/src/github.com/deild/myapp
    ```

1. Install the build dependencies:

    Note that Myapp uses [Go Dep](https://github.com/golang/dep) to vendor dependencies, rather than a simple `go get`. We don't commit the vendored packages themselves to the Myapp git repository.

    The simplest way is to use [mage](https://magefile.org/) (a Make alternative for Go projects.).

    ```bash
    mage vendor
    ```

1. A good way of making sure everything is all right is running the test suite:

    ```bash
    mage test
    ```

## Submit a pull request

1. Create a new branch for your changes:

    Please use a short and descriptive branch name, e.g. **NOT** "patch-1". It's very common but creates a naming conflict each time when a submission is pulled for a review.

    ```bash
    git checkout -b iss1234
    ```

1. After making your changes, commit them to your new branch follow the [Git Commit Message Guidelines](#git-commit-message-guidelines):

    ```bash
    git commit -a -v
    ```

1. [Fork](https://help.github.com/articles/fork-a-repo/) Myapp in GitHub.

1. Add your fork as a new remote (the remote name, "fork" in this example, is arbitrary):

    ```bash
    git remote add fork git://github.com/USERNAME/hugo.git
    ```

1. Push the changes to your new remote:

    ```bash
    git push --set-upstream fork iss1234
    ```

1. Ensure that `mage ci` succeeds. [Travis CI](https://travis-ci.org/deild/myapp) will fail the build if `mage ci` fails.

1. You're now ready to submit a PR based upon the new branch in your forked repository.

## Updating the Myapp Sources

If you want to stay in sync with the Myapp repository, you can easily pull down the source changes, but you'll need to keep the vendored packages up-to-date as well.

```bash
git pull
mage vendor
```

## Git Commit Message Guidelines

This [blog article](http://chris.beams.io/posts/git-commit/) is a good resource for learning how to write good commit messages,
the most important part being that each commit message should have a title/subject in imperative mood starting with a capital letter and no trailing period:
*"Return error on wrong use of the Paginator"*, **NOT** *"returning some error."*

Sometimes it makes sense to prefix the commit message with with the type. Choose one of the following: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `chore`, `revert`, `bump`

Also, if your commit references one or more GitHub issues, always end your commit message body with *See #1234* or *Fixes #1234*.
Replace *1234* with the GitHub issue ID. The last example will close the issue when the commit is merged into *master*.

Here’s a useful heuristic for writing better commit messages. Set your commit message template to:

```text
# <type>: If applied, this commit will <subject>
# |<----  Using a Maximum Of 50 Characters  ---->|

# Explain what and why not how this change is being made
# |<----   Try To Limit Each Line to a Maximum Of 72 Characters   ---->|

# Provide any references to tickets, articles or other resources
# Example: See #123, Fixes #456

# --- COMMIT END ---
# Type can be
#    feat     (new feature)
#    fix      (bug fix)
#    refactor (refactoring production code)
#    style    (formatting, missing semi colons, etc; no code change)
#    docs     (changes to documentation)
#    test     (adding or refactoring tests; no production code change)
#    chore    (updating grunt tasks etc; no production code change)
#    revert   (reverts a previous commit)
#    perf     (changes that improves performance)
#    bump     (increase the version of something)
# --------------------
# Remember to
#    Capitalize the subject line
#    Use the imperative mood in the subject line
#    Do not end the subject line with a period
#    Separate subject from body with a blank line
#    Use the body to explain what and why vs. how
```

To do this in Git, save the above content in a file (eg ~/.git_commit_msg) and run:

```bash
git config --global commit.template ~/.git_commit_msg
```
