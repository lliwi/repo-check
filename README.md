

## Contenido
* [repo-check](#repo-check)
* [Instalación](#instalaci%C3%B3n)
* [Uso](#uso)
* [Inspiración](#inspiraci%C3%B3n)

## repo-check
repo-check es un pequeño script escrito en go que nos ayuda a buscar distintos repositorios para posteriormente poder obtener todos los datos con dvcs-ripper.

Se puede usar tanto para una única URL o leyendo un fichero con distintos dominios o URLs. Para la lectura de fichero se ha implementado una expresión regular que identificara línea por línea los dominios o URLs válidas.

En el caso de los dominios el sistema intentará buscar los distintos repos tanto en http como en https.

## Instalación
Al ser un sript escrito en go solo es necesario ejecutar:
```
$ go build repo_check.go
```

## Uso
```
$ ./repo_check -u [url o dominio]
```


```
$ ./repo_check -f [path del fichero]
```

Una vez obtengamos los resultados podemos usar dvcs-ripper para obtener los datos del repo.


## Inspiración
Este script está inspirado en el proyecto de kost, dvcs-ripper. Pretendiendo únicamente localizar con agilidad los repositorios.
