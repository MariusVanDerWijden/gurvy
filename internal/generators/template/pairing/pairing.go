package pairing

const Pairing = `
// FinalExponentiation computes the final expo x**(p**6-1)(p**2+1)(p**4 - p**2 +1)/r
func (curve *Curve) FinalExponentiation(z *{{.Fp12Name}}, _z ...*{{.Fp12Name}}) {{.Fp12Name}} {
	var result {{.Fp12Name}}
	result.Set(z)

	// if additional parameters are provided, multiply them into z
	for _, e := range _z {
		result.Mul(&result, e)
	}

	result.FinalExponentiation(&result)

	return result
}

// MillerLoop Miller loop
func (curve *Curve) MillerLoop(P G1Affine, Q G2Affine, result *{{.Fp12Name}}) *{{.Fp12Name}} {

	// init result
	result.SetOne()

	if P.IsInfinity() || Q.IsInfinity() {
		return result
	}

	// the line goes through QCur and QNext
	var QCur, QNext, QNextNeg G2Jac
	var QNeg G2Affine

	// Stores -Q
	QNeg.Neg(&Q)

	// init QCur with Q
	Q.ToJacobian(&QCur)

	var lEval lineEvalRes

	// Miller loop
	for i := len(curve.loopCounter) - 2; i >= 0; i-- {
		QNext.Set(&QCur)
		QNext.Double()
		QNextNeg.Neg(&QNext)

		result.Square(result)

		// evaluates line though Qcur,2Qcur at P
		lineEvalJac(QCur, QNextNeg, &P, &lEval)
		lEval.mulAssign(result)

		if curve.loopCounter[i] == 1 {
			// evaluates line through 2Qcur, Q at P
			lineEvalAffine(QNext, Q, &P, &lEval)
			lEval.mulAssign(result)

			QNext.AddMixed(&Q)

		} else if curve.loopCounter[i] == -1 {
			// evaluates line through 2Qcur, -Q at P
			lineEvalAffine(QNext, QNeg, &P, &lEval)
			lEval.mulAssign(result)

			QNext.AddMixed(&QNeg)
		}
		QCur.Set(&QNext)
	}

	{{template "ExtraWork" dict "all" . }}

	return result
}

// lineEval computes the evaluation of the line through Q, R (on the twist) at P
// Q, R are in jacobian coordinates
// The case in which Q=R=Infinity is not handled as this doesn't happen in the SNARK pairing
func lineEvalJac(Q, R G2Jac, P *G1Affine, result *lineEvalRes) {
	// converts Q and R to projective coords
	Q.ToProjFromJac()
	R.ToProjFromJac()

	// line eq: w^3*(QyRz-QzRy)x +  w^2*(QzRx - QxRz)y + w^5*(QxRy-QyRxz)
	// result.r1 = QyRz-QzRy
	// result.r0 = QzRx - QxRz
	// result.r2 = QxRy-QyRxz

	result.r1.Mul(&Q.Y, &R.Z)
	result.r0.Mul(&Q.Z, &R.X)
	result.r2.Mul(&Q.X, &R.Y)

	Q.Z.Mul(&Q.Z, &R.Y)
	Q.X.Mul(&Q.X, &R.Z)
	Q.Y.Mul(&Q.Y, &R.X)

	result.r1.Sub(&result.r1, &Q.Z)
	result.r0.Sub(&result.r0, &Q.X)
	result.r2.Sub(&result.r2, &Q.Y)

	// multiply P.Z by coeffs[2] in case P is infinity
	result.r1.MulByElement(&result.r1, &P.X)
	result.r0.MulByElement(&result.r0, &P.Y)
	//result.r2.MulByElement(&result.r2, &P.Z)
}

// Same as above but R is in affine coords
func lineEvalAffine(Q G2Jac, R G2Affine, P *G1Affine, result *lineEvalRes) {

	// converts Q and R to projective coords
	Q.ToProjFromJac()

	// line eq: w^3*(QyRz-QzRy)x +  w^2*(QzRx - QxRz)y + w^5*(QxRy-QyRxz)
	// result.r1 = QyRz-QzRy
	// result.r0 = QzRx - QxRz
	// result.r2 = QxRy-QyRxz

	result.r1.Set(&Q.Y)
	result.r0.Mul(&Q.Z, &R.X)
	result.r2.Mul(&Q.X, &R.Y)

	Q.Z.Mul(&Q.Z, &R.Y)
	Q.Y.Mul(&Q.Y, &R.X)

	result.r1.Sub(&result.r1, &Q.Z)
	result.r0.Sub(&result.r0, &Q.X)
	result.r2.Sub(&result.r2, &Q.Y)

	// multiply P.Z by coeffs[2] in case P is infinity
	result.r1.MulByElement(&result.r1, &P.X)
	result.r0.MulByElement(&result.r0, &P.Y)
	// result.r2.MulByElement(&result.r2, &P.Z)
}

type lineEvalRes struct {
	r0 {{.Fp2Name}} // c0.b1
	r1 {{.Fp2Name}} // c1.b1
	r2 {{.Fp2Name}} // c1.b2
}

func (l *lineEvalRes) mulAssign(z *E12) *E12 {

	{{template "MulAssign" dict "all" . }}

	return z
}

`

// ExtraWork extra operations needed when the loop shortening is used (cf Vecauteren, Optimal Pairing)
const ExtraWork = `
{{define "ExtraWork" }}
	{{if eq $.all.Fpackage "bn256" }}
		// cf https://eprint.iacr.org/2010/354.pdf for instance for optimal Ate Pairing
		var Q1, Q2 G2Affine

		//Q1 = Frob(Q)
		Q1.X.Conjugate(&Q.X).MulByNonResiduePower2(&Q1.X)
		Q1.Y.Conjugate(&Q.Y).MulByNonResiduePower3(&Q1.Y)

		// Q2 = -Frob2(Q)
		Q2.X.MulByNonResiduePowerSquarE2(&Q.X)
		Q2.Y.MulByNonResiduePowerSquare3(&Q.Y).Neg(&Q2.Y)

		lineEvalAffine(QCur, Q1, &P, &lEval)
		lEval.mulAssign(result)

		QCur.AddMixed(&Q1)

		lineEvalAffine(QCur, Q2, &P, &lEval)
		lEval.mulAssign(result)
	{{end}}
{{- end}}
`

// MulAssign multiplies the result of a line evalution to a E12 elmt.
// The line evaluation result is sparse therefore there is a special optimized method to handle this case.
const MulAssign = `
{{define "MulAssign" }}
	{{if eq $.all.Fpackage "bn256" }}	
		var a, b, c E12
		a.MulByVW(z, &l.r1)
		b.MulByV(z, &l.r0)
		c.MulByV2W(z, &l.r2)
		z.Add(&a, &b).Add(z, &c)
	{{else if eq $.all.Fpackage "bls377" }}
		var a, b, c E12
		a.MulByVW(z, &l.r1)
		b.MulByV(z, &l.r0)
		c.MulByV2W(z, &l.r2)
		z.Add(&a, &b).Add(z, &c)
	{{else if eq $.all.Fpackage "bls381" }}
		var a, b, c E12
		a.MulByVWNRInv(z, &l.r1)
		b.MulByV2NRInv(z, &l.r0)
		c.MulByWNRInv(z, &l.r2)
		z.Add(&a, &b).Add(z, &c)
	{{end}}
{{end}}
`
