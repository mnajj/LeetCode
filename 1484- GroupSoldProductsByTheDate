SELECT sell_date,
       count(DISTINCT product) AS num_sold,
       string_agg(product, ',') within GROUP (
                                              ORDER BY product ASC) AS products
FROM
  (SELECT DISTINCT sell_date,
                   product
   FROM activities) a
GROUP BY sell_date
ORDER BY sell_date
