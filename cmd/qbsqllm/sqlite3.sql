--*- mode: sql; sql-product: sqlite; -*-
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
);

CREATE TABLE args (
  entry INTEGER NOT NULL REFERENCES entry(id)
, name TEXT NOT NULL
, value TEXT NOT NULL
);
