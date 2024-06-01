#include "textflag.h"

/*
 * Return the value of "g.goid".
 * The macro takes an address pointing to "struct g",
 * adds the offset of member "goid", dereferenced it
 * and returns the value present there.
 *
 * Ref(s):
 *   - struct definition of "g" [1]
 *   - go's pseudo-registers: "TLS" and "g" [2]
 *
 * Note(s):
 *   - offset tested on go-1.22, might break on
 *     other versions
 *
 * [1]: https://go.dev/src/runtime/runtime2.go
 * [2]: https://go.dev/doc/asm#architectures
 */
#define g_goid(r)   0x98(r)

/*
 * Return the internal thread ID of the calling
 * goroutine from thread local storage (TLS).
 */  
TEXT ·goidFast(SB), NOSPLIT|NOFRAME, $0x0-0x8
#ifdef GOARCH_386
    MOVL (TLS), CX
    MOVL g_goid(CX), CX
    MOVL CX, ret_lo+0x0(FP)
#endif
#ifdef GOARCH_amd64
    MOVQ (TLS), CX
    MOVQ g_goid(CX), CX
    MOVQ CX, ret+0x0(FP)
#endif
#ifdef GOARCH_amd64p32
    MOVQ (TLS), CX
    MOVQ g_goid(CX), CX
    MOVQ CX, ret+0x0(FP)
#endif
#ifdef GOARCH_arm
    MOVW g, R1
    MOVW g_goid(R1), R1
    MOVW R1, ret_lo+0x0(FP)
#endif
#ifdef GOARCH_arm64
    MOVD g, R1
    MOVD g_goid(R1), R1
    MOVD R1, ret+0x0(FP)
#endif
#ifdef GOARCH_riscv64
    MOV (g), A0
    MOV g_goid(A0), A0
    MOV A0, ret+0x0(FP)
#endif
    RET
