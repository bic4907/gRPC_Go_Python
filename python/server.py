import logging
import os
import uuid
import cv2 as cv
from concurrent import futures
from datetime import datetime

import grpc
import protobuf.media_pb2 as media_pb
import protobuf.media_pb2_grpc as media_grpc

Chunk_Dir = 'chunks'


class Service(media_grpc.ServiceServicer):

    def StreamVideo(self, iterator, context):
        for req in iterator:

            fName = Chunk_Dir + '/' + str(uuid.uuid4()) + '.mp4'
            f = open(fName, 'wb')
            f.write(req.Chunk)
            f.close()

            cap = cv.VideoCapture(fName)
            while True:
                # Capture frame-by-frame
                ret, frame = cap.read()
                # if frame is read correctly ret is True
                if not ret:
                    break

                # Display the resulting frame
                cv.imshow('frame', frame)
                cv.waitKey(1)
            cap.release()
            os.remove(fName)

        message = media_pb.ReceiveReply(Result="1")
        return message


    def SendMessage(self, request, context):
        print(request.Content)
        message = media_pb.RplMessage(Content=request.Content)

        return message





def serve():
    global Chunk_Dir

    if not os.path.exists(Chunk_Dir):
        os.makedirs(Chunk_Dir)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    media_grpc.add_ServiceServicer_to_server(Service(), server)
    server.add_insecure_port('0.0.0.0:10002')
    server.start()
    print('OPEN')
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
