# Pré-requisitos:

Docker Versão 20.10.12

Google Chrome Versão 99.0.4844.84 (Versão oficial)


# Acessando o webserver

1 - Certifique-se de que a porta 8088 nao está em uso em seu computador
2 - Inicie o Docker (TODO)
3 - Acesse o endereço http://localhost:8088


# Premissas do Projeto

* Os dados dos clientes (válidos) serão inseridos na tabela customer_data.
* Os dados rejeitados (por nao atenderem ao padrão) serão inseridos na tabela customer_data_rejected sem qualquer tipo de tratamento.
* A aplicação não deve assumir nenhum tipo de correção nos dados (ex.: Linhas quebradas ou ausência de colunas pois não há como inferir a integridade do dado).
* Os arquivos importados serão inseridos na tabela "imported_files" com a data/hora da iportacao para registro e consulta futura.
* Não há necessidade de validacao de letras maiúsculas ou acentuação pois os dados são basicamente numéricos (em suas variações). Isto geraria um consumo desnecessário de recursos de processamento. Serão validados apenas a pontuação do CNPJ ou CPF.


# TO DO
* Validar a formatacao do campo CPF
* Validar a formatação do campo CNPJ





# Referências de Código

## Validação de CPF e CNPJ: 

https://medium.com/@andersonantunes/fun%C3%A7%C3%B5es-de-valida%C3%A7%C3%A3o-de-cpf-e-cnpj-em-pl-pgsql-5720d49b4215#:~:text=Valida%C3%A7%C3%A3o%20do%20CPF,descrente%20de%2010%20at%C3%A9%202.

## File Upload tutorial

https://github.com/TannerGabriel/learning-go/tree/master/beginner-programs/FileUpload
https://tutorialedge.net/golang/go-file-upload-tutorial/

# Remover acentuacao

https://devtools.com.br/blog/retirando-acentuacao-no-postgresql/

# REST

https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/

