-- 1683. Invalid Tweets
SELECT tweet_id
FROM Tweets
WHERE LENGTH(content) > 15;



-- 1148. Article Views I

Select DISTINCT author_id as id 
from Views 
where author_id = viewer_id 
order by author_id;


-- 595. Big Countries

SELECT name, population, area 
FROM World
WHERE area >= 3000000 OR population >= 25000000;

-- 584. Find Customer Referee

SELECT name 
FROM Customer 
WHERE COALESCE(referee_id, 0) <> 2;
-- or
Select name 
from Customer 
where referee_id != 2 or referee_id is null;

-- 1757. Recyclable and Low Fat Products


Select product_id
from Products
where low_fats = 'Y' and recyclable = 'Y';