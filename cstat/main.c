#include "httpd/httpd.h"
#include "netstat/netstat.h"
#include "load/load.h"

int main(int c, char** v)
{
    serve_forever("8080");
    return 0;
}

void route()
{
    ROUTE_START()

    ROUTE_GET("/") {
        printf("HTTP/1.1 200 OK\r\n\r\n");
        printf("Whoa there!");
    }

    ROUTE_GET("/network") {
        printf("HTTP/1.1 200 OK\r\n\r\n");
        //printf("Hello! You are using %s", request_header("User-Agent"));
        netstat();
    }

    ROUTE_GET("/load") {
        printf("HTTP/1.1 200 OK\r\n\r\n");
        load();
    }
    /*
    ROUTE_POST("/") {
        printf("HTTP/1.1 200 OK\r\n\r\n");
        printf("Wow, seems that you POSTed %d bytes. \r\n", payload_size);
        printf("Fetch the data using `payload` variable.");
    }
    */
    ROUTE_END()
}