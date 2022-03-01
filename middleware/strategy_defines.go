package middleware

type Strategy string

const (
	StrategyCheckAuth       Strategy = "check-auth"
	StrategyForwardDirectly Strategy = "forward"
)
