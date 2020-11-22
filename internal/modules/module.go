package modules

type Module interface {
	Name() string
	// Installation script for the module
	Install()
	// Uninstallation script for the module
	Uninstall()
}
