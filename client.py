import sys
from datetime import datetime
import grpc

from py_client.proto import calc_pb2
from py_client.proto import calc_pb2_grpc


if __name__ == '__main__':
    address = 'localhost:8080'
    if len(sys.argv) == 2:
        address = sys.argv[1]
    print("Start service: " + str(datetime.now()), "on", address, "\n")
    try:
        channel = grpc.insecure_channel(address)
        stub = calc_pb2_grpc.CalcServiceStub(channel)

        print("Performing Operations on CalcService:")

        request = calc_pb2.Request(a=10, b=20)
        response = stub.Add(request)
        print(request, "Add:", response)

        request = calc_pb2.Request(a=100, b=20)
        response = stub.Sub(request)
        print(request, "Sub:", response)

        request = calc_pb2.Request(a=10, b=20)
        response = stub.Mul(request)
        print(request, "Mul:", response)

        request = calc_pb2.Request(a=100, b=20)
        response = stub.Div(request)
        print(request, "Div:", response)

    except Exception as e:
        print(e)
