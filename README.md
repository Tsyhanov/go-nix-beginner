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

![Result](https://github.com/Tsyhanov/go-nix-beginner/blob/5e407262027b3f3f728b5ec4c5571ba88df305e9/task-3/task-3.png)

## Task-4 (Gorutines)
- Get post's structure concurrently
- Print it in console
- Note: system can not guarantee the execution time for gorutines. 
Without  sync we always see different sequence of posts in console output

![Result](https://github.com/Tsyhanov/go-nix-beginner/blob/5e407262027b3f3f728b5ec4c5571ba88df305e9/task-4/task-4.png)

## Task-5 (File system)
- Get post's structure concurrently
- Write structures in files. Separate file for each post
- Use ioutil and bufio as well
- Note: iobuf stores data in memory. And any changes are in memory while flush() not being called.

![Result](https://github.com/Tsyhanov/go-nix-beginner/blob/5e407262027b3f3f728b5ec4c5571ba88df305e9/task-5/task-5.png)

## Task-6 (Database)
- Get posts for UserId = 7 concurrently
- Insert Posts and Comments for them into MySql database concurrently

![Posts table](https://github.com/Tsyhanov/go-nix-beginner/blob/5e407262027b3f3f728b5ec4c5571ba88df305e9/task-6/task-6-posts.png)

![Comments table](https://github.com/Tsyhanov/go-nix-beginner/blob/5e407262027b3f3f728b5ec4c5571ba88df305e9/task-6/task-6-Comments.png)




