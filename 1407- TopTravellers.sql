SELECT name,
       ISNULL(SUM(distance), 0) AS travelled_distance
FROM Users AS u
LEFT JOIN Rides AS r ON u.id = user_id
GROUP BY u.id,
         name
ORDER BY travelled_distance DESC,
         name
