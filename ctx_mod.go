package nlp

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

func (ctx *Ctx[T]) WithModifier(mod Modifier[T]) *Ctx[T] {
	ctx.mod = append(ctx.mod, mod)
	ctx.SetBit(flagMod, false)
	return ctx
}

func (ctx *Ctx[T]) Modify() *Ctx[T] {
	if ctx.chkSrcErr() {
		return ctx
	}
	if ctx.CheckBit(flagMod) {
		return ctx
	}
	defer ctx.SetBit(flagMod, true)
	for i := 0; i < len(ctx.cln); i++ {
		ctx.bufR = ctx.mod[i].AppendModify(ctx.bufR[:0], byteseq.B2Q[T](ctx.buf))
		ctx.buf = fastconv.AppendR2B(ctx.buf[:0], ctx.bufR)
	}
	return ctx
}

func (ctx *Ctx[T]) ModifyT(t T) T {
	return ctx.SetText(t).
		Modify().
		GetText()
}

func (ctx *Ctx[T]) ResetModifiers() *Ctx[T] {
	ctx.mod = ctx.mod[:0]
	return ctx
}
