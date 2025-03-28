install-cli:
	brew install gh # https://github.com/cli/cli
	go install github.com/goreleaser/goreleaser/v2@v2.8.1 # https://goreleaser.com/quick-start/
	brew install mdomke/git-semver/git-semver # https://github.com/mdomke/git-semver?tab=readme-ov-file#installation


auth-github:
	gh auth login

release:
	./release.sh