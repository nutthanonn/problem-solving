SELECT user_id, CONCAT(UPPER(LEFT(name, 1)), LOWER(SUBSTR(NAME, 2, LENGTH(NAME))))  
AS 'name'
FROM USERS
order by user_id