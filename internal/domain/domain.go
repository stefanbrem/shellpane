package domain

type ShellpaneConfig struct {
	Views 		[]ViewSpec
}

type ViewSpec struct {
	Name        string
	Command     string
	Env         []EnvSpec
}

type ViewOutput struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

type EnvSpec struct {
	Name string
}
