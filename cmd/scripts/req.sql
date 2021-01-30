INSERT INTO user(login, first_name, last_name, birth_date, job_title, city) values (?, ?, ?, ?, ?, ?);

SELECT u.* FROM users AS u INNER JOIN ((SELECT DISTINCT fst FROM friends WHERE snd = 2) UNION (SELECT DISTINCT snd FROM friends WHERE fst = 2)) AS fr ON u.user_id = fr.fst;

SELECT h.* FROM hobbies as h INNER JOIN (SELECT hobby_id FROM user2hobbies WHERE user_id = 5) AS u ON h.id = u.hobby_id;