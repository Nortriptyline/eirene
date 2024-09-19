package logging

type ISugaredLogger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	// Ajoutez d'autres méthodes si nécessaire
}
