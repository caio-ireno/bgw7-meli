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
select series.title, COUNT(seasons.number) AS total_seasons 
from series 
join seasons on seasons.serie_id = series.id 
group by series.title;
```
#### Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.
```sql
select count(movies.title) as movies, genres.name 
from movies 
join genres on genres.id= movies.genre_id 
group by genres.name

select count(movies.title) as movies, genres.name 
from movies 
join genres on genres.id= movies.genre_id 
and movies.rating >3
group by genres.name 
```
#### Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.
```sql
SELECT 
  a.first_name, 
  a.last_name
FROM actors a
JOIN actor_movie am ON am.actor_id = a.id
JOIN movies m ON m.id = am.movie_id
WHERE m.title LIKE '%guerra de las%'
GROUP BY a.id, a.first_name, a.last_name
HAVING COUNT(DISTINCT m.id) = (
    SELECT COUNT(*) FROM movies WHERE title LIKE '%guerra de las%'
);


and 


select actors.first_name, actors.last_name
from actor_movie 
join actors on actors.id=actor_movie.actor_id
join movies on movies.id = actor_movie.movie_id
where movies.title like '%guerra de las%'
GROUP BY actors.id, actors.first_name, actors.last_name;
```