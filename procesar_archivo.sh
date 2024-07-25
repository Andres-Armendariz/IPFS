#!/bin/bash
# Verifica si se ha pasado un archivo como argumento
if [ $# -eq 0 ]; then
 echo "Uso: $0 nombre_del_archivo"
 exit 1
fi
archivo=$1
# Verifica si el archivo existe
if [ ! -f "$archivo" ]; then
 echo "El archivo '$archivo' no existe."
 exit 1
fi
# Obtiene el tamaño del archivo
tamano=$(stat -c%s "$archivo")
echo "Tamaño del archivo: $tamano bytes"
# Imprime las tres primeras líneas
echo "Primeras 3 líneas del archivo:"
head -n 3 "$archivo"
# Imprime las tres últimas líneas
echo "Últimas 3 líneas del archivo:"
tail -n 3 "$archivo"
