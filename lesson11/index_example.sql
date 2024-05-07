CREATE TABLE index_test(
	id int,
	reason text COLLATE "C",
	annotation text COLLATE "C"
);
-- DROP TABLE index_test;


INSERT INTO index_test
SELECT s.id, md5(random()::text), null
FROM generate_series(1, 10000000) AS s(id)
ORDER BY random();

UPDATE index_test
SET annotation = UPPER(md5(random()::text));


-- начало
SELECT * FROM index_test LIMIT 5;


-- Пример sequence scan/index scan
EXPLAIN SELECT * FROM index_test 
WHERE id = 5000;

SELECT * FROM index_test 
WHERE id = 5000;

CREATE INDEX id_index_test ON index_test(id);
DROP INDEX id_index_test;


-- Пример index only scan
EXPLAIN SELECT id FROM index_test 
WHERE id > 5000;


-- Пример составного index
EXPLAIN SELECT * FROM index_test
WHERE reason LIKE 'bc%' AND annotation LIKE 'AB%';

CREATE INDEX reas_annot_index_test ON index_test(reason, annotation);
DROP INDEX reas_annot_index_test;

EXPLAIN SELECT * FROM index_test
WHERE reason LIKE 'ab%';

EXPLAIN SELECT * FROM index_test
WHERE annotation LIKE 'AB%';

