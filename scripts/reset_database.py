#!/usr/bin/python3

import sqlite3

DATABASE_PATH = 'apps/api/database/kingclover.sqlite'

def main():
  conn = sqlite3.connect(DATABASE_PATH)
  cur = conn.cursor()
  res = cur.execute('SELECT name FROM sqlite_master WHERE type = "table" AND name NOT LIKE "sqlite_%"')
  tables = res.fetchall()

  for table in tables:
    cur.execute(f'DROP TABLE IF EXISTS {table[0]};')

  conn.close()

if __name__ == '__main__':
  main()
