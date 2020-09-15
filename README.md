
# Api en Go para rotar matriz NxN

  Dado un array NxN enviado por post se retornara un array de la misma dimensión pero rotado en sentido antiorario 9 grados.
  
  Se utilizó [Gin Gonic Framework](https://github.com/gin-gonic/gin) para el servidor http.

  
### Parámetros de entrada (JSON):
- matrix: Matriz bidimensional de enteros N x N.
Ejemplo:
{ "matrix":[[1,2],[3,4]] }
 
### Parámetros de salida (JSON):
matrix: Matriz bidimensional de enteros N x N.

Oupput: {  "matrix":[[2,4],[1,3]]  }
### Ejecución:
1. Clonar repositorio
2. Instalar las depentencias utilizadas:
	- [ ] go get github.com/gin-gonic/gin
	- [ ] go get github.com/gin-contrib/cors
3. Ejecutar con el comando go run main.go

### Pruebas
curl --request POST --url http://127.0.0.1:8080/ --header 'content-type: application/json' --data '{"array": [[1,2],[3,4]]}'

**Output**:  {"array":[[2,4],[1,3]]}
  
## Author
By [Orlando A. Yepes](https://www.linkedin.com/in/orlando-andr%C3%A9s-yepes-miquilena-9649544a/)