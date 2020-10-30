from concurrent import futures
import logging

from protobuf.opencv_pb2_grpc import *
from protobuf.opencv_pb2 import *


class Greeter(GreeterServicer):

    def SayHello(self, request, context):
        print(request.name)
        return HelloReply(message='Hello, %s!' % request.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('127.0.0.1:10002')
    server.start()
    print('OPEN')
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()