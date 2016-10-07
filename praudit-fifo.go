package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include <bsm/libbsm.h>

#define AU_HEADER_32_TOKEN 0x14

extern char *optarg;
extern int optind, optopt, opterr,optreset;

static char *del = ","; // default delimiter
static int oneline = 1;
static int raw = 0;
static int shortfrm = 1;
static int partial = 0;

static void usage()
{
	printf("Usage: praudit [-lrs] [-ddel] [filenames]\n");
	exit(1);
}

//
// token printing for each token type
//
static int print_tokens(FILE *fp, FILE *out)
{
	u_char *buf;
	tokenstr_t tok;
	int reclen;
   	int bytesread;

	// allow tail -f | praudit to work
        if (partial) {
            u_char type = 0;
            // record must begin with a header token
            do {
                type = fgetc(fp);
            } while(type != AU_HEADER_32_TOKEN);
            ungetc(type, fp);
        }

	while((reclen = au_read_rec(fp, &buf)) != -1) {

		bytesread = 0;

		while (bytesread < reclen) {

			if(-1 == au_fetch_tok(&tok, buf + bytesread, reclen - bytesread)) {
				// is this an incomplete record ?
				break;
			}

			au_print_tok(out, &tok, del, raw, shortfrm);
			bytesread += tok.len;

			if(oneline) {
				fprintf(out, "%s", del);
			}
			else {
				fprintf(out, "\n");
			}
		}

		free(buf);

		if(oneline) {
			fprintf(out, "\n");
		}
	}

	return 0;
}
#cgo LDFLAGS: -L/usr/lib
#cgo LDFLAGS: -lbsm
*/
import "C"

// A BSMMessage actually maps quite cleanly to a Netlink audit message.
type BSMMessage struct {
	Header struct {
		Len   uint32
		Type  uint16
		Flags uint16
		Seq   uint32
		Pid   uint32
	}
	Data []byte
}

func main() {
	// Manual parse implementation

	/*
		// Unix FIFO implementation
		outpath := "/tmp/go-audit.pipe"
		syscall.Mkfifo(outpath, 0777)

		file := C.CString(os.Args[1])
		outpipe := C.CString(outpath)
		read := C.CString("r")
		write := C.CString("w")

		defer func() {
			C.free(unsafe.Pointer(file))
			C.free(unsafe.Pointer(outpipe))
			C.free(unsafe.Pointer(read))
			C.free(unsafe.Pointer(write))
		}()

		fp := C.fopen(file, read)
		outp := C.fopen(outpipe, write)
		go C.print_tokens(fp, outp)

		t, _ := tail.TailFile(outpath, tail.Config{Follow: true})
		for line := range t.Lines {
			fmt.Println(line.Text)
		}*/

	/*
		// StdoutPipe implementation
		auditParserCmd := "/usr/sbin/praudit"
		cmd := exec.Command(auditParserCmd, "-l", "/dev/auditpipe")
		stdout, err := cmd.StdoutPipe()
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		go func() {
			s := bufio.NewScanner(stderr)
			for s.Scan() {
				fmt.Fprintf(os.Stderr, s.Text())
			}
		}()

		go func() {
			s := bufio.NewScanner(stdout)
			for s.Scan() {
				fmt.Println(s.Text())
			}
		}()

		if err := cmd.Run(); err != nil {
			fmt.Println("ERROR:", err)
		}*/
}
