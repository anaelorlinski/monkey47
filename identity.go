package monkey47

type Identity struct {
	component string
	version   string
}

var identity Identity

func SetComponent(component string) {
	identity.component = component
}

func Component() string {
	return identity.component
}

func SetVersion(version string) {
	identity.version = version
}

func Version() string {
	return identity.version
}
