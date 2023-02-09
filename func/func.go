package func

func ComposeInt (f func(int) int, g func(int) int) func(int) int {
    return func(x int) int { return f(g(x)) }
}
