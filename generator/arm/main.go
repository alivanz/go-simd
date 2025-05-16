package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/alivanz/go-simd/generator/scanner"
	"github.com/alivanz/go-simd/generator/types"
	"github.com/alivanz/go-simd/generator/utils"
	"github.com/alivanz/go-simd/generator/writer"
	"github.com/iancoleman/strcase"
)

func Source() ([]byte, error) {
	cmd := exec.Command("clang", "-E", "-")
	cmd.Stdin = bytes.NewBufferString(strings.Join(writer.Includes([]string{
		"arm_neon.h",
	}), "\n"))
	cmd.Stderr = os.Stderr
	return cmd.Output()
}

func main() {
	src, err := Source()
	if err != nil {
		log.Fatal(err)
	}
	// write raw
	if err := writer.WriteToFile("raw.h", func(w io.Writer) error {
		_, err := w.Write(src)
		return err
	}); err != nil {
		log.Fatal(err)
	}
	// scan
	result, err := scanner.Scan(src)
	if err != nil {
		log.Fatal(err)
	}
	// filter functions
	result.Functions = utils.Filter(result.Functions, func(fn types.Function) bool {
		if strings.HasPrefix(fn.Name, "vbf") {
			return false
		}
		if strings.Contains(fn.Name, "bf16") {
			return false
		}
		return true
	})
	// filter types
	mtype := make(map[string]bool)
	for _, fn := range result.Functions {
		if fn.Return != nil {
			mtype[fn.Return.Name] = true
		}
		for _, arg := range fn.Args {
			mtype[arg.Name] = true
		}
	}
	result.Types = utils.Filter(result.Types, func(t types.Type) bool {
		return mtype[t.Name]
	})
	// sort functions
	sort.Slice(result.Functions, func(i, j int) bool {
		g0, i0, _ := sortGroup(result.Functions[i].Name)
		g1, i1, _ := sortGroup(result.Functions[j].Name)
		if g0 != g1 {
			return g0 < g1
		}
		return i0 < i1
	})
	// sort types
	sort.Slice(result.Types, func(i, j int) bool {
		return result.Types[i].Name < result.Types[j].Name
	})
	// write types
	if err := writer.WriteToFile("types.go", func(w io.Writer) error {
		if err := writer.Package(w, "arm"); err != nil {
			return err
		}
		if err := writer.ImportC(w, func(w io.Writer) error {
			_, err := io.WriteString(w, strings.Join(writer.Includes([]string{
				"arm_neon.h",
			}), "\n"))
			return err
		}); err != nil {
			return err
		}
		if err := writer.Types(w, result.Types); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// patch intrinsics info
	intrins, err := GetIntrinsics()
	if err != nil {
		log.Fatal(err)
	}
	for i, fn := range result.Functions {
		if info := intrins.Find(fn.Name); info != nil {
			result.Functions[i].Comment = info.Description
		}
	}
	// write C
	if err := writer.WriteToFile("neon/functions.c", func(w io.Writer) error {
		if _, err := io.WriteString(w, "#include <arm_neon.h>\n\n"); err != nil {
			return err
		}
		for _, fn := range result.Functions {
			if fn.Blacklisted() {
				continue
			}
			if err := writer.RewriteC(w, fn); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// write functions
	if err := writer.WriteToFile("neon/functions.go", func(w io.Writer) error {
		if err := writer.Package(w, "neon"); err != nil {
			return err
		}
		if err := writer.Import(w, []string{
			"github.com/alivanz/go-simd/arm",
		}); err != nil {
			return err
		}
		if err := writer.ImportC(w, func(w io.Writer) error {
			if _, err := io.WriteString(w, "#include <arm_neon.h>"); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
		for _, fn := range result.Functions {
			if fn.Blacklisted() {
				continue
			}
			writer.DeclareFuncBypass(w, fn, "arm")
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// C loops
	var (
		loops = make(map[string]bool)
	)
	if err := writer.WriteToFile("neon/loops.c", func(w io.Writer) error {
		if _, err := io.WriteString(w, "#include <arm_neon.h>\n\n"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "#define save(dst, src) *dst = src\n"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "#define load(src) (*src)\n"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, `#define LOOP1(name, rtype, itype, f, set, load, rstep, istep) \
    void name(rtype *r, itype *v, int32_t n)                  \
    {                                                         \
        while (n >= rstep)                                    \
        {                                                     \
            set(r, f(load(v)));                               \
            r += rstep;                                       \
            n -= rstep;                                       \
            v += istep;                                       \
        }                                                     \
    }

`); err != nil {
			return err
		}
		for _, fn := range result.Functions {
			if fn.Blacklisted() {
				continue
			}
			if len(fn.Args) != 1 {
				continue
			}
			og, o0, o1 := parseType(fn.Return.Name)
			if og == "" {
				continue
			}
			ig, i0, i1 := parseType(fn.Args[0].Name)
			if ig == "" {
				continue
			}
			if o0 != i0 {
				continue
			}
			var rq, iq string
			if o0*o1 == 128 {
				rq = "q"
			}
			if i0*i1 == 128 {
				iq = "q"
			}
			if o1 == -1 {
				o1 = 1
			}
			if i1 == -1 {
				i1 = 1
			}
			group, _, suffix := sortGroup(fn.Name)
			rg, r0, _ := parseType(fn.Return.Name)
			if rg == "" {
				continue
			}
			io.WriteString(w,
				fmt.Sprintf(
					"LOOP1(%s, %s, %s, %s, %s, %s, %d, %d)\n",
					strcase.ToCamel(group+suffix+"N"),
					fmt.Sprintf("%s%d_t", rg, r0),
					fmt.Sprintf("%s%d_t", ig, i0),
					fn.Name,
					setter(fn.Return.Name, "save", fmt.Sprintf("vst1%s_%s%d", rq, typeShort[rg], r0)),
					setter(fn.Args[0].Name, "load", fmt.Sprintf("vld1%s_%s%d", iq, typeShort[ig], i0)),
					o1,
					i1,
				),
			)
			loops[fn.Name] = true
		}
		io.WriteString(w, "\n")
		if _, err := io.WriteString(w, `#define LOOP2(name, rtype, itype, f, set, load, rstep, istep) \
    void name(rtype *r, itype *v1, itype *v2, int32_t n)      \
    {                                                         \
        while (n >= rstep)                                    \
        {                                                     \
            set(r, f(load(v1), load(v2)));                    \
            r += rstep;                                       \
            n -= rstep;                                       \
            v1 += istep;                                      \
            v2 += istep;                                      \
        }                                                     \
    }

`); err != nil {
			return err
		}
		for _, fn := range result.Functions {
			if fn.Blacklisted() {
				continue
			}
			if len(fn.Args) != 2 {
				continue
			}
			if fn.Args[0].Name != fn.Args[1].Name {
				continue
			}
			og, o0, o1 := parseType(fn.Return.Name)
			if og == "" {
				continue
			}
			ig, i0, i1 := parseType(fn.Args[0].Name)
			if ig == "" {
				continue
			}
			if o0 != i0 {
				continue
			}
			var rq, iq string
			if o0*o1 == 128 {
				rq = "q"
			}
			if i0*i1 == 128 {
				iq = "q"
			}
			if o1 == -1 {
				o1 = 1
			}
			if i1 == -1 {
				i1 = 1
			}
			group, _, suffix := sortGroup(fn.Name)
			rg, r0, _ := parseType(fn.Return.Name)
			if rg == "" {
				continue
			}
			io.WriteString(w,
				fmt.Sprintf(
					"LOOP2(%s, %s, %s, %s, %s, %s, %d, %d)\n",
					strcase.ToCamel(group+suffix+"N"),
					fmt.Sprintf("%s%d_t", rg, r0),
					fmt.Sprintf("%s%d_t", ig, i0),
					fn.Name,
					setter(fn.Return.Name, "save", fmt.Sprintf("vst1%s_%s%d", rq, typeShort[rg], r0)),
					setter(fn.Args[0].Name, "load", fmt.Sprintf("vld1%s_%s%d", iq, typeShort[ig], i0)),
					o1,
					i1,
				),
			)
			loops[fn.Name] = true
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// loop functions
	if err := writer.WriteToFile("neon/loops.go", func(w io.Writer) error {
		if err := writer.Package(w, "neon"); err != nil {
			return err
		}
		if err := writer.Import(w, []string{
			"github.com/alivanz/go-simd/arm",
		}); err != nil {
			return err
		}
		if err := writer.ImportC(w, func(w io.Writer) error {
			if _, err := io.WriteString(w, "#include <arm_neon.h>"); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
		for _, fn := range result.Functions {
			if !loops[fn.Name] {
				continue
			}
			// add suffix
			fn.Name += "N"
			// write
			fmt.Fprintf(w, "\n")
			if len(fn.Comment) > 0 {
				fmt.Fprintf(w, "// %s\n", fn.Comment)
			} else {
				fmt.Fprintf(w, "// %s\n", fn.Name)
			}
			fmt.Fprintf(w, "//\n")
			fmt.Fprintf(w, "//go:linkname %s %s\n", strcase.ToCamel(fn.Name), strcase.ToCamel(fn.Name))
			fmt.Fprintf(w, "//go:noescape\n")
			fmt.Fprintf(w, "func %s(", strcase.ToCamel(fn.Name))
			if fn.Return != nil {
				var parts = strings.SplitN(strings.TrimSuffix(fn.Return.Name, "_t"), "x", 2)
				fmt.Fprintf(w, "r *arm.%s, ", strcase.ToCamel(parts[0]))
			}
			fmt.Fprintf(w, "%s, n int32)\n", strings.Join(utils.Transform(fn.Args, func(i int, t types.Type) string {
				var parts = strings.SplitN(strings.TrimSuffix(t.Name, "_t"), "x", 2)
				return fmt.Sprintf("v%d *arm.%s", i, strcase.ToCamel(parts[0]))
			}), ", "))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func setter(t string, direct string, def string) string {
	_, _, r1 := parseType(t)
	if r1 == -1 {
		return direct
	}
	return def
}

func parseType(t string) (string, int, int) {
	var (
		group string
	)
	t = strings.TrimSuffix(t, "_t")
	if strings.HasPrefix(t, "uint") {
		group = "uint"
		t = t[4:]
	} else if strings.HasPrefix(t, "int") {
		group = "int"
		t = t[3:]
	} else if strings.HasPrefix(t, "float") {
		group = "float"
		t = t[5:]
	}
	parts := strings.Split(t, "x")
	switch len(parts) {
	case 1:
		w, err := strconv.ParseUint(parts[0], 10, 32)
		if err != nil {
			return "", 0, 0
		}
		return group, int(w), -1
	case 2:
		w, err := strconv.ParseUint(parts[0], 10, 32)
		if err != nil {
			return "", 0, 0
		}
		h, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			return "", 0, 0
		}
		return group, int(w), int(h)
	}
	return "", 0, 0
}

var (
	typeShort = map[string]string{
		"uint":    "u",
		"uint8":   "u8",
		"uint16":  "u16",
		"uint32":  "u32",
		"uint64":  "u64",
		"int":     "s",
		"int8":    "s8",
		"int16":   "s16",
		"int32":   "s32",
		"int64":   "s64",
		"float":   "f",
		"float32": "f32",
		"float64": "f64",
	}
)
