import socket
import sys

def send_http_request(server_host, server_port, filename):
    try:
        client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        
        client_socket.connect((server_host, int(server_port)))
        
        request = f"GET /{filename} HTTP/1.1\r\nHost: {server_host}\r\nConnection: close\r\n\r\n"
        client_socket.sendall(request.encode())
        
        response = b""
        while True:
            data = client_socket.recv(1024)
            if not data:
                break
            response += data

        client_socket.close()

        response = response.decode()
        header, _, body = response.partition("\r\n\r\n")

        if "200 OK" in header:
            print(f"Found {filename}:")
        elif "404 Not Found" in header:
            print(f"Not Found {filename}:")
        elif "400 Bad Request" in header:
            print(f"Bad Request : ")

        print(response)

    except ConnectionRefusedError:
        print(f"Server Not Responding:")
        print(f"Error: Unable to connect to the server at {server_host}:{server_port}.")
        print(f"Please check if the server is running and the port is correct.")

    except Exception as e:
        print(f"HTTP request failed: {e}")


if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Format: py webClient.py <ip> <port> <file_name>")
        sys.exit(1)

    send_http_request(sys.argv[1], sys.argv[2], sys.argv[3])
