package base

type (
	//UseCaseBase Interfaz creada para definir los métodos necesarios para implementar los casos de uso.
	UseCaseBase interface {
		//Execute Método a implementar.
		Execute(command CommandBase, responder ResponderBase)
	}
	//CommandBase Interfaz para lograr definir los parámetros comandos en la interfaz UseCaseBase.
	CommandBase interface {
		//Get Método a implementar, el cual retorna un número de día.
		Get() int
	}

	//ResponderBase Se crea la interfaz simplemente para lograr usarlo como parámetro en la interfaz UseCaseBase.
	ResponderBase interface {
	}
)
