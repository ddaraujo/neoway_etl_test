## Pré-requisitos:

Docker Versão 20.10.12.  
Google Chrome Versão 99.0.4844.84 (Versão oficial).  
Certifique-se de que as portas 8888 e 5433 não estão em uso no computador onde o sistema será executado.  


## Premissas do Projeto

* Os dados dos clientes (válidos) serão inseridos na tabela customer_data.
* Os dados rejeitados (por nao atenderem ao padrão) serão inseridos na tabela customer_data_rejected sem qualquer tipo de tratamento.
* A aplicação não deve assumir nenhum tipo de correção nos dados (ex.: Linhas quebradas ou ausência de colunas pois não há como inferir a integridade do dado).
* Os arquivos importados serão inseridos na tabela "imported_files" com a data/hora da iportacao para registro e consulta futura.
* Não há necessidade de validacao de letras maiúsculas ou acentuação pois os dados são basicamente numéricos (em suas variações). Isto geraria um consumo desnecessário de recursos de processamento. Serão validados apenas o CNPJ e CPF.

## Download do repositório

git clone https://github.com/ddaraujo/neoway_etl_test.git

## Inicializando (Building Server)

Antes de começar, inicie o serviço do Docker em seu computador.

```bash
$ cd neoway_etl_test
$ docker compose up --build
```


## Importando o arquivo de dados

+ Acesse o endereço http://localhost:8888
+ Selecione o arquivo desejado e clique em "upload".
+ Aguarde o retorno de sucesso ou falha da importação.
+ Utilize a API ou PGadmin para checar os dados.


## Utilizando a API

* **/items/valid**   ->  retorna todos os items importados válidos
* **/items/valid/cpf/{cpf}**   ->  retorna todos os dados importados válidos por CPF (Apenas números)
* **/items/valid/lastSale/{cnpj}**   ->  retorna todos os dados importados validos (ultima loja) por CNPJ
* **/items/valid/frequentSale/{cnpj}**   ->  retorna todos os dados importados validos (loja mais frequente) por CNPJ 
* **/items/valid/count**   ->  retorna a quantidade de registros válidos
* **/items/valid/delete**   ->  limpa a tabela de registros válidos
* **/items/invalid**   ->  retorna todos os items importados válidos
* **/items/invalid/count**   ->  retorna a quantidade de registros inválidos (rejeitados)
* **/items/invalid/delete**   ->  limpa a tabela de registros inválidos

## Conexão ao database

**host:** localhost  
**porta:** 5433  
**usuário:** postgres  
**senha:** postgres  

## Estrutura
### Database
#### Tabelas
* **customer_data** - Dados válidos importados já contendo validação de CPF e CNPJ
* **customer_data_rejected** - Dados rejeitados (inválidos) que não passaram pela validacao de campos
* **imported_files** - Dados RAW importados do arquivo antes da sanitização e classificação. São removidos após o fim da importação.

#### Projeto
```bash
├── db
│   ├── db.go         
|   ├── file.go
│   └── item.go
├── handler                        
│   ├── files.go
│   ├── handler.go
│   └── items.go
├── models
│   ├── file.go
│   ├── item.go
│   └── record.go
├── sql
│   ├── create_triggers.sql
│   ├── init_tables.sql
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── index.html
├── main.go
└── README.md
```

**db:** Pacote responsável pela interação direta com o banco de dados, separando as camadas de acesso ao DB do restante da aplicação.handler: Cria os handlers do app e as rodas da API utilizando gorilla/mux.  
**models:** Structs de objetos para acesso e consulta ao database ou transformados em formato JSON.  
**sql:** Scripts de inicializacao das tabelas, funcões, triggers, etc.  
**.env:** Variáveis de ambiente utilizadas pela aplicação (conexão ao database).  
**docker-compose:** Define as dependencias dos microserviços do app (app e Database).  
**Dockerfile:** Imagem base e comandos para inicialização do app.  
**index.html:** Página simples para upload do arquivo.  
**main.go:** Entrypoint da app. Responsável pela inicialização do database, leitura de variáveis do .env e inicializar/encerrar a API e handlers.  
