const {HelloRequest, HelloReply} = require('./helloworld.proto')
const {GreeterClient} = require('./helloworld_grpc_web_pb')

const client = new GreeterClient('http://localhost:8080')

const request = new HelloRequest()
request.setName('World')

client.sayHello(request, {}, (err, response) => {
    console.log(response.getMessage())
})