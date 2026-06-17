// SYNTAX TEST "source.go" "Go fork syntax (enum, lambdas, result types)"

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

func nullable(x int?) int {
	return x ?? 0
}

func lambdaChain(posts []Post) {
	_ = posts.Where(p => p.Score >= 10)
		.Select(p => p.Title)
}

func switchExpr(v Color) string {
	return switchv{case Red:"red"case Green:"green"default:"other"}
}

func ifExpr(ok bool) int {
	return ifok{1}else{0}
}

func nullCond(obj *Person) string? {
	return obj?.Name
}

type Post struct {
	Title string
	Score int
}

type Person struct {
	Name  string
	Title string
}
