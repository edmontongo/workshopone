# very rough code to obtain line counts for .go files in primarily
# the standard library
#
# function get_file_list() {
#     find "$@" -iname "*.go" | grep -Ev 'vendor|windows|runtime|cmd|_(test|32bit|386|64bit|abbrs_windows|amd64|amd64p32|amd64x|android|arm|arm64|arm_linux|arm_test|asm|bsd|bsd_test|bsdvar|darwin|darwin_386|darwin_amd64|darwin_arm|darwin_arm64|darwin_arm_gen|darwin_armx|darwin_test|dragonfly|dragonfly_amd64|err_android|freebsd|freebsd_32bit|freebsd_386|freebsd_64bit|freebsd_amd64|freebsd_arm|freebsd_test|ios|linux|linux_32bit|linux_386|linux_amd64|linux_amd64_test|linux_arm|linux_arm64|linux_arm_test|linux_generic|linux_mips64|linux_mips64le|linux_mips64x|linux_noauxv|linux_ppc64|linux_ppc64le|linux_ppc64x|linux_s390x|linux_test|mips64|mips64le|mips64x|nacl_386|nacl_amd64p32|nacl_arm|netbsd|netbsd_386|netbsd_amd64|netbsd_arm|nonppc64x|nonunix_test|notbsd|notunix|openbsd|openbsd_386|openbsd_amd64|openbsd_arm|plan9|plan9_386|plan9_amd64|plan9_arm|plan9_test|posix|posix_test|ppc64|ppc64le|ppc64x|s390x|solaris|solaris_amd64|solaris_test|unix|unix_test|windows|windows_386|windows_amd64|windows_test|x86|nacl).go'
# }
#
# function get_count() {
#     get_file_list "$@" | xargs wc -l | tail -n1 | awk '{print $1}'
# }
#
# git checkout go1; get_count src/pkg
# git checkout go1.1; get_count src/pkg
# git checkout go1.2; get_count src/pkg
# git checkout go1.3; get_count src/pkg
# git checkout go1.4; get_count src
# git checkout go1.5; get_count src
# git checkout go1.6; get_count src
# git checkout go1.7; get_count src
# git checkout go1.8; get_count src

pdf("gocode.pdf", width=6, height=3)

d <- data.frame(
    release=c("go1", "go1.1", "go1.2", "go1.3", "go1.4", "go1.5", "go1.6", "go1.7", "go1.8"),
    lines=c(134764, 148589, 157755, 163431, 167855, 195370, 209755, 217034, 237592)
)

par(mar=c(2.3, 2.3, 0.6, 0.3))
par(mgp=c(1.3, 0.3, 0))
par(tcl=-0.2)

plot(d$lines/1000, type="l", ylab="Lines (1000s)", xaxt="n", xlab="Release")
axis(1, at=1:length(d$release), labels=d$release)

dev.off()
