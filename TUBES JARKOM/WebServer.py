import socket

SERVER_HOST = ""
SERVER_PORT = 8080
BUFFER_SIZE = 1024


def start_server():
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    server_socket.bind((SERVER_HOST, SERVER_PORT))

    server_socket.listen(5)

    while True:
        print("Ready to serve...")

        client_socket, client_address = server_socket.accept()
        print(f"Get Connection from {client_address}")

        try:
            request = client_socket.recv(BUFFER_SIZE).decode()
            print("Request : ")
            print(request)
            
            request_lines = request.split("\n")
        
            if request_lines:
                parts = request_lines[0].split()
                if len(parts) >= 2:
                    filename = parts[1].lstrip('/')
                    
                    if filename == "":
                        filename = "index.html"
                        
                    fin = open(filename)

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
    server_socket.close()


if __name__ == "__main__":
    start_server()
