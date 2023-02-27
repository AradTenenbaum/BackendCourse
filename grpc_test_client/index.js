
const grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
const PROTO_PATH = "../proto/service_simple_bank.proto";

const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);


const SimpleBank = grpc.loadPackageDefinition(packageDefinition).pb.SimpleBank;

const client = new SimpleBank(
  "localhost:9090",
  grpc.credentials.createInsecure()
);

console.log("Create a user:")
client.CreateUser({ 
  username: "gRPC-user", full_name: "gRPC", email: "gRPC@email.com", password: "secret"
}, (error, data) => {
  if (error) throw error
  console.log(data);
  console.log("Login from user:")
  client.LoginUser({ 
    username: "gRPC-user", password: "secret"
  }, (error, data) => {
    if (error) throw error
    console.log(data);
  });
});