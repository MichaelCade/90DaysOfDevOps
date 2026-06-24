Go es una herramienta para gestionar el código fuente de Go.

Sintaxis:

	go <command> [arguments]

Los comandos son:

	bug         Iniciar un informe de error
	build       Compilar paquetes y dependencias
	clean       Eliminar ficheros de objetos y en caché
	doc         Mostrar la documentación del paquete o símbolo
	env         Imprimir información del entorno de Go
	fix         Actualizar paquetes para que usen las nuevas APIs
	fmt         gofmt (reformatear) las fuentes de los paquetes
	generate    Generar ficheros Go procesando el código fuente
	get         Añadir dependencias al módulo actual e instalarlas
	install     Compilar e instalar paquetes y dependencias
	list        Listar paquetes o módulos
	mod         Mantenimiento de módulos
	work        Mantenimiento del workspace
	run         Compilar y ejecutar el programa Go
	test        Test de paquetes
	tool        Ejecutar la herramienta go especificada
	version     Imprimir la versión de Go
	vet         Informar de posibles errores en los paquetes

Use "go help <command>" para obtener más información sobre un comando.

Temas de ayuda adicionales:

	buildconstraint Restricciones de construcción
	buildmode       Modos de construcción
	c               Llamada entre Go y C
	cache           construcción y prueba de caché
	environment     variables de entorno
	filetype        Tipos de ficheros
	go.mod          El fichero go.mod
	gopath          Variable de entorno GOPATH
	gopath-get      Obtener legacy GOPATH
	goproxy         Protocolo modulo proxy
	importpath      Sintaxis path de importación
	modules         Módulos, versión de módulos y más
	module-get      module-aware go get
	module-auth     Autenticación de módulos mediante go.sum
	packages        Listas de paquetes y patrones
	private         Configuración para la descarga de código no público
	testflag        testing flags
	testfunc        testing functions
	vcs             Control de versiones con GOVCS

Use "go help <topic>" para obtener más información sobre ese tema.
