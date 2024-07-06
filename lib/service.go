package lib

type Store struct {
	Scripts []App
}

var store = Store{
	Scripts: []App{},
}

func GetScripts() *[]App {
	return &store.Scripts
}

func (script *Script) AddScript() {
	script.Id = len(store.Scripts) + 1
	store.Scripts = append(store.Scripts, *script)
}
