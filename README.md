# Arquitetura hexagonal

---

## Explicando a arquitetura:

### Pasta `application`

Vamos utilizar a pasta `application` para definir o **coração da nossa aplicação**, que ficará isolada do mundo externo.

Utilizaremos a `application/product.go` para definir as interfaces, estrutura do produto e as regras de negócio.

Em `application/product_service.go` criaremos os métodos para chamar a classe de persistência (que ficará fora da application) e conta
com algumas validações caso aconteça algum erro durante a persistência dos dados.

### Pasta `adapters`

Vamos utilizar a pasta `adapters` para definir os componentes externos da nossa aplicação, como a persistência no banco
de dados, filas, web, entre outros.

#### Pasta `adapters/cli`

Esse será o lado esquerdo da arquitetura hexagonal, que representará o cliente.

Utilizaremos a `adapters/cli/product.go` para passar uma instrução de criar/buscar/atualizar um produto por linha de comando.

Em `adapters/cli/product_test.go`, criaremos os testes do adapter CLI.

#### Pasta `adapters/db`

Esse será o lado direito da arquitetura hexagonal, que fará a conexão com o db.

Utilizaremos a `adapters/db/product.go` para criar a persistência com o banco de dados.

O arquivo `sqlite.db` na raiz do projeto, **será o nosso banco de dados**, para **criar as tabelas** abra o bash e digite:
`sqlite3 sqlite.db` e rode os comando: `create table products(id string, name string, price float, status string);`

Em `adapters/db/product_test.go`, criaremos os testes no banco.

### Raiz do projeto

Na raiz do projeto contamos com o arquivo `mainteste.go` que é um modo de **criarmos uma conexão** dos `adapters` na `application`
e realizar uma inserção e atualização de um produto no banco.

Temos o arquivo `sqlite.db` que é utilizado na `mainteste.go` para salvar os **insert/update/delete** do banco.

Para rodar essa classe, basta digitar o comando no bash: `go run mainteste.go`.

---

### Como subir o projeto:

Abra um terminal na **pasta raiz**, e logo em seguida digite: `docker-compose up -d`.

Para checar se o docker está de pé, rode o comando: `docker ps`.

Acesse o bash com: `docker exec -it appproduct bash`.

**Para rodar as classes de testes Go**, rode o comando no mesmo bash: `go test ./...`.
> Caso queira rodar uma classe de teste em específico, utilize `go test ./application/{arquivo_teste}`.

![img.png](readme_images/img.png)

### Para evitar bugs de versionamento:

Tenha preferência por utilizar as **mesmas versões do arquivo** `go.mod` existente, **copie** todos os `require`,
apague o arquivo `go.mod` existente e cole no arquivo `go.mod` que você criará logo abaixo:

Vamos criar um arquivo **go.mod** dentro do bash com o comando: `go mod init github.com/gui-meireles/go-hexagonal`, para facilitar
o download dos pacotes que vamos trabalhar.

---

### Comandos úteis do application

Para criar mocks facilmente em Go, podemos utilizar o **_mockgen_**, que é um auto gerador de mocks, para utilizá-lo
digite o comando a seguir: `mockgen -destination=application/mocks/application.go -source=application/product.go application`
dentro do bash do container.

### Comandos úteis do db

**Para criar o arquivo sqlite.db**, utilize o comando dentro do bash: `touch sqlite.db`.

Para abrir o terminal do db, rode o comando dentro do bash: `sqlite3 sqlite.db`.

Para checar as tabelas: `.tables`.

### Sobre o Cobra

É uma **biblioteca popular em Go** que permite criar aplicativos de linha de comando de forma _fácil e eficiente_.

Fornecendo uma estrutura para definir comandos, flags e argumentos.

Para utilizá-lo em sua aplicação Go, você deve baixá-lo através do **Dockerfile**. (Utilize a versão v1.1.3)

E iniciar o Cobra com o comando: `cobra init --pkg-name=github.com/gui-meireles/go-hexagonal/application`.

Ele criará os arquivos `main.go` e `cmd/root.go`.

Caso você tente rodar o go run main.go e apareça alguns erros de pacotes, rode o comando: `go mod tidy`, ele removerá
os pacotes que não estiverem sendo utilizados e baixará os necessários.