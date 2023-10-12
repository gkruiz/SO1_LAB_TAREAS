const WebSocket = require('ws');
const redis = require('redis');

const http = require("http");
const { Server } = require("socket.io");


const httpServer = http.createServer();

var valores=[]

const PORT = 3000,
  HOST = "localhost";
// Configuration: adapt to your environment
const REDIS_SERVER = "redis://127.0.0.1:6379";

// Connect to Redis and subscribe to "app:notifications" channel
const subscriber = redis.createClient(REDIS_SERVER);
subscriber.connect();


const io = new Server(httpServer, {
  cors: {
    origin: "*", // or a list of origins you want to allow, e.g. ["http://localhost:3000"]
    credentials: true,
  },
});

io.on("connection", (socket) => {
  console.log(socket.id);
  // broadcast on web socket when receving a Redis PUB/SUB Event
  subscriber.subscribe('pruebas', (message) => {
    console.log(valores); 
    valores.push(JSON.parse(message))
    
    subscriber.unsubscribe();
    //subscriber.quit();
    io.send(JSON.stringify(valores));
  });

  

});

httpServer.listen(PORT, HOST, () => {
  console.log("Server running on port:", PORT);
});















 /*
// Configuration: adapt to your environment
const REDIS_SERVER = "redis://127.0.0.1:6379";
const WEB_SOCKET_PORT = 3000;
 
// Connect to Redis and subscribe to "app:notifications" channel
var redisClient = redis.createClient(REDIS_SERVER);
redisClient.subscribe('app:notifications');
 
// Create & Start the WebSocket server
const server = new WebSocket.Server({ port : WEB_SOCKET_PORT });
 
// Register event for client connection
server.on('connection', function connection(ws) {
 
  // broadcast on web socket when receving a Redis PUB/SUB Event
  redisClient.on('message', function(channel, message){
    console.log(message);
    ws.send(message);
  })
 
});*/
