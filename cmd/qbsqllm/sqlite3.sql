--*- mode: sql; sql-product: sqlite; -*-
CREATE TABLE file (
  id INTEGER PRIMARY KEY
, name TEXT NOT NULL
);

CREATE TABLE form (
  id INTEGER PRIMARY KEY
, hash BLOB NOT NULL UNIQUE
, text TEXT NOT NULL
);

CREATE TABLE entry (
  id INTEGER PRIMARY KEY
, ts TEXT NOT NULL
, level TEXT NOT NULL
, log TEXT NOT NULL
, code TEXT
, form INTEGER NOT NULL REFERENCES form(id)
, file INTEGER REFERENCES file(id)
, line INTEGER NOT NULL
);

CREATE TABLE args (
  entry INTEGER NOT NULL REFERENCES entry(id)
, name TEXT NOT NULL
, value TEXT NOT NULL
);

CREATE VIEW list AS
SELECT f.name, e.line, e.ts, e.level, e.log, e.code, t.text
FROM entry e
JOIN file f ON (e.file = f.id)
JOIN form t ON (e.form = t.id)
ORDER BY e.ts, f.name, e.line;

