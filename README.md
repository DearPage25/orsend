# orsend
## Info
This proyect is a simple, client and server that send and receive files.

For start this program, first see it like two proyects, client and server.

## Client

1. Write in the terminal: 
      ```
         $ client
           go build src/*.go
      ```
  
   This will create a file calls ```client.exe``` or simply ```client```.
   
2. Start the ```client``` write in terminal one a two commands:
   
   2.1 For ```client``` will be a ```receive``` state, write this:
   
       $ ./client receive channel 1
 
      This will made the client in ```listening``` mode and you can to change value channel (this value is a string not a number or int value).
  
   2.2 For ```client``` will be a ```send``` state, write this:
      ```
      $ ./client send -file something.txt -channel 1
      ```
      
      This state ```send``` mode. ```-file``` take any file in the directory ```./client/docs``` with the one condition: **write exactly name file.** And file will send in the ```channel``` value, in this case is ```1```.
      
 **Important**
 
   For this program to run correctly, the commands are required.
   
  receive:
      
      
       $ ./client receive channel <name or number channel>
      
       
   send:
      
     
      $ ./client send -file <name file> -channel <name or number channel>
      
## Server

In this case just write in the terminal:
      ```
         $ server
           go build src/*.go
      ```
Also write in the terminal: 

```
   $ ./app
```
And that all, the server is running.
