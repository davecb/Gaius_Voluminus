package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

/*
 * voluminous -- generate a test suite via exhaustion
 *	It's full name is Gaius Voluminous Fartus, after the romans in Asterix
 *	The icon is actually of Langelus, a blowhard centurion
 */

func main() {
	fd := Open("/dev/null", unix.O_APPEND)
	fmt.Printf("Open returned %d\n", fd)

}

// Open -- an open-like function to test
func Open(filename string, flags int32) int {
	// if the path points to readable file, open and return it using
	// os.Open.  This is just a stalking-horse for the testing
	return 0
}

// generate a test suite form path, flags and optional mode
// path == (
// 	nil,
//	a path to nowhere,
//	one to a readable file,
//	an unreadable file or directory in the path
//  an unwritable file or dir
//  a symlink loop
//  path argument exceeds {PATH_MAX} or a pathname component is longer than {NAME_MAX}.
//  path contains a non-directory
// )
// flags == (
// 64-bit int has 8 bytes, we use 7
// rightmost bits, octal 1
// O_RDONLY             00
// O_WRONLY             01
// O_RDWR               02
// O_ACCMODE            03 this is a mask for two bits

// third byte, octal  3
// O_CREAT           0100
// O_EXCL            0200
// O_NOCTTY          0400

// fourth byte       4
// O_TRUNC          01000
// O_APPEND         02000
// O_NONBLOCK       04000
// O_NDELAY    O_NONBLOCK

// fifth byte       5
// O_DSYNC	       010000
// O_ASYNC         020000
// O_DIRECT        040000

// seventh and fifth bytes
//                7 5
// O_SYNC        04010000 fifth sets dsync
// O_FSYNC         O_SYNC
// O_RSYNC         O_SYNC

// sixth bytes     6
// O_LARGEFILE    0100000
// O_DIRECTORY    0200000
// O_NOFOLLOW     0400000

// seventh byte   7
// O_CLOEXEC     02000000
// O_NOATIME     01000000

// Eighth byte   8
// O_PATH       010000000
// O_TMPFILE   (020000000 | O_DIRECTORY) sixth sets directory
// )

//
// mode == (S_IRWXU  and friends,
// S_IREAD        0400    /* Read by owner. */
// S_IWRITE       0200    /* Write by owner.  */
// S_IEXEC        0100    /* Execute by owner.  */
// extends to S_IRGRP | S_IROTH
// S_IRGRP         040    /* Read by group.  */
// S_IWRGRP        020    /* Write by group.  */
// S_IEXGRP        010    /* Execute by group.  */
// S_IROTH          04    /* Read by other.     0b100 */
// S_IWROT          02    /* Write by other.    0b010 */
// S_IEXOT          01    /* Execute by other.  0b001 */
// // and then
// S_ISUID       04000   /* Set user ID on execution.  */
// S_ISGID       02000   /* Set group ID on execution.  */
// S_ISVTX       01000   /* Save swapped text after use (sticky).  */
// S_IREAD        0400    /* Read by owner.  */
// S_IWRITE       0200    /* Write by owner.  */
// S_IEXEC        0100    /* Execute by owner.  */
// )

// unreachables!
// [EMFILE] {OPEN_MAX} file descriptors are currently open in the calling process.
// [ENFILE] The maximum allowable number of files is currently open in the system.
// [ENOSR] The path argument names a STREAMS-based file and the system is unable to allocate a STREAM. [Option End]
// [ENOSPC] The directory or file system that would contain the new file cannot be expanded, the file does not exist, and O_CREAT is specified.
// [ENXIO] O_NONBLOCK is set, the named file is a FIFO, O_WRONLY is set, and no process has the file open for reading.
// [ENXIO] The named file is a character special or block special file, and the device associated with this special file does not exist.
// [EOVERFLOW] The named file is a regular file and the size of the file cannot be represented correctly in an object of type off_t.
// [EAGAIN] The path argument names the slave side of a pseudo-terminal device that is locked. [Option End]
