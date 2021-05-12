# PFA - Fullcycle 👨‍💻

## Desafio 01

Crie um programa utilizando sua linguagem de programação favorita que faça uma listagem simples do nome de alguns módulos do curso Full Cycle os trazendo de um banco de dados MySQL. Gere a imagem desse container e a publique no DockerHub.

Gere uma imagem do nginx que seja capaz que receber as solicitações http e encaminhá-las para o container.

Crie um repositório no github com todo o fonte do programa e das imagens geradas.

Crie um arquivo README.md especificando quais comandos precisamos executar para que a aplicação funcione recebendo as solicitações na porta 8080 de nosso computador. Lembrando que NÃO utilizaremos Docker-compose nesse desafio.


## Executando o desafio

1. Primeiro passo é criar a rede onde todos os containers estarão inseridos: `docker network create pfa`
2. Em sequencia iremos subir o banco de dados mysql, para tal basta executar o seguinte comando docker `docker run -d --network=pfa --name=mysqlcontainer --rm -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_ROOT_HOST=% -v database:/var/lib/mysql juniorsdj/pfa-mysql`

OBS: Como estamos usando  a base de dados em um volume, quando for a primeira execução do banco precisamos executar um script para popular o banco. para tal basta executar a seguinte sequencia de comandos
. `docker exec -it mysqlcontainer mysql -uroot -ppass`
. `source /usr/src/app/initial-database.sql`

Feito isso, iremos  ter criado o banco PFA dentro do servidor mysql e inicializado com alguns dados.

3.  Seguindo os próximos passos precisamos executar a nossa imagem do node `docker run --name=nodecontainer  --network=pfa --rm -d juniorsdj/pfa-node`

4. Por fim, precisamos agora subir o container com o nginx `docker run --name=nginxcontainer --rm --network=pfa -p 8080:80 juniorsdj/pfa-nginx`