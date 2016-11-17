[![Build Status](https://travis-ci.org/tcc-denis-raul/Projeto-api.svg?branch=master)](https://travis-ci.org/tcc-denis-raul/Projeto-api)
Sobre:
======
Servidor web escrita em go

Pré-dependencias:
================
- Go >= 1.5.4
- Govendor
	- go get -u github.com/kardianos/govendor

Configurações:
==============
- Edite um arquivo chamado paloma.json com as configurações do banco
Exemplo:
`{
      "database": {
         "url": "127.0.0.1",
         "name": "paloma"       
      }
}`
- Copie 'paloma.json' para a pasta '/etc'

Instalando depencias:
=====================
- export GO15VENDOREXPERIMENT=1
- make deps
- Obs. A pasta precisa estar no diretório $GOPATH

Compilando e executando:
========================
- export API_PORT=Port (default: 5000)
- make run

Executando os testes:
=====================
- make test
