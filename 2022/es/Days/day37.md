## Conociendo a Git

Pido disculpas por los terribles juegos de palabras del título (Título original: *Gitting to know Git*). Seguro que no soy la primera persona que convierte Git en un chiste de padres.

En los dos últimos posts aprendimos acerca de los sistemas de control de versiones, y algunos de los flujos de trabajo fundamentales de git ([Día 35](day35.md)). Luego conseguimos instalar git en nuestro sistema, actualizarlo y configurarlo. También vimos algo de la teoría entre los tipos de sistema de control de versiones Cliente-Servidor y los distribuidos en el [Día 36](day36.md).

Ahora vamos a repasar algunos de los comandos y casos de uso  habituales con git.

### Git ayuda con git

Habrá momentos en los que simplemente no recuerdes o no sepas el comando que necesitas para hacer algo con git. Necesitarás ayuda. Google o cualquier motor de búsqueda será probablemente tu primer puerto de escala cuando busques ayuda. (También las IAs)

El siguiente lugar, sobretodo a modo de confirmación, va a ser el sitio oficial de git y la documentación: [git-scm.com/docs](http://git-scm.com/docs). Aquí no sólo encontrarás una referencia sólida de todos los comandos disponibles, sino también un montón de recursos diferentes.

![](Images/Day37_Git1.png)

También podemos acceder a esta misma documentación desde la terminal, lo cual es muy útil si trabajas sin conexión. Si elegimos el comando `git add` por ejemplo podemos ejecutar `git add --help` y veremos el manual concreto del subcomando.

![](Images/Day37_Git2.png)

También podemos usar en la shell `git add -h` que nos va a dar un resumen de las opciones que tenemos disponibles.

![](Images/Day37_Git3.png)

### Mitos en torno a Git

"Git no tiene control de acceso" - Puedes autorizar a quien quieras a mantener el código fuente.

"Git es demasiado pesado" - Git puede proporcionar repositorios poco profundos, lo que significa una cantidad reducida de historial si tienes proyectos grandes.

### Deficiencias reales

No es ideal para archivos binarios. Es ideal para código fuente, pero no para archivos ejecutables o vídeos, por ejemplo.

Git no es fácil de aprender, el hecho de que tengamos que pasar tiempo hablando de comandos y funciones de la herramienta es probablemente una señal de ello. Sin embargo, en general, git es difícil de aprender pero fácil de usar.

### El ecosistema git

Quiero cubrir brevemente el ecosistema alrededor de git pero no profundizar en estas áreas, pero creo que es importante señalarlas aquí a un alto nivel.

Casi todas las herramientas de desarrollo modernas soportan Git.

- **Herramientas para desarrolladores** - Ya hemos mencionado visual studio code pero encontrarás plugins de git e integraciones en sublime text y otros editores de texto e IDEs.
- **Herramientas de equipo** - También se han mencionado herramientas como Jenkins desde un punto de vista de CI/CD, Slack desde un marco de mensajería y Jira para la gestión de proyectos y seguimiento de incidencias.
- **Proveedores de nube** - Todos los grandes proveedores de nube soportan git: Microsoft Azure, Amazon AWS y Google Cloud Platform.
- **Servicios basados en Git** - Luego tenemos GitHub, GitLab y BitBucket que cubriremos en más detalle más adelante. Estos servicios se conocen como las redes sociales del código.

### The Git Cheatsheet

No hemos cubierto la mayoría de comandos de git, ya que tiene un amplio abanico de posibilidades, pero después de haber mirado algunas hojas de trucos disponibles en línea vamos a documentar algunos de ellos y su propósito. No necesitamos recordarlos todos, ya que a través de la práctica y el uso aprenderás al menos mucho más de git.

Estos comandos se han sacado de [atlassian](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet), escribirlos y leer la descripción es una buena forma de conocer su uso, así que manos a la obra en las tareas diarias:

- [Conceptos básicos de Git](#conceptos-básicos-de-git)
- [Deshacer cambios Git](#deshacer-cambios-git)
- [Historial de reescritura de Git](#historial-de-reescritura-de-git)
- [Git Branches](#git-branches)
- [Repositorios remotos Git](#repositorios-remotos-git)
- [Git Diff](#git-diff)
- [Git Config](#git-config)
- [Git Rebase](#git-rebase)
- [Git Pull](#git-pull)
- [Git Reset](#git-reset)
- [Git Push](#git-push)

### Conceptos básicos de Git

| Comando       | Ejemplo                     | Descripción                                                                                                           |
| ------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| git init      | `git init <directory>`      | Crea un repositorio git vacío en el directorio especificado.                                                          |
| git clone     | `git clone <repo>`          | Clona el repositorio ubicado en <repo> en la máquina local.                                                           |
| git config    | `git config user.name`      | Definir opciones que se usará para todos los commits según el entorno `system`, `global` o `local`.                   |
| git add       | `git add <directory>`       | Prepara los cambios en <directorio> para la próxima confirmación. También podemos añadir <ficheros> y <.> para todo.  |
| git commit -m | `git commit -m "<message>"` | Confirmar la instantánea preparada, usando <mensaje> para detallar lo que se está confirmando.                        |
| git status    | `git status`                | Lista los archivos que están organizados, no organizados y sin seguimiento.                                           |
| git log       | `git log`                   | Muestra todo el historial de confirmaciones usando el formato por defecto. Hay opciones adicionales con este comando. |
| git diff      | `git diff`                  | Muestra los cambios no organizados entre tu índice y el directorio de trabajo.                                        |

### Deshacer cambios Git

| Comando    | Ejemplo               | Descripción                                                                                                        |
| ---------- | --------------------- | ------------------------------------------------------------------------------------------------------------------ |
| git revert | `git revert <commit>` | Crear un nuevo commit que deshaga todos los cambios realizados en <commit> y aplicarlo a la rama actual.           |
| git reset  | `git reset <file>`    | Elimina <fichero> del área de preparación dejando el directorio de trabajo sin cambios.                            |
| git clean  | `git clean -n`        | Muestra qué archivos se eliminarán del directorio de trabajo. Usa `-f` en lugar de `-n` para ejecutar la limpieza. |

### Historial de reescritura de Git

| Comando    | Ejemplo              | Descripción                                                                                                                                                 |
| ---------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| git commit | `git commit --amend` | Reemplazar el último commit con los cambios en staged y el último commit combinados. Editas el mensaje del último commit.                                   |
| git rebase | `git rebase <base>`  | Rebase la rama actual en <base> que puede ser un commit ID, un nombre de rama, una etiqueta o una referencia relativa a HEAD.                               |
| git reflog | `git reflog`         | Muestra registro de cambios en el HEAD del repositorio local. Con --relative-date muestra información de fechas y --all para mostrar todas las referencias. |

### Git Branches

| Comando      | Ejemplo                    | Descripción                                                                                                            |
| ------------ | -------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| git branch   | `git branch`               | Lista todas las ramas de tu repositorio. Añade un argumento <branch> para crear una nueva rama con el nombre <branch>. |
| git checkout | `git checkout -b <branch>` | Crea y comprueba una nueva rama con el nombre <branch>. Elimina la opción -b para obtener una rama existent            |
| git merge    | `git merge <branch>`       | Fusiona <branch> en la rama actual.                                                                                    |

### Repositorios remotos Git

| Comando        | Ejemplo                       | Descripción                                                                                                                                 |
| -------------- | ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| git remote add | `git remote add <name> <url>` | Crea una nueva conexión a un repositorio remoto. Después de añadir un remoto, puedes usar <nombre> como atajo para <url> en otros comandos. |
| git fetch      | `git fetch <remote> <branch>` | Obtiene un <branch> específico del repositorio. Deja <branch> para obtener todas las referencias remotas.                                   |
| git pull       | `git pull <remote>`           | Obtiene la copia remota especificada de la rama actual e inmediatamente la fusiona con la copia local.                                      |
| git push       | `git push <remote> <branch>`  | Empuja la rama a <remote>, junto con los commits y objetos necesarios. Crea la rama con el nombre en el repositorio remoto si no existe.    |

### Git Diff

| Comando           | Ejemplo             | Descripción                                                              |
| ----------------- | ------------------- | ------------------------------------------------------------------------ |
| git diff HEAD     | `git diff HEAD`     | Muestra la diferencia entre el directorio de trabajo y el último commit. |
| git diff --cached | `git diff --cached` | Muestra la diferencia entre los cambios realizados y el último commit.   |

### Git Config

| Comando                                              | Ejemplo                                                | Descripción                                                                                                                             |
| ---------------------------------------------------- | ------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------- |
| git config --global user.name <name>                 | `git config --global user.name <name>`                 | Define el nombre de autor que se usará para todos los commits del usuario actual.                                                       |
| git config --global user.email <email>               | `git config --global user.email <email>`               | Define el email del autor que se usará para todas las confirmaciones del usuario actual.                                                |
| git config --global alias <alias-name> <git-command> | `git config --global alias <alias-name> <git-command>` | Crear un acceso directo para un comando git.                                                                                            |
| git config --system core.editor <editor>             | `git config --system core.editor <editor>`             | Establece el editor de texto que usarán todos los usuarios de la máquina. <editor> arg debe ser el comando que lanza el editor deseado. |
| git config --global --edit                           | `git config --global --edit `                          | Abrir el archivo de configuración global en un editor de texto para su edición manual.                                                  |

### Git Rebase

| Comando              | Ejemplo                | Descripción                                                                                                                                              |
| -------------------- | ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| git rebase -i <base> | `git rebase -i <base>` | Rebase interactivamente la rama actual en <base>. Inicia el editor para introducir comandos sobre cómo se transferirá cada confirmación a la nueva base. |

### Git Pull

| Comando                    | Ejemplo                      | Descripción                                                                                                                                    |
| -------------------------- | ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| git pull --rebase <remote> | `git pull --rebase <remote>` | Obtiene la copia remota de la rama actual y la vuelve a basar en la copia local. Utiliza git rebase en lugar de merge para integrar las ramas. |

### Git Reset

| Comando                   | Ejemplo                     | Descripción                                                                                                                                                                                         |
| ------------------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| git reset                 | `git reset `                | Restablece el área de preparación para que coincida con la confirmación más reciente, pero deja el directorio de trabajo sin cambios.                                                               |
| git reset --hard          | `git reset --hard`          | Restablece el área de preparación y el directorio de trabajo para que coincidan con la confirmación más reciente y sobreescribe todos los cambios en el directorio de trabajo.                      |
| git reset <commit>        | `git reset <commit>`        | Mueve la punta de la rama actual hacia atrás hasta <commit>, resetea el área de preparación para que coincida, pero deja el directorio de trabajo solo                                              |
| git reset --hard <commit> | `git reset --hard <commit>` | Igual que el anterior, pero reinicia tanto el área de preparación como el directorio de trabajo para que coincidan. Borra los cambios no comprometidos, y todos los commits posteriores a <commit>. |

### Git Push

| Comando                   | Ejemplo                     | Descripción                                                                                                                                                         |
| ------------------------- | --------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| git push <remote> --force | `git push <remote> --force` | Fuerza el git push incluso si resulta en una fusión no rápida. No uses la opción --force a menos que estés seguro de lo que haces.                                  |
| git push <remote> --all   | `git push <remote> --all`   | Envía todas tus ramas locales a la remota especificada.                                                                                                             |
| git push <remote> --tags  | `git push <remote> --tags`  | Las etiquetas no se envían automáticamente cuando envías una rama o usas la opción --all. La opción --tags envía todas tus etiquetas locales al repositorio remoto. |

## Recursos

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)
- [En español] [Comandos Git](https://gitea.vergaracarmona.es/man-linux/comandos-git)
- [En español][Apuntes Curso de Git](https://vergaracarmona.es/wp-content/uploads/2022/10/Curso-git_vergaracarmona.es_.pdf) de [Juan Carlos Rubio](https://www.linkedin.com/in/juan-carlos-rubio-pineda/Curso-git_vergaracarmona-es).
- [En español] En los [apuntes](https://vergaracarmona.es/apuntes/) del traductor:
  - ["Instalar git en ubuntu"](https://vergaracarmona.es/instalar-git-en-ubuntu/)
  - ["Comandos de git"](https://vergaracarmona.es/comandos-de-git/)
  - ["Estrategias de fusión en git: Ship / Show / Ask"](https://vergaracarmona.es/estrategias-bifurcacion-git-ship-show-ask/)
  - ["Resolver conflictos en Git. Merge, Squash, Rebase o Pull"](https://vergaracarmona.es/merge-squash-rebase-pull/)
  - ["Borrar commits de git: reset, rebase y cherry-pick"](https://vergaracarmona.es/reset-rebase-cherry-pick/)

Nos vemos en el [Día 38](day38.md)
