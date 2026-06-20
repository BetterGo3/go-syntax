// SYNTAX TEST "source.go" "Go fork syntax (enum, lambdas, result types, default args)"

// Go fork syntax highlighting samples (see go/doc/new_features/).

package fork

enum reportSection {
	TopPosts
	SubredditBreakdown
	CommentDepth
}

enum Color {
	Red { r, g, b int }
	Green
	Blue(int)
}

enum Message {
	Write {
		text string,
		bytes int
	}
}

enum Option[T any] {
	None
	Some(T)
}

func fetchPosts() []Post! {
	return nil
}

func fetchOne() Post! {
	return Post{}
}

func readTitle() string! {
	return fetchOne()!.Title
}

func greet(name string, greeting string = "hello") string {
	return greeting + " " + name
}

func open(path string, mode Mode = Read) Mode {
	return mode
}

func nullable(x int?) int {
	return x ?? 0
}

// Real code needs import "linq"; omitted here so this grammar sample stays import-free.
func lambdaChain(posts []Post) {
	_ = posts.Where(p => p.Score >= 10)
		.Select(p => p.Title)
}

func switchExpr(v Color) string {
	return switch v {
	case Red {}:
		"red"
	case Green:
		"green"
	case Blue(_):
		"blue"
	default:
		"other"
	}
}

func ifExpr(ok bool) int {
	return if ok { 1 } else { 0 }
}

func nullCond(obj *Person) string? {
	return obj?.Name
}

type Post struct {
	Title string
	Score int
}

struct Person {
	Name  string
	Title string
}

interface Named {
	Name() string
}
