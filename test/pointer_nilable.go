// SYNTAX TEST "source.go" "Pointer nilable return type"

package test

type nsLockMap struct{}

func newNSLock(isDistErasure bool) *nsLockMap? {
	return nil
}

func find(path string) *dataUsageEntry? {
	return nil
}
