package main


/*================================= imports ======================*/

import (
	"io"
	"os/exec"
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

/*================================= types =========================*/

type selpg_args struct{
	start_page int
	end_page int
	in_filename string
	page_len int /* default value, can be overriden by "-l number" on command line */
	page_type int /* 'l' for lines-delimited, 'f' for form-feed-delimited */
					/* default is 'l' */
	print_dest string
}

/*================================= globals =======================*/

var progname string	/* program name, for error messages */

var inputS = flag.IntP("start_page", "s", -1, "Input start page.")
var inputE = flag.IntP("end_page", "e", -1, "Input end page.")
/*这两个分别是选择两种类型的输入文本:页行数固定和不固定(遇到\f换页)*/
var inputL = flag.IntP("fixed_rows", "l", 72, "Fixed row num. Input page rows.")
var inputF = flag.BoolP("unsure_rows", "f", false, "Unsure row num.")
var inputD = flag.StringP("print_dest", "d", "", "Input print destination.")

/*================================= main()=== =====================*/

func main() {
	/* save name by which program is invoked, for error messages */
	progname = os.Args[0]

	sp_args := selpg_args{
		start_page: -1,
		end_page: -1,
		in_filename: "",
		page_len: 72,
		page_type: 'l',
		print_dest: "",
	}

	process_args(&sp_args)
	process_input(&sp_args)
}

/*================================= process_args() ================*/

func process_args(sp_args *selpg_args) {
	flag.Parse()
	/*handle -s and -e*/
	if *inputS == -1 || *inputE == -1 {
		fmt.Fprintf(os.Stderr, "%s: you must input startPage and endPage.\n", progname)
		os.Exit(1)
	}
	if *inputS < 1 || *inputE < 1 {
		fmt.Fprintf(os.Stderr, "%s: you must input valid startPage and endPage(>=0).\n", progname)
		os.Exit(2)
	}
	if *inputS > *inputE {
		fmt.Fprintf(os.Stderr, "%s: startPage must be greater than endPage.\n", progname)
		os.Exit(3)
	}
	sp_args.start_page = *inputS
	sp_args.end_page = *inputE

	/*handle -l and -f*/
	if (*inputL < 1) {
		fmt.Fprintf(os.Stderr, "%s: rows of page must be greater than 0.\n", progname)
		os.Exit(4)
	}
	sp_args.page_len = *inputL
	if *inputF == true {
		sp_args.page_type = 'f'
	}
	
	/*handle -d*/
	sp_args.print_dest = *inputD

	/*handle input filename*/
	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "%s: there can only be one destination\n", progname)
		os.Exit(5)
	}
	if flag.NArg() == 1 {
		if _, err := os.Stat(flag.Arg(0)); err != nil {
			fmt.Fprintf(os.Stderr, "%s: file doesn't exist, check your input!\n", progname)
			os.Exit(6)
		}
		
		if _, err := os.Open(flag.Arg(0)); err != nil {
			fmt.Fprintf(os.Stderr, "%s: can't open this file.\n", progname)
			os.Exit(7)
		}
		sp_args.in_filename = flag.Arg(0)
	}

	if (sp_args.page_type != 'l' && sp_args.page_type != 'f') {
		fmt.Fprintf(os.Stderr, "%s: there are only 2 types of rows of page!\n", progname)
		os.Exit(8)
	}
}

func process_input(sp_args *selpg_args) {
	var fin *bufio.Reader /* input stream */
	var fout *bufio.Writer /* output stream */
	var err error
	var file *os.File
	var cmd *exec.Cmd
	var pipe io.WriteCloser
	
	/* set the input source */
	if (sp_args.in_filename == "") {
		fin = bufio.NewReader(os.Stdin)
	} else {
		file, err = os.Open(sp_args.in_filename)
		/*if err != nil {
			fmt.Fprintln(os.Stderr, "%s: failed to open file.\n", progname)
			os.Exit(9)
		}*/
		fin = bufio.NewReader(file)
		defer file.Close()
	}

	/* set the output destination */
	if sp_args.print_dest == "" {
		fout = bufio.NewWriter(os.Stdout)
	} else {
		cmd = exec.Command("lp", "-d", sp_args.print_dest)
		pipe, err = cmd.StdinPipe()
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: no printer\n", progname)
		}
		defer pipe.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: failed to open pipe to %s.\n", progname, sp_args.print_dest)
			os.Exit(10)
		}
	}


	/* begin one of two main loops based on page type */
	var page_ctr int
	if (sp_args.page_type == 'l') {
		line_ctr := 0
		page_ctr = 1
		for {
			line, err := fin.ReadBytes('\n')
			if err == io.EOF || page_ctr > sp_args.end_page {
				break
			}
			line_ctr++
			if line_ctr > sp_args.page_len {
				page_ctr++
				line_ctr = 1
			}
			if page_ctr >= sp_args.start_page && page_ctr <= sp_args.end_page {
				if sp_args.print_dest == "" {
					fout.Write(line)
					fout.Flush()
				} else {
					_, err := io.WriteString(pipe, string(line))
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
				}
			}
		}
	} else {
		page_ctr = 1
		for {
			line, err := fin.ReadBytes('\f')
			if err == io.EOF || page_ctr > sp_args.end_page {
				break
			}
			page_ctr++
			if page_ctr >= sp_args.start_page && page_ctr <= sp_args.end_page {
				if sp_args.print_dest == "" {
					fout.Write(line)
					fout.Flush()
				} else {
					_, err := io.WriteString(pipe, string(line))
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
				}
			}
		}
	}

	if sp_args.print_dest != "" {
		pipe.Close()
		stderr, _ := cmd.CombinedOutput()
		fmt.Fprintln(os.Stderr, string(stderr))
	}

	/* end main loop */
	if page_ctr < sp_args.start_page {
		fmt.Fprintf(os.Stderr, 
			"%s: start_page (%d) greater than total pages (%d)," + 
			" no output written\n", progname, sp_args.start_page, page_ctr)
	}
	if page_ctr < sp_args.end_page {
		fmt.Fprintf(os.Stderr, 
			"%s: end_page (%d) greater than total pages (%d)," + 
			" less output than expected\n", progname, sp_args.end_page, page_ctr)
	}
}