from concurrent import futures
import logging
import grpc

import cv2


import protobuf.media_pb2 as media_pb
import protobuf.media_pb2_grpc as media_grpc

from dpkt.rtp import RTP

class Service(media_grpc.ServiceServicer):

    def StreamVideo(self, iterator, context):
        for req in iterator:

            rtp = RTP(req.Rtp).data
            print(rtp)
            cv2.imshow("rtp", rtp)
            cv2.waitKey(1)
        message = media_pb.ReceiveReply(Result="1")
        return message


    def SendMessage(self, request, context):
        print(request.Content)
        message = media_pb.RplMessage(Content=request.Content)

        return message



def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    media_grpc.add_ServiceServicer_to_server(Service(), server)
    server.add_insecure_port('0.0.0.0:10002')
    server.start()
    print('OPEN')
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
