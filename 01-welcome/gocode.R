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
