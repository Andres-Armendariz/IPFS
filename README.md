# Utilización de la Aplicación
Para utilizar esta aplicación, puede clonar el repositorio o descargarlo directamente. Asegúrese de que el sistema en el que va a ejecutar la aplicación tenga un sistema operativo basado en Linux.

## Instalación de IPFS
Para automatizar el proceso de instalación, puede ejecutar el archivo ejecutable. Es posible que necesite otorgarle los permisos necesarios. Para ello, ejecute lo siguiente:
```
chmod +x ipfs.sh
```
Una vez otorgados los permisos, ejecute el archivo en su terminal:
```
./ipfs.sh
```
Probablemente se le solicitarán sus credenciales para completar el proceso. Cuando la ejecución termine, verá una salida similar a la siguiente:
Compruebe la instalación con:

![image](https://github.com/user-attachments/assets/e8719b15-b414-4334-a595-8cd32b70ee79)


Levante el nodo IPFS con:
```
ipfs --version
```
Levante el nodo IPFS
```
ipfs daemon
```
Después de unos minutos, obtendrá una salida similar a:

![image](https://github.com/user-attachments/assets/c62d2858-4137-43f6-8034-5a9e84362582)


Puede ingresar a la dirección indicada en la terminal para comprobar que la interfaz gráfica está operativa.

## Utilización de la aplicación 

Dentro del directorio ipfs-cli, encontrará el ejecutable para efectuar las tres operaciones del sistema de almacenamiento. Primero, debe otorgar permisos de ejecución al archivo ipfs-linux:
```
chmod +x ipfs-linux
```
Puede efectuar una de las siguientes operaciones:

Subir acta
```
./ipfs-linux upload [path del archivo almacenado]
```
 Recupear acta
```
./ipfs-linux download [CID]
```
Auditar/Listar versiones de acta
```
 ./ipfs-linux list [MFS path donde esta el acta almacenada en IPFS]
```
## Comprobar truncamiento de resultados

Dentro del repositorio, encontrará un script y el archivo de las tablas de resultados. En el directorio raíz, otorgue permisos de ejecución:
```
chmod +x procesar_archivo.sh
```
Ejecute el script con:
```
./procesar_archivo [Qma...]
```
No olvide ingresar el nombre completo del archivo.
