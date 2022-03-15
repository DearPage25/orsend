# orsend

This proyect is a simple, client and server that send and receive files.

For start this program, first see it like two proyects, client and server.

## Client

1. Write in the terminal: 
      ```
         $ cd client
           go build src/*.go
      ```
  
   This will create a file calls ```client.exe``` or simply ```client```.
   
2. For start the ```client``` write in the terminal one of two commands:
   
   2.1 For ```client``` to be in ```receive``` mode, write this in the terminal :
   
       $ ./client receive channel 1
 
      This put the client in ```listening``` mode and you can change the value channel (this value is a string not a number or int value).
  
   2.2 For ```client``` to be in ```send``` mode, write this in the terminal:
      ```
      $ ./client send -file something.txt -channel 1
      ```
      
      This ```send``` mode include two sub-commands: ```-file``` take any file in the directory ```./client/docs``` with the only condition: **write exactly name file.** And file will send in the ```channel``` value, in this case is ```1```.
      
 **Important**
 
   For this program to run correctly, the commands are required.
   
  receive:
      
      
       $ ./client receive channel <name or number channel>
      
       
   send:
      
     
      $ ./client send -file <name file> -channel <name or number channel>
      
## Server

In this case just write in the terminal:
      ```
         $ cd server
           go build src/*.go
      ```
Also write in the terminal: 

```
   $ ./app
```
And that's all, the server is running.
