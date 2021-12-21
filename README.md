# NIX Education Golang Beginner

## Task-1 (OS/Go/IDE)
Environment setup:
Ubuntu 20.04
Golang 1.13.8
VSCode (Goland will be used fo future tasks)

## Task-2 (Github link)
https://github.com/Tsyhanov/go-nix-beginner

## Task-3 (Get request)
- Get posts from https://jsonplaceholder.typicode.com/
- Print it in console

## Task-4 (Gorutines)
- Get post's structure concurrently
- Print it in console
Note: system can not guarantee the execution time for gorutines. 
Without  sync we always see different sequence of posts in console output

## Task-5 (File system)
- Get post's structure concurrently
- Write structures in files. Separate file for each post
- Use ioutil and bufio as well
Note: iobuf stores data in memory. And any changes are in memory while flush() not being called.

## Task-6 (Database)
- Get posts for UserId = 7 concurrently
- Insert it into MySql database in parallel


