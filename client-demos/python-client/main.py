import grpc
from datetime import datetime

# import the generated classes
from hellosvc import helloworld_pb2, helloworld_pb2_grpc


def main():
    # open a gRPC channel
    channel = grpc.insecure_channel('localhost:50051')

    # create a stub (client)
    stub = helloworld_pb2_grpc.HelloWorldSvcStub(channel)

    # create a valid request message
    t1 = datetime.now()
    for i in range(1000):
        req = helloworld_pb2.HelloRequest(name=f'yang {i}')
        # make the call
        response = stub.SayHello(req)
        print(response)
        stub.SayBye(helloworld_pb2.ByeRequest(name=f'yang {i}'))
    print((datetime.now() -t1 ).total_seconds()*1000)

if __name__ == "__main__":
    main()
