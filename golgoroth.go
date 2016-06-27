package golgoroth

import (
	"math/rand"
	"strings"
	"time"
)

// Generate creates a Cthulhu-inspired name
func Generate() string {
	var cthulhu string
	for cthulhu == "" || strings.Contains(cthulhu, "''") {
		cthulhu = createName()
	}
	return strings.Title(cthulhu)
}

func createName() string {
	var bits []string
	rand.Seed(time.Now().Unix())
	parts := parts()

	upper := rand.Intn(5) + 2

	for i := 0; i < upper; i++ {
		bits = append(bits, parts[rand.Intn(len(parts))])
	}
	for bits[0] == "'" {
		bits[0] = parts[rand.Intn(len(parts))]
	}
	for bits[len(bits)-1] == "'" {
		bits[len(bits)-1] = parts[rand.Intn(len(parts))]
	}

	return strings.Join(bits, "")
}

func parts() []string {
	return []string{
		"'", "te", "co", "ra", "ll", "ta", "la", "to", "po", "be", "ci", "vi", "no", "ck",
		"'", "wa", "tu", "do", "rr", "fu", "mu", "ry", "ps", "sw", "tc", "dd", "ys", "nf",
		"'", "wr", "ja", "lf", "hs", "rh", "zi", "cc", "dn", "hr", "sr", "sg", "nm", "fs",
		"'", "df", "hw", "tg", "pn", "dp", "lh", "kr", "nz", "gt", "bv", "kf", "km", "dv",
		"'", "vr", "wp", "cn", "cq", "sv", "bn", "rx", "md", "pc", "xw", "bf", "xl", "xm",
		"'", "fp", "gj", "kj", "wu", "zp", "on", "ar", "ac", "em", "iv", "ir", "if", "ip",
		"'", "od", "ob", "up", "ud", "ew", "eq", "ox", "ax", "iz", "ik", "iq", "aq", "uv",
		"'", "oq", "ez", "ij", "ih", "cm", "sg", "xh", "ts",
		"'", "ae", "ai", "ao", "au", "ay",
		"'", "ei", "eo", "eu", "ey", "ia",
		"'", "ii", "io", "iu", "iy", "oa",
		"'", "oe", "oi", "oo", "oy", "ua",
		"'", "ue", "ui", "uo", "uu", "uy",
		"'", "ya", "ye", "yi", "yu", "yy",
		"'",
	}
}
