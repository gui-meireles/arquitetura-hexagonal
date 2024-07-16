# Arquitetura hexagonal

---

### Como subir o projeto:

Abra um terminal na **pasta raiz**, e logo em seguida digite: `docker-compose up -d`.

Para checar se o docker está de pé, rode o comando: `docker ps`.

Acesse o bash com: `docker exec -it appproduct bash`.

### Para evitar bugs de versionamento:

Tenha preferência por utilizar as **mesmas versões do arquivo** `go.mod` existente, **copie** todos os `require`,
apague o arquivo `go.mod` existente e cole no arquivo `go.mod` que você criará logo abaixo:

Vamos criar um arquivo **go.mod** dentro do bash com o comando: `go mod init github.com/gui-meireles/go-hexagonal`, para facilitar
o download dos pacotes que vamos trabalhar.

---

### Explicando a arquitetura:

### Pasta `application`

Vamos utilizar a pasta `application` para definir o **coração da nossa aplicação**, que ficará isolada do mundo externo.

Utilizaremos a `application/product.go` para definir as interfaces, estrutura do produto e as regras de negócio.

Em `application/product_service.go` criaremos os métodos para chamar a classe de persistência (que ficará fora da application) e conta
com algumas validações caso aconteça algum erro durante a persistência dos dados.

### Pasta `adapters`

Vamos utilizar a pasta `adapters` para definir os componentes externos da nossa aplicação, como a persistência no banco
de dados, filas, web, entre outros.

Utilizaremos a `adapters/db/product.go` para criar a persistência com o banco de dados.

O arquivo `sqlite.db` na raiz do projeto, **será o nosso banco de dados**, para **criar as tabelas** abra o bash e digite:
`sqlite3 sqlite.db` e rode os comando: `create table products(id string, name string, price float, status string);`

Em `adapters/db/product_test.go`, criaremos os testes no banco.

---

### Comandos úteis do application

Para criar mocks facilmente em Go, podemos utilizar o **_mockgen_**, que é um auto gerador de mocks, para utilizá-lo
digite o comando a seguir: `mockgen -destination=application/mocks/application.go -source=application/product.go application`
dentro do bash do container.

**Para rodar as classes de testes Go**, rode o comando no mesmo bash: `go test ./...`.
> Caso queira rodar uma classe de teste em específico, utilize `go test ./application/{arquivo_teste}`.

![img.png](readme_images/img.png)

### Comandos úteis do db

Para abrir o terminal do db, rode o comando: `sqlite3 sqlite.db`.

Para checar as tabelas: `.tables`.