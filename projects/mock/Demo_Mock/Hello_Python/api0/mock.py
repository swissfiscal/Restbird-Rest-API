# encoding: utf-8

def HandleRequest(req, state) :
    print(req.path)
    req.send_response(200)
    req.send_header('Content-type', 'text/html')
    req.end_headers()
    req.wfile.write("\nwelcome to advanced Mock Handler\n".encode())
    req.wfile.flush()
