# Ensure 'vscode-tmgrammar-test' is installed before running these tests.

tests:
	vscode-tmgrammar-test ./test/semantic_tokens.go
	vscode-tmgrammar-test ./test/newgo_fork_syntax.go
snap:
	vscode-tmgrammar-snap ./test/semantic_tokens.go -u
	vscode-tmgrammar-snap ./test/newgo_fork_syntax.go -u
new-snap:
	vscode-tmgrammar-snap ./test/semantic_tokens.go
	vscode-tmgrammar-snap ./test/newgo_fork_syntax.go
stress:
	vscode-tmgrammar-test ./test/stress.go
	vscode-tmgrammar-snap ./test/stress.go -u
ready: tests snap
