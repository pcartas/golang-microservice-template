# golang-microservice-template

* **Main**: Aquí se inicia el server y se cargan las rutas definidas en routes.go

* **Actions**: Aquí van todos los Handlers de las rutas

* **Routes**: Aquí se definen todas las rutas

* **Config**: Aquí se encuentran los archivos .env de cada ambiente. Durante la ejecucion de la pipeline de acuerdo a la variable ENVIRONMENT (de la pipeline...) prod.env, qa.dev o dev.env pasaran a ser .env y se cargaran en el main. 

    **Aquí NO se colocan secretos y tokens**
    en caso de necesitarlo para el desarrollo local colocarlo directamente en el .env el cual se encuentra en el .gitignore y no hay forma de que se suba al repo. 

    **NO compartir secrets por slack, para eso tenemos lastpass o onepassword**

    Si estas usando ECS recordá que las variables definidas en la Task Definition **siempre** cobraran prioridad sobre las definidas aquí: por ejemplo PORT aqui lo definimos como 8080, pero si en la task definition dice PORT=80, entonces docker tomará esa variable y no la del .env.


## Ejecutar el proyecto en local
* Asegurate de tener una version de go compatible

* Copiar config/dev.env a .env

* Abrir una terminal en la base del proyecto y ejecutar:
    ```
    go get
    ```

    ```
    go build .
    ```

    ```
    go run .
    ```
* El servidor empezará a ejecutarse en el puerto 8080 (si no lo cambiaste) para testearlo ir a Postman o a un navegador y hacer una request a: http://localhost:8080/api/template/healthy

## Despliegue en AWS

Este repo contiene un archivo llamado **buildspec.yml** donde se define el archivo de build para CodeBuild. Está configurado para:

* Configurar .env de acuerdo a la variable ENVIRONMENT
* Se configura el nombre del container con la variable CONTAINER_NAME
* Se Logea en Github de forma PRIVADA cosa de que se pueda hacer referencia a repos y librerías privadas. Para ello en Codebuild se debe configurar una variable de ambiente llamada GITHUB_TOKEN, para el cual se debe referenciar un Secret (Que es eso de copiar tokens como texto plano??). También se debe configurar cual es la cuenta de github en cuestión (en mi caso github.com/pcartas)
* Buildea la imagen de docker y le coloca 2 tags: latest y hash. El tag con hash será utilizado para desplegar en ECS o EKS ya que EC2 cachea imágenes y si no explicitamos la versión al hacer el deploy puede no tomar efecto
* Se pushea esto al repo de ECR definido en ECR_NAME


## TO DO
Falta agregar la parte de Sonar y Testeos Unitarios
