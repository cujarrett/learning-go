# learning-go

A small set of hacks to learn Go.

## Rules

- **Never run `git commit`, `git push`, or any git command that writes to or modifies repository history or remotes.** If a task requires committing or pushing, stop and tell the user to run the git command manually.

### Pre-commit safety check

Before telling the user to commit, always run `/pre-commit-review`. It checks for secrets, sensitive identifiers, PII, credential templates, and cluster safety, and returns explicit verdicts on whether the changes are safe for a public repo. Once it confirms the changes are safe, offer the user a suggested commit message — do not run `git commit` yourself.
