package config

type ProgramLanguage struct {
	InstallPython bool
	PythonConfig  PythonConfig
	InstallNode   bool
	NodeConfig    NodeConfig
	InstallJava   bool
	JavaConfig    JavaConfig
	InstallGo     bool
	GoConfig      GoConfig
}

type DesktopProgram struct {
	InstallVSCode        bool
	VSCodeConfig         VSCodeConfig
	InstallEdge          bool
	EdgeConfig           EdgeConfig
	InstallGit           bool
	GitConfig            GitConfig
	InstallDockerDesktop bool
	DockerDesktopConfig  DockerDesktopConfig
}

type Config struct {
	ProgramLanguage ProgramLanguage
	DesktopProgram  DesktopProgram
}

type PythonConfig struct {
	Version string
}

type NodeConfig struct {
	Version string
}

type JavaConfig struct {
	Version string
}

type GoConfig struct {
	Version string
}

type VSCodeConfig struct {
	Version string
}

type EdgeConfig struct {
	Version string
}

type GitConfig struct {
	Version string
}

type DockerDesktopConfig struct {
	Version string
}
