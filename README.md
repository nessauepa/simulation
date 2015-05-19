# simulation

### Leitura

  Tutorial (tour) da linguagem GO: https://tour.golang.org/welcome/1
	Boas práticas: https://golang.org/doc/effective_go.html (pelo menos o comecinho e importante)

### Instalar GO: https://golang.org/doc/install

	tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
	export PATH=$PATH:/usr/local/go/bin

### Configurar workspace: https://golang.org/doc/code.html

	vschissato@vschissato:~$ tree workspace-golang/
	workspace-golang/
	├── bin
	├── pkg
	└── src
   	 └── github.com
        	└── nessauepa
        	    └── simulation

	export GOPATH=$HOME/workspace-golang
	export PATH=$PATH:$GOPATH/bin (conveniência)

### Build e install com GO Tools: https://golang.org/cmd/go/

	$ go install github.com/nessauepa/simulation
	$ simulation

### Convenções de testes unitários: https://golang.org/doc/code.html#Testing

	$ go test github.com/nessauepa/simulation
