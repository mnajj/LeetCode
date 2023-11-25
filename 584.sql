SELECT name
FROM   customer
WHERE  Isnull(referee_id, 0) != 2; 
