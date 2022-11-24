#!/usr/bin/python3

import subprocess

DATABASE_PATH = "apps/api/database/kingclover.sqlite"

def main():
  subprocess.Popen(["sqlitebrowser", DATABASE_PATH], start_new_session=True)

if __name__ == '__main__':
  main()