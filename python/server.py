from concurrent import futures
import logging
import grpc


import protobuf.media_pb2 as media_pb
import protobuf.media_pb2_grpc as media_grpc

class Service(media_grpc.ServiceServicer):

    def SendMessage(self, request, context):
        print(request.Content)
        message = media_pb.RplMessage(Content=request.Content)

        return message


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    media_grpc.add_ServiceServicer_to_server(Service(), server)
    server.add_insecure_port('127.0.0.1:10002')
    server.start()
    print('OPEN')
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()