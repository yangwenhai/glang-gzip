

#pragma once
#ifdef __cplusplus
extern "C" {
#endif

#include <zlib.h>
    
/* Compress gzip data */
long GzipCompress(char *data, uLong ndata, char *zdata);
long GzipUncompress(char *zdata, uLong nzdata, char *data);
 
#ifdef __cplusplus    
 }
 #endif