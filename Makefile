# Ensure 'vscode-tmgrammar-test' is installed before running these tests.

VSCE = node ./node_modules/@vscode/vsce/vsce

tests:
	vscode-tmgrammar-test ./test/semantic_tokens.go
	vscode-tmgrammar-test ./test/newgo_fork_syntax.go
	vscode-tmgrammar-test ./test/pointer_nilable.go
snap:
	vscode-tmgrammar-snap ./test/semantic_tokens.go -u
	vscode-tmgrammar-snap ./test/newgo_fork_syntax.go -u
	vscode-tmgrammar-snap ./test/pointer_nilable.go -u
new-snap:
	vscode-tmgrammar-snap ./test/semantic_tokens.go
	vscode-tmgrammar-snap ./test/newgo_fork_syntax.go
	vscode-tmgrammar-snap ./test/pointer_nilable.go
stress:
	vscode-tmgrammar-test ./test/stress.go
	vscode-tmgrammar-snap ./test/stress.go -u
ready: tests snap

# Package a .vsix for local install (Extensions → Install from VSIX…).
# Requires: npm install
vsix:
	$(VSCE) package --no-dependencies
