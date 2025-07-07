# Go-Cache-Client

#### For server side please visit  [repo](https://github.com/praveenmahasena/go_cache_server).
#### This project is still on development

Go-Cache-Client is a part of in memory database that is "Redis like" built using Go programming language. </br>
This app communicates to server via TCP

## Features
 - In memory database
 - Concurrent
 - No race conditions
 - Has string,set,hash,list datastructures to store data

## Tech Stack
 - (Golang)[]
 - (Git)[]optional


## Client side Setup

 1. ### Clone repo

    ```bash
    git clone https://github.com/praveenmahasena/go_cache_client.git
    ```

 2. ### Change working dir

    ```bash
    cd /go_cache_client
    ```

 3. ### Build binary

    ```bash
    make
    ```

 4. ### Run
    ```bash
    ./bin/go_cache_client -p "<portID_to_server>"
    ```

## Data structures

 1. ### strings
 - `SET/GET` : Basic set and get
 - `APPEND` : Add to end
 - `STRLEN` : String length
 - `SET/GET` : Basic set and get
 - `SET/GET` : Basic set and get
 - `SET/GET` : Basic set and get

 2. ### list
 3. ### set
 4. ### hash


## TODO

