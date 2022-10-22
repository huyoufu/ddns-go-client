package util

type ServiceAction int

const (
	Init ServiceAction = iota
	Install
	Uninstall
	Run
)

func DetectServiceAction(actionStr string) ServiceAction {
	if actionStr == "init" {
		return Init
	}
	if actionStr == "install" {
		return Install
	}
	if actionStr == "uninstall" {
		return Uninstall
	}
	return Run

}
