#include "netstat.h"

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char **array;
char *netres;

int line() {
    FILE *fp = fopen("/proc/net/dev", "r");
    int ch = 0, lines = 0;
    while(!feof(fp)) {
        ch = fgetc(fp);
        if(ch == '\n') {
            lines++;
        }   
    }
    fclose(fp);
    return lines;
}

int arrLen;

void init() {
    arrLen = line() - 2;
    array = (char**)malloc(arrLen*sizeof(int*));
    for(int i = 0; i < arrLen; i++){  
        array[i]=(char*)malloc(200*sizeof(int));  
    }
    netres = (char*)malloc(200*arrLen*sizeof(int));
}

void netstat() {
    FILE *fp = fopen("/proc/net/dev", "r");
    char buf[200], ifname[20];
    unsigned long int r_bytes, t_bytes, r_packets, t_packets;

    // skip first two lines
    for (int i = 0; i < 2; i++) {
        fgets(buf, 200, fp);
    }

    int l = 0;
    while (fgets(buf, 200, fp)) {
        sscanf(buf, "%[^:]: %lu %lu %*lu %*lu %*lu %*lu %*lu %*lu %lu %lu",
               ifname, &r_bytes, &r_packets, &t_bytes, &t_packets);
        if (l != arrLen - 1) {
            sprintf(array[l],"{\"Name\":\"%s\",\"rbytes\":%lu,\"rpackets\":%lu,\"tbytes\":%lu,\"tpackets\":%lu},",
               ifname, r_bytes, r_packets, t_bytes, t_packets);
        } else {
            sprintf(array[l],"{\"Name\":\"%s\",\"rbytes\":%lu,\"rpackets\":%lu,\"tbytes\":%lu,\"tpackets\":%lu}",
               ifname, r_bytes, r_packets, t_bytes, t_packets);
        }
        l++;
    }
    
    for (int i = 0; i < arrLen; i++) {
        if (i == 0) {
            strcpy(netres, "{");
            strcat(netres, array[i]);
        } else {
            if (i == arrLen - 1) {
                strcat(netres, array[i]);
                strcat(netres, "}");
            } else {
                strcat(netres, array[i]);
            }
        }
    }
    fclose(fp);
    printf("%s", netres);
}