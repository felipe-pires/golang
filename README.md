# go get url-pacote - pacote externo
# ex: go getgithub.com/badoux/checkmail 


# go mod tidy - remover dependencias nao utilizadas

# go mod init nome-modulo - criar modulo

# go test - executar testes
# go test ./... - testes de todos os pacotes
# go test -v
# go test --cover - cobertura dos testes
# go test --coverprofile nomeArquivo.txt - relatorio de cobertura

# go tool cover --func=nomeArquivo.txt - ler relatorio de cobertura
# go tool cover --html=nomeArquivo.txt - mostra uma pagina html com as linhas cobertas e nao cobertas pelos testes com base no relatorio de cobertura