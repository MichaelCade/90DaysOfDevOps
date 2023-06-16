## El panorama general: Gestión de datos

![](Images/Day84_Data1.png)

La gestión de datos no es en absoluto un obstáculo nuevo, aunque sabemos que los datos son más importantes de lo que quizás eran hace unos años. Valiosos y siempre cambiantes, también pueden ser una pesadilla masiva cuando hablamos de automatización e integración continua, pruebas e implementación de versiones de software frecuentes. Los datos persistentes y los servicios de datos subyacentes suelen ser los principales culpables cuando las cosas van mal.

Pero antes de entrar en la Gestión de datos nativa de la nube, debemos subir un nivel. Hemos hablado de muchas plataformas diferentes a lo largo de este desafío. Ya sea física, virtual, en la nube o nativa de la nube, incluido Kubernetes, ninguna de estas plataformas prescinde de la necesidad de gestionar datos.

Sea cual sea nuestro negocio, es muy probable que encontremos una base de datos acechando en algún lugar del entorno, ya sea para el sistema más crítico para el negocio o al menos algún componente del sistema que almacene esos datos persistentes en algún nivel.

### DevOps y datos

Al igual que al comienzo de esta serie donde hablamos de los principios de DevOps, para obtener un mejor proceso en lo que respecta a los datos, es necesario incluir a las personas adecuadas. Esto puede incluir a los administradores de bases de datos (DBAs), pero también a las personas que se preocupan por la copia de seguridad de esos servicios de datos.

En segundo lugar, también necesitamos identificar los diferentes tipos de datos, dominios y límites asociados con nuestros datos. De esta manera, no se trata solo de un enfoque aislado entre administradores de bases de datos, ingenieros de almacenamiento o ingenieros centrados en copias de seguridad. De esta manera, todo el equipo puede determinar la mejor ruta de acción cuando se trata de desarrollar y alojar aplicaciones para el negocio en general, y centrarse en la arquitectura de datos en lugar de considerarla como un añadido posterior.

Ahora, esto puede abarcar muchas áreas diferentes del ciclo de vida de los datos. Podríamos hablar de la ingestión de datos, dónde y cómo se ingresarán los datos en nuestro servicio o aplicación. ¿Cómo accederá el servicio, la aplicación o los usuarios a estos datos? Pero también requiere que entendamos cómo aseguraremos los datos y cómo los protegeremos.

### Gestión de datos 101

La gestión de datos, según el [Cuerpo de conocimientos de gestión de datos](https://www.dama.org/cpages/body-of-knowledge), es "el desarrollo, ejecución y supervisión de planes, políticas, programas y prácticas que controlan, protegen, entregan y mejoran el valor de los datos y los activos de información".

- Los datos son el aspecto más importante de tu negocio: los datos son solo una parte de tu negocio en general. He visto la frase "Los datos son la vida de nuestro negocio" y es muy probablemente cierto. Esto me hizo pensar en que la sangre es bastante importante para el cuerpo, pero por sí sola no es nada, todavía necesitamos los demás aspectos del cuerpo para convertir la sangre en algo más que un líquido.

- La calidad de los datos es más importante que nunca: debemos tratar los datos como un activo empresarial, lo que significa que debemos darles la consideración que necesitan y requieren para trabajar con nuestros principios de automatización y DevOps.

- Acceso a los datos de manera oportuna: nadie tiene paciencia para no tener acceso a los datos adecuados en el momento adecuado para tomar decisiones efectivas. Los datos deben estar disponibles de manera ágil y oportuna, independientemente de la presentación.

- La gestión de datos debe ser un habilitador de DevOps: mencioné anteriormente la agilización, debemos incluir los requisitos de gestión de datos en nuestro ciclo y asegurarnos no solo de la disponibilidad de esos datos, sino también de incluir otras políticas importantes de protección de esos puntos de datos, junto con modelos de recuperación completamente probados.

### DataOps

Tanto DataOps como DevOps aplican las mejores prácticas de desarrollo y operaciones tecnológicas para mejorar la calidad, aumentar la velocidad, reducir las amenazas de seguridad, deleitar a los clientes y proporcionar un trabajo significativo y desafiante para profesionales cualificados. DevOps y DataOps comparten objetivos para acelerar la entrega de productos mediante la automatización de la mayor cantidad de pasos del proceso posible. Para DataOps, el objetivo es una tubería de datos resiliente y conocimientos confiables a partir del análisis de datos.

Algunas de las áreas de nivel superior más comunes en las que se centra DataOps son el aprendizaje automático, el big data y el análisis de datos, incluida la inteligencia artificial.

### La gestión de datos es la gestión de la información

Mi enfoque en esta sección no será adentrarme en el aprendizaje automático o la inteligencia artificial, sino centrarme en la protección de los datos desde un punto de vista de la protección de datos. El título de esta subsección es "La gestión de datos es la gestión de la información" y podemos relacionar esa información con los datos.

Tres áreas clave que debemos considerar en este viaje con los datos son:

- Precisión: asegurarnos de que los datos de producción sean precisos, del mismo modo, debemos asegurarnos de que nuestros datos en forma de copias de seguridad también funcionen y se prueben en caso de recuperación, para asegurarnos de que, si surge un fallo o un problema, podamos volver a funcionar lo más rápido posible.

- Consistencia: si nuestros servicios de datos se extienden a múltiples ubicaciones, debemos asegurarnos de tener consistencia en todas las ubicaciones de datos para obtener datos precisos. Esto también se extiende a la protección de datos cuando se trata de proteger estos servicios de datos, especialmente los servicios de datos, debemos asegurar la consistencia en diferentes niveles para garantizar que obtengamos una copia limpia y completa de esos datos para nuestras copias de seguridad, réplicas, etc.

- Seguridad: el control de acceso y simplemente mantener los datos, en general, es un tema de actualidad en todo el mundo. Asegurarse de que las personas adecuadas tengan acceso a sus datos es primordial. Nuevamente, esto lleva a la protección de datos, donde debemos asegurarnos de que solo el personal necesario tenga acceso a las copias de seguridad y la capacidad de restaurar desde ellas, clonar y proporcionar otras versiones de los datos empresariales.

Mejores datos = Mejores decisiones

### Días de gestión de datos

Durante las próximas 6 sesiones, examinaremos de cerca las bases de datos, la copia de seguridad y recuperación, la recuperación ante desastres y la movilidad de aplicaciones, todo con elementos de demostración y práctica a lo largo del camino.

## Recursos

- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

Nos vemos en el [Día 85](day85.md)
