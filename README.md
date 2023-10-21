# API para Site de Estudos - Em construção :hammer:

Site desenvolvido durante o curso de devops do ittalent2023. O objetivo do repositório é a criação de uma api com a finalidade de desafiar-se durante o curso. 

O estudo de golang veio como uma proposta do curso, com o intuíto de aprender uma nova linguagem de programação comumente utilizada no devops.

## Arquitetura

Arquitetura em camadas. 

Objetivo: Utilizar Drive Domain Design (DDD) para modelagem do software. 

##### Links:

Links de apoio para criação da arquitetura hexagonal caso adotemos a mesma no decorrer do desenvolvimento.

- https://alistair.cockburn.us/hexagonal-architecture/
- https://www.linkedin.com/pulse/arquitetura-hexagonal-com-golang-vinicius-nordi-esperan%C3%A7a/?originalSubdomain=pt

## Base dados

Nesse projeto utilizaremos a base de dados SQL.

## Documentação

Utilizamos o swagger para documentar a API.

##### Pacotes

- go get -u github.com/swaggo/swag/cmd/swag
- go get -u github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/files

##### Necessário exportar a variável de ambiente
 
- export PATH=$(go env GOPATH)/bin:$PATH

##### Evolução da modelagem do banco


![Database](/src/utils/svg/initialDatabase.svg)