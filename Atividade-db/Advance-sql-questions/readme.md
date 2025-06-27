## O que é um JOIN em banco de dados e para que ele é usado?

Similar a uma expressão de junção em algebra linear, uma cláusula JOIN, combina colunas de uma ou mais tabelas em um banco de dados.

Dentro que chamamos de junção, temos alguns comandos comuns como, JOIN: INNER JOIN, LEFT JOIN, RIGHT JOIN, FULL JOIN e CROSS JOIN.

A imagem abaixo representa esses comandos mais comuns.

![alt text](image.png)

## Querys

#### Exibir o título e o nome do gênero de todas as séries.

```sql
select s.title, g.name 
from series s 
left join genres g on s.genre_id = g.id;
```

#### Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.

```sql
select e.title, a.first_name, a.last_name  from actor_episode ae 
join actors a on ae.actor_id = a.id
join episodes e on ae.episode_id = e.id;
```


#### Mostre o título de todas as séries e o número total de temporadas de cada série.
```sql
```
#### Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.
```sql
```
#### Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.
```sql
```