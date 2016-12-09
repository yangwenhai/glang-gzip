#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "gzip.h"
#define BUF_SIZE 65535
int main()
{
        char *data = "kjdalkfjdflkjdlkfjdklfjdlkfjlkdjflkdjflddajfkdjfkdfaskf;ldsfk;ldakf;ldskfl;dskf;ld";
        uLong ndata = strlen(data);
        char zdata[BUF_SIZE];
        uLong nzdata = BUF_SIZE;
        char  odata[BUF_SIZE];
        uLong nodata = BUF_SIZE;

        memset(zdata, 0, BUF_SIZE);
        int retlen=GzipCompress(data, ndata, zdata);
        fprintf(stdout, "retlen:%d \n", retlen);
        if( retlen >  0)
        {
                fprintf(stdout, "compress ok!nzdata:%lu %d\n", nzdata, retlen);
                memset(odata, 0, BUF_SIZE);
                int retlen2=GzipUncompress(zdata, ndata, odata);
                if( retlen2> 0)
                {
                  fprintf(stdout, "uncompress ok! %d %s\n", retlen2, odata);
                }
        }else{
             fprintf(stdout, "compress error!");
        }
}
