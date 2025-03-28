#!/bin/zsh

export GORELEASER_CURRENT_TAG=$(git-semver -target patch  -prefix v)
unset GORELEASER_PREVIOUS_TAG
export GITHUB_TOKEN=$(gh auth token)
git add .
git commit -m "release commit"
git tag -a ${GORELEASER_CURRENT_TAG} -m "release tag"
git push --tags
goreleaser release --clean
