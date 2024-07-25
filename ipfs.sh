#!/bin/bash

# Función para imprimir mensajes de error y salir
error_exit() {
    echo "$1" 1>&2
    exit 1
}

# Descargar Kubo (IPFS)
echo "Descargando Kubo (IPFS)..."
wget https://dist.ipfs.tech/kubo/v0.29.0/kubo_v0.29.0_linux-amd64.tar.gz -O kubo.tar.gz || error_exit "Error descargando Kubo."

# Descomprimir el archivo descargado
echo "Descomprimiendo Kubo..."
tar -xzf kubo.tar.gz || error_exit "Error descomprimiendo Kubo."

# Mover el binario a /usr/local/bin
echo "Instalando Kubo..."
sudo mv kubo/ipfs /usr/local/bin/ipfs || error_exit "Error moviendo Kubo a /usr/local/bin."

# Limpiar archivos temporales
sudo rm -rf kubo kubo.tar.gz

# Inicializar el nodo IPFS
echo "Inicializando el nodo IPFS..."
ipfs init || error_exit "Error inicializando el nodo IPFS."

# Configurar IPFS para arrancar automáticamente al iniciar el sistema (opcional)
echo "Configurando IPFS para iniciar automáticamente..."
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["*"]' || error_exit "Error configurando CORS."
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "GET", "POST"]' || error_exit "Error configurando métodos CORS."

# Mostrar información del nodo
ipfs id
