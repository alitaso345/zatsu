const PROTO_PATH = __dirname + '/helloworld.proto'

const grpc = require('grpc')
const _ = require('lodash')
const async = require('async')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
)
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)
const helloworld = protoDescriptor.helloworld

/**
 * @param {!Object} call
 * @param {function():?} callback
 */
function doSayHello(call, callback) {
    callback(null, {message: 'Hello! '+ call.request.name});
}

/**
 * @param {!Object} call
 */
function doSayRepeatHello(call) {
    const senders = [];
    function sender(name) {
        return (callback) => {
            call.write({
                message: 'Hey! ' + name
            });
            _.delay(callback, 500); // in ms
        };
    }
    for (const i = 0; i < call.request.count; i++) {
        senders[i] = sender(call.request.name + i);
    }
    async.series(senders, () => {
        call.end();
    });
}

/**
 * @param {!Object} call
 * @param {function():?} callback
 */
function doSayHelloAfterDelay(call, callback) {
    function dummy() {
        return (cb) => {
            _.delay(cb, 5000);
        };
    }
    async.series([dummy()], () => {
        callback(null, {
            message: 'Hello! '+call.request.name
        });
    });
}

/**
 * @return {!Object} gRPC server
 */
function getServer() {
    const server = new grpc.Server();
    server.addService(helloworld.Greeter.service, {
        sayHello: doSayHello,
        sayRepeatHello: doSayRepeatHello,
        sayHelloAfterDelay: doSayHelloAfterDelay
    });
    return server;
}

if (require.main === module) {
    const server = getServer();
    server.bind('0.0.0.0:9090', grpc.ServerCredentials.createInsecure());
    server.start();
}

exports.getServer = getServer;
