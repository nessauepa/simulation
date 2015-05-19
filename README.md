# simulation

## Leitura

* Tutorial (tour) da linguagem GO: https://tour.golang.org/welcome/1
* Boas práticas: https://golang.org/doc/effective_go.html (pelo menos o comecinho e importante)

## Instalar GO: https://golang.org/doc/install

	tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
	export PATH=$PATH:/usr/local/go/bin

## Configurar workspace: https://golang.org/doc/code.html

### Criar estrutura de pastas:

	vschissato@vschissato:~$ tree workspace-golang/
	workspace-golang/
	├── bin       <- executáveis (entregáveis)
	├── pkg       <- pacotes de bibliotecas
	└── src       <- código fonte

### Criar variáveis de ambiente

	export GOPATH=$HOME/workspace-golang
	export PATH=$PATH:$GOPATH/bin (conveniência)
	
### Baixar projeto:
	
	go get github.com/nessauepa/simulation
	
	Resultado (código baixado, compilado e instalado):
	vschissato@vschissato:~$ tree workspace-golang/
	workspace-golang/
	├── bin
	│   └── simulation                           <- executável (entregável), hard link com a biblioteca sellerapi.a (portanto, entregável é auto funcional, único arquivo que precisa ser deployado)
	├── pkg
	│   └── linux_amd64
	│       └── github.com
	│           └── nessauepa
	│               └── simulation
	│                   └── sellerapi.a          <- biblioteca seller api client
	└── src
	    └── github.com
	        └── nessauepa
	            └── simulation                    
	                ├── README.md
	                ├── sellerapi
	                │   ├── sellerapi.go        <- Fonte da biblioteca seller api client
	                │   └── sellerapi_test.go   <- Teste unitário
	                └── simulation.go           <- Arquivo principal (importa a biblioteca sellerapi)

## Build e install com GO Tools: https://golang.org/cmd/go/

	$ go install github.com/nessauepa/simulation
	$ simulation

## Convenções de testes unitários: https://golang.org/doc/code.html#Testing

	$ go test github.com/nessauepa/simulation
