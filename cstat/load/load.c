#include "load.h"

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void load() {
    double loadavg[3];
    getloadavg(&loadavg[0], 3);
    char loadres[40];
    sprintf(loadres, "{\"loadavg1\"=\"%.2f\",\"loadavg5\"=\"%.2f\",\"loadavg15\"=\"%.2f\"}", loadavg[0], loadavg[1], loadavg[2]);
    printf("%s", loadres);
}