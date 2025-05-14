import socket
import threading

SERVER_HOST = ""
SERVER_PORT = 8080
BUFFER_SIZE = 1024


def handle_client(client_socket, client_address):
    try:
        request = client_socket.recv(BUFFER_SIZE).decode()
        print(f"Request from {client_address} : ")
        print(request)

        request_lines = request.split("\n")
        
        if request_lines:
            parts = request_lines[0].split()
            if len(parts) >= 2:
                filename = parts[1].lstrip('/')
                
                if filename == "":
                    filename = "index.html"
                
                fin = open(filename, encoding='utf-8', errors='replace')

                content = fin.read()

                response = f"HTTP/1.1 200 OK\r\n\r\n" + content
            else:
                print("Request Lines Not Complete")
                content = "<html> <head> </head> <body> <h1> 400 Bad Request </h1> </body> </html> \r\n"
                response = f"HTTP/1.1 400 Bad Request\r\n\r\n" + content
        else:
            print("Request Lines Not Found")
            content = "<html> <head> </head> <body> <h1> 400 Bad Request </h1> </body> </html> \r\n"
            response = f"HTTP/1.1 400 Bad Request\r\n\r\n" + content

    except IOError:
        content = "<html> <head> </head> <body> <h1> 404 Not Found </h1> </body> </html> \r\n"

        response = f"HTTP/1.1 404 Not Found\r\n\r\n" + content
    
    client_socket.sendall(response.encode())
    client_socket.close()


def start_server():
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    server_socket.bind((SERVER_HOST, SERVER_PORT))

    server_socket.listen(5)

    print("Multithread server is ready to serve...")

    while True:

        client_socket, client_address = server_socket.accept()
        print(f"Get Connection from {client_address}")
        threading.Thread(target=handle_client, daemon=True,
                         args=(client_socket, client_address,)).start()

    server_socket.close()


if __name__ == "__main__":
    start_server()
