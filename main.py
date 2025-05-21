from concurrent import futures
import logging

import grpc

import py_protos.hi_pb2 as hi__pb2
import py_protos.hi_pb2_grpc as hi__pb2_grpc

class Greeter(hi__pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        return hi__pb2.HelloReply(message=f'Hello from Python, {request.name}!')

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    hi__pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Python gRPC server started on port 50051")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()