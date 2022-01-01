package builder

import (
	"os"
	"path/filepath"

	"github.com/tinygo-org/tinygo/goenv"
)

// Picolibc is a C library for bare metal embedded devices. It was originally
// based on newlib.
var Picolibc = Library{
	name: "picolibc",
	makeHeaders: func(target, includeDir string) error {
		f, err := os.Create(filepath.Join(includeDir, "picolibc.h"))
		if err != nil {
			return err
		}
		return f.Close()
	},
	cflags: func(target, headerPath string) []string {
		picolibcDir := filepath.Join(goenv.Get("TINYGOROOT"), "lib/picolibc/newlib/libc")
		return []string{
			"-Werror",
			"-Wall",
			"-std=gnu11",
			"-D_COMPILING_NEWLIB",
			"-DHAVE_ALIAS_ATTRIBUTE",
			"-DTINY_STDIO",
			"-nostdlibinc",
			"-Xclang", "-internal-isystem", "-Xclang", picolibcDir + "/include",
			"-I" + picolibcDir + "/tinystdio",
			"-I" + headerPath,
		}
	},
	sourceDir: func() string { return filepath.Join(goenv.Get("TINYGOROOT"), "lib/picolibc/newlib/libc") },
	librarySources: func(target string) []string {
		return picolibcSources
	},
}

var picolibcSources = []string{
	"../../../picolibc-stdio.c",

	"tinystdio/asprintf.c",
	"tinystdio/atod_engine.c",
	"tinystdio/atod_ryu.c",
	"tinystdio/atof_engine.c",
	"tinystdio/atof_ryu.c",
	//"tinystdio/atold_engine.c", // have_long_double and not long_double_equals_double
	"tinystdio/clearerr.c",
	"tinystdio/compare_exchange.c",
	"tinystdio/dtoa_data.c",
	"tinystdio/dtoa_engine.c",
	"tinystdio/dtoa_ryu.c",
	"tinystdio/ecvtbuf.c",
	"tinystdio/ecvt.c",
	"tinystdio/ecvt_data.c",
	"tinystdio/ecvtfbuf.c",
	"tinystdio/ecvtf.c",
	"tinystdio/ecvtf_data.c",
	"tinystdio/exchange.c",
	//"tinystdio/fclose.c", // posix-io
	"tinystdio/fcvtbuf.c",
	"tinystdio/fcvt.c",
	"tinystdio/fcvtfbuf.c",
	"tinystdio/fcvtf.c",
	"tinystdio/fdevopen.c",
	//"tinystdio/fdopen.c", // posix-io
	"tinystdio/feof.c",
	"tinystdio/ferror.c",
	"tinystdio/fflush.c",
	"tinystdio/fgetc.c",
	"tinystdio/fgets.c",
	"tinystdio/fileno.c",
	"tinystdio/filestrget.c",
	"tinystdio/filestrputalloc.c",
	"tinystdio/filestrput.c",
	//"tinystdio/fopen.c", // posix-io
	"tinystdio/fprintf.c",
	"tinystdio/fputc.c",
	"tinystdio/fputs.c",
	"tinystdio/fread.c",
	"tinystdio/fscanf.c",
	"tinystdio/fseek.c",
	"tinystdio/ftell.c",
	"tinystdio/ftoa_data.c",
	"tinystdio/ftoa_engine.c",
	"tinystdio/ftoa_ryu.c",
	"tinystdio/fwrite.c",
	"tinystdio/gcvtbuf.c",
	"tinystdio/gcvt.c",
	"tinystdio/gcvtfbuf.c",
	"tinystdio/gcvtf.c",
	"tinystdio/getchar.c",
	"tinystdio/gets.c",
	"tinystdio/matchcaseprefix.c",
	"tinystdio/perror.c",
	//"tinystdio/posixiob.c", // posix-io
	//"tinystdio/posixio.c", // posix-io
	"tinystdio/printf.c",
	"tinystdio/putchar.c",
	"tinystdio/puts.c",
	"tinystdio/ryu_divpow2.c",
	"tinystdio/ryu_log10.c",
	"tinystdio/ryu_log2pow5.c",
	"tinystdio/ryu_pow5bits.c",
	"tinystdio/ryu_table.c",
	"tinystdio/ryu_umul128.c",
	"tinystdio/scanf.c",
	"tinystdio/setbuf.c",
	"tinystdio/setvbuf.c",
	//"tinystdio/sflags.c", // posix-io
	"tinystdio/snprintf.c",
	"tinystdio/snprintfd.c",
	"tinystdio/snprintff.c",
	"tinystdio/sprintf.c",
	"tinystdio/sprintfd.c",
	"tinystdio/sprintff.c",
	"tinystdio/sscanf.c",
	"tinystdio/strfromd.c",
	"tinystdio/strfromf.c",
	"tinystdio/strtod.c",
	"tinystdio/strtod_l.c",
	"tinystdio/strtof.c",
	//"tinystdio/strtold.c",   // have_long_double and not long_double_equals_double
	//"tinystdio/strtold_l.c", // have_long_double and not long_double_equals_double
	"tinystdio/ungetc.c",
	"tinystdio/vasprintf.c",
	"tinystdio/vfiprintf.c",
	"tinystdio/vfiscanf.c",
	"tinystdio/vfprintf.c",
	"tinystdio/vfprintff.c",
	"tinystdio/vfscanf.c",
	"tinystdio/vfscanff.c",
	"tinystdio/vprintf.c",
	"tinystdio/vscanf.c",
	"tinystdio/vsnprintf.c",
	"tinystdio/vsprintf.c",
	"tinystdio/vsscanf.c",

	"string/bcmp.c",
	"string/bcopy.c",
	"string/bzero.c",
	"string/explicit_bzero.c",
	"string/ffsl.c",
	"string/ffsll.c",
	"string/fls.c",
	"string/flsl.c",
	"string/flsll.c",
	"string/gnu_basename.c",
	"string/index.c",
	"string/memccpy.c",
	"string/memchr.c",
	"string/memcmp.c",
	"string/memcpy.c",
	"string/memmem.c",
	"string/memmove.c",
	"string/mempcpy.c",
	"string/memrchr.c",
	"string/memset.c",
	"string/rawmemchr.c",
	"string/rindex.c",
	"string/stpcpy.c",
	"string/stpncpy.c",
	"string/strcasecmp.c",
	"string/strcasecmp_l.c",
	"string/strcasestr.c",
	"string/strcat.c",
	"string/strchr.c",
	"string/strchrnul.c",
	"string/strcmp.c",
	"string/strcoll.c",
	"string/strcoll_l.c",
	"string/strcpy.c",
	"string/strcspn.c",
	"string/strdup.c",
	"string/strerror.c",
	"string/strerror_r.c",
	"string/strlcat.c",
	"string/strlcpy.c",
	"string/strlen.c",
	"string/strlwr.c",
	"string/strncasecmp.c",
	"string/strncasecmp_l.c",
	"string/strncat.c",
	"string/strncmp.c",
	"string/strncpy.c",
	"string/strndup.c",
	"string/strnlen.c",
	"string/strnstr.c",
	"string/strpbrk.c",
	"string/strrchr.c",
	"string/strsep.c",
	"string/strsignal.c",
	"string/strspn.c",
	"string/strstr.c",
	"string/strtok.c",
	"string/strtok_r.c",
	"string/strupr.c",
	"string/strverscmp.c",
	"string/strxfrm.c",
	"string/strxfrm_l.c",
	"string/swab.c",
	"string/timingsafe_bcmp.c",
	"string/timingsafe_memcmp.c",
	"string/u_strerr.c",
	"string/wcpcpy.c",
	"string/wcpncpy.c",
	"string/wcscasecmp.c",
	"string/wcscasecmp_l.c",
	"string/wcscat.c",
	"string/wcschr.c",
	"string/wcscmp.c",
	"string/wcscoll.c",
	"string/wcscoll_l.c",
	"string/wcscpy.c",
	"string/wcscspn.c",
	"string/wcsdup.c",
	"string/wcslcat.c",
	"string/wcslcpy.c",
	"string/wcslen.c",
	"string/wcsncasecmp.c",
	"string/wcsncasecmp_l.c",
	"string/wcsncat.c",
	"string/wcsncmp.c",
	"string/wcsncpy.c",
	"string/wcsnlen.c",
	"string/wcspbrk.c",
	"string/wcsrchr.c",
	"string/wcsspn.c",
	"string/wcsstr.c",
	"string/wcstok.c",
	"string/wcswidth.c",
	"string/wcsxfrm.c",
	"string/wcsxfrm_l.c",
	"string/wcwidth.c",
	"string/wmemchr.c",
	"string/wmemcmp.c",
	"string/wmemcpy.c",
	"string/wmemmove.c",
	"string/wmempcpy.c",
	"string/wmemset.c",
	"string/xpg_strerror_r.c",
}
