<p align="center">
 <img src="https://csharpcorner-mindcrackerinc.netdna-ssl.com/article/restfull-minimal-web-api-with-net-6/Images/image-20220522162113-1.png" width="35%"/>
</p>

Golang Web Api Structure
=======================================

[![Licence: MIT](https://img.shields.io/badge/Licence-MIT-green)](LICENCE)

* * *

Este projeto é um microframework de aplicação web com sintaxe expressiva e elegante.
Ao utilizar a estrutura deste projeto você elimina a dor do desenvolvimento facilitando tarefas
comuns usadas em muitos projetos da web, como:

- Serviço de injeção de dependência
- Database ORM
- Gerenciamento de rotas
- Autenticação com JWT

### Requisitos

O framework Laravel possui alguns requisitos:

- Go 1.18
- Docker

### Instalação

A maneira recomendada de instalar este projeto é seguindo estas etapas:

1. Realize o clone do projeto para a sua máquina

```shell
git clone git@github.com:ReinanHS/golang-web-api-structure.git
```

2. Acessar as pastas do projeto

```shell
cd golang-web-api-structure
cp .env.example .env
make up
make server
```

### Software stack

Esse projeto roda nos seguintes softwares:

- Git 2.33+
- Go 1.18
- Gin
- Gorm

### Routing

As rotas aceitam um URI e um encerramento, fornecendo um método muito simples e 
expressivo de definir rotas e comportamento sem arquivos de configuração de roteamento complicados.

Para você definir uma nova rota você deve editar o seguinte arquivo: `internal/http/config/route.go`

```go
func AddRoutes(ctx context.Context, router *gin.Engine) *gin.Engine {

	// Adicione suas rotas aqui
	router.GET("/", user.NewUserController(ctx).Index)
	
	return router
}
```

### Changelog

Por favor, veja [CHANGELOG](CHANGELOG.md) para obter mais informações sobre o que mudou recentemente.

### Seja um dos contribuidores

Quer fazer parte desse projeto? Clique AQUI e leia [como contribuir](CONTRIBUTING.md).

## Segurança

Se você descobrir algum problema relacionado à segurança, envie um e-mail para reinangabriel1520@gmail.com em vez de
usar o issue.

### Licença

Esse projeto está sob licença. Veja o arquivo [LICENÇA](LICENSE.md) para mais detalhes.
