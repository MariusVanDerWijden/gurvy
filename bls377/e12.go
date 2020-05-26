// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gurvy/internal/generators DO NOT EDIT

package bls377

// Code generated by internal/fp12 DO NOT EDIT

import (
	"github.com/consensys/gurvy/bls377/fp"
)

// E12 is a degree-two finite field extension of fp6:
// C0 + C1w where w^3-v is irrep in fp6

// fp2, fp12 are both quadratic field extensions
// template code is duplicated in fp2, fp12
// TODO make an abstract quadratic extension template

type E12 struct {
	C0, C1 E6
}

// Equal compares two E12 elements
// TODO can this be deleted?
func (z *E12) Equal(x *E12) bool {
	return z.C0.Equal(&x.C0) && z.C1.Equal(&x.C1)
}

// String puts E12 in string form
func (z *E12) String() string {
	return (z.C0.String() + "+(" + z.C1.String() + ")*w")
}

// SetString sets a E12 from string
func (z *E12) SetString(s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11 string) *E12 {
	z.C0.SetString(s0, s1, s2, s3, s4, s5)
	z.C1.SetString(s6, s7, s8, s9, s10, s11)
	return z
}

// Set copies x into z and returns z
func (z *E12) Set(x *E12) *E12 {
	z.C0 = x.C0
	z.C1 = x.C1
	return z
}

// SetOne sets z to 1 in E12 in Montgomery form and returns z
func (z *E12) SetOne() *E12 {
	z.C0.B0.A0.SetOne()
	z.C0.B0.A1.SetZero()
	z.C0.B1.A0.SetZero()
	z.C0.B1.A1.SetZero()
	z.C0.B2.A0.SetZero()
	z.C0.B2.A1.SetZero()
	z.C1.B0.A0.SetZero()
	z.C1.B0.A1.SetZero()
	z.C1.B1.A0.SetZero()
	z.C1.B1.A1.SetZero()
	z.C1.B2.A0.SetZero()
	z.C1.B2.A1.SetZero()
	return z
}

// ToMont converts to Mont form
// TODO can this be deleted?
func (z *E12) ToMont() *E12 {
	z.C0.ToMont()
	z.C1.ToMont()
	return z
}

// FromMont converts from Mont form
// TODO can this be deleted?
func (z *E12) FromMont() *E12 {
	z.C0.FromMont()
	z.C1.FromMont()
	return z
}

// Add set z=x+y in E12 and return z
func (z *E12) Add(x, y *E12) *E12 {
	z.C0.Add(&x.C0, &y.C0)
	z.C1.Add(&x.C1, &y.C1)
	return z
}

// Sub set z=x-y in E12 and return z
func (z *E12) Sub(x, y *E12) *E12 {
	z.C0.Sub(&x.C0, &y.C0)
	z.C1.Sub(&x.C1, &y.C1)
	return z
}

// SetRandom used only in tests
// TODO eliminate this method!
func (z *E12) SetRandom() *E12 {
	z.C0.B0.A0.SetRandom()
	z.C0.B0.A1.SetRandom()
	z.C0.B1.A0.SetRandom()
	z.C0.B1.A1.SetRandom()
	z.C0.B2.A0.SetRandom()
	z.C0.B2.A1.SetRandom()
	z.C1.B0.A0.SetRandom()
	z.C1.B0.A1.SetRandom()
	z.C1.B1.A0.SetRandom()
	z.C1.B1.A1.SetRandom()
	z.C1.B2.A0.SetRandom()
	z.C1.B2.A1.SetRandom()
	return z
}

// Mul set z=x*y in E12 and return z
func (z *E12) Mul(x, y *E12) *E12 {
	// Algorithm 20 from https://eprint.iacr.org/2010/354.pdf

	var t0, t1, xSum, ySum E6

	t0.Mul(&x.C0, &y.C0) // step 1
	t1.Mul(&x.C1, &y.C1) // step 2

	// finish processing input in case z==x or y
	xSum.Add(&x.C0, &x.C1)
	ySum.Add(&y.C0, &y.C1)

	// step 3
	{ // begin: inline z.C0.MulByNonResidue(&t1)
		var result E6
		result.B1.Set(&(&t1).B0)
		result.B2.Set(&(&t1).B1)
		{ // begin: inline result.B0.MulByNonResidue(&(&t1).B2)
			buf := (&(&t1).B2).A0
			{ // begin: inline MulByNonResidue(&(result.B0).A0, &(&(&t1).B2).A1)
				buf := *(&(&(&t1).B2).A1)
				(&(result.B0).A0).Double(&buf).Double(&(result.B0).A0).AddAssign(&buf)
			} // end: inline MulByNonResidue(&(result.B0).A0, &(&(&t1).B2).A1)
			(result.B0).A1 = buf
		} // end: inline result.B0.MulByNonResidue(&(&t1).B2)
		z.C0.Set(&result)
	} // end: inline z.C0.MulByNonResidue(&t1)
	z.C0.Add(&z.C0, &t0)

	// step 4
	z.C1.Mul(&xSum, &ySum).
		Sub(&z.C1, &t0).
		Sub(&z.C1, &t1)

	return z
}

// Square set z=x*x in E12 and return z
func (z *E12) Square(x *E12) *E12 {
	// TODO implement Algorithm 22 from https://eprint.iacr.org/2010/354.pdf
	// or the complex method from fp2
	// for now do it the dumb way
	var b0, b1 E6

	b0.Square(&x.C0)
	b1.Square(&x.C1)
	{ // begin: inline b1.MulByNonResidue(&b1)
		var result E6
		result.B1.Set(&(&b1).B0)
		result.B2.Set(&(&b1).B1)
		{ // begin: inline result.B0.MulByNonResidue(&(&b1).B2)
			buf := (&(&b1).B2).A0
			{ // begin: inline MulByNonResidue(&(result.B0).A0, &(&(&b1).B2).A1)
				buf := *(&(&(&b1).B2).A1)
				(&(result.B0).A0).Double(&buf).Double(&(result.B0).A0).AddAssign(&buf)
			} // end: inline MulByNonResidue(&(result.B0).A0, &(&(&b1).B2).A1)
			(result.B0).A1 = buf
		} // end: inline result.B0.MulByNonResidue(&(&b1).B2)
		b1.Set(&result)
	} // end: inline b1.MulByNonResidue(&b1)
	b1.Add(&b0, &b1)

	z.C1.Mul(&x.C0, &x.C1).Double(&z.C1)
	z.C0 = b1

	return z
}

// Inverse set z to the inverse of x in E12 and return z
func (z *E12) Inverse(x *E12) *E12 {
	// Algorithm 23 from https://eprint.iacr.org/2010/354.pdf

	var t [2]E6

	t[0].Square(&x.C0) // step 1
	t[1].Square(&x.C1) // step 2
	{                  // step 3
		var buf E6
		{ // begin: inline buf.MulByNonResidue(&t[1])
			var result E6
			result.B1.Set(&(&t[1]).B0)
			result.B2.Set(&(&t[1]).B1)
			{ // begin: inline result.B0.MulByNonResidue(&(&t[1]).B2)
				buf := (&(&t[1]).B2).A0
				{ // begin: inline MulByNonResidue(&(result.B0).A0, &(&(&t[1]).B2).A1)
					buf := *(&(&(&t[1]).B2).A1)
					(&(result.B0).A0).Double(&buf).Double(&(result.B0).A0).AddAssign(&buf)
				} // end: inline MulByNonResidue(&(result.B0).A0, &(&(&t[1]).B2).A1)
				(result.B0).A1 = buf
			} // end: inline result.B0.MulByNonResidue(&(&t[1]).B2)
			buf.Set(&result)
		} // end: inline buf.MulByNonResidue(&t[1])
		t[0].Sub(&t[0], &buf)
	}
	t[1].Inverse(&t[0])               // step 4
	z.C0.Mul(&x.C0, &t[1])            // step 5
	z.C1.Mul(&x.C1, &t[1]).Neg(&z.C1) // step 6

	return z
}

// InverseUnitary inverse a unitary element
// TODO deprecate in favour of Conjugate
func (z *E12) InverseUnitary(x *E12) *E12 {
	return z.Conjugate(x)
}

// Conjugate set z to (x.C0, -x.C1) and return z
func (z *E12) Conjugate(x *E12) *E12 {
	z.Set(x)
	z.C1.Neg(&z.C1)
	return z
}

// MulByVW set z to x*(y*v*w) and return z
// here y*v*w means the E12 element with C1.B1=y and all other components 0
func (z *E12) MulByVW(x *E12, y *E2) *E12 {
	var result E12
	var yNR E2

	{ // begin: inline yNR.MulByNonResidue(y)
		buf := (y).A0
		{ // begin: inline MulByNonResidue(&(yNR).A0, &(y).A1)
			buf := *(&(y).A1)
			(&(yNR).A0).Double(&buf).Double(&(yNR).A0).AddAssign(&buf)
		} // end: inline MulByNonResidue(&(yNR).A0, &(y).A1)
		(yNR).A1 = buf
	} // end: inline yNR.MulByNonResidue(y)
	result.C0.B0.Mul(&x.C1.B1, &yNR)
	result.C0.B1.Mul(&x.C1.B2, &yNR)
	result.C0.B2.Mul(&x.C1.B0, y)
	result.C1.B0.Mul(&x.C0.B2, &yNR)
	result.C1.B1.Mul(&x.C0.B0, y)
	result.C1.B2.Mul(&x.C0.B1, y)
	z.Set(&result)
	return z
}

// MulByV set z to x*(y*v) and return z
// here y*v means the E12 element with C0.B1=y and all other components 0
func (z *E12) MulByV(x *E12, y *E2) *E12 {
	var result E12
	var yNR E2

	{ // begin: inline yNR.MulByNonResidue(y)
		buf := (y).A0
		{ // begin: inline MulByNonResidue(&(yNR).A0, &(y).A1)
			buf := *(&(y).A1)
			(&(yNR).A0).Double(&buf).Double(&(yNR).A0).AddAssign(&buf)
		} // end: inline MulByNonResidue(&(yNR).A0, &(y).A1)
		(yNR).A1 = buf
	} // end: inline yNR.MulByNonResidue(y)
	result.C0.B0.Mul(&x.C0.B2, &yNR)
	result.C0.B1.Mul(&x.C0.B0, y)
	result.C0.B2.Mul(&x.C0.B1, y)
	result.C1.B0.Mul(&x.C1.B2, &yNR)
	result.C1.B1.Mul(&x.C1.B0, y)
	result.C1.B2.Mul(&x.C1.B1, y)
	z.Set(&result)
	return z
}

// MulByV2W set z to x*(y*v^2*w) and return z
// here y*v^2*w means the E12 element with C1.B2=y and all other components 0
func (z *E12) MulByV2W(x *E12, y *E2) *E12 {
	var result E12
	var yNR E2

	{ // begin: inline yNR.MulByNonResidue(y)
		buf := (y).A0
		{ // begin: inline MulByNonResidue(&(yNR).A0, &(y).A1)
			buf := *(&(y).A1)
			(&(yNR).A0).Double(&buf).Double(&(yNR).A0).AddAssign(&buf)
		} // end: inline MulByNonResidue(&(yNR).A0, &(y).A1)
		(yNR).A1 = buf
	} // end: inline yNR.MulByNonResidue(y)
	result.C0.B0.Mul(&x.C1.B0, &yNR)
	result.C0.B1.Mul(&x.C1.B1, &yNR)
	result.C0.B2.Mul(&x.C1.B2, &yNR)
	result.C1.B0.Mul(&x.C0.B1, &yNR)
	result.C1.B1.Mul(&x.C0.B2, &yNR)
	result.C1.B2.Mul(&x.C0.B0, y)
	z.Set(&result)
	return z
}

// MulByV2NRInv set z to x*(y*v^2*(0,1)^{-1}) and return z
// here y*v^2 means the E12 element with C0.B2=y and all other components 0
func (z *E12) MulByV2NRInv(x *E12, y *E2) *E12 {
	var result E12
	var yNRInv E2

	{ // begin: inline yNRInv.MulByNonResidueInv(y)
		buf := (y).A1
		{ // begin: inline MulByNonResidueInv(&(yNRInv).A1, &(y).A0)
			nrinv := fp.Element{
				330620507644336508,
				9878087358076053079,
				11461392860540703536,
				6973035786057818995,
				8846909097162646007,
				104838758629667239,
			}
			(&(yNRInv).A1).Mul(&(y).A0, &nrinv)
		} // end: inline MulByNonResidueInv(&(yNRInv).A1, &(y).A0)
		(yNRInv).A0 = buf
	} // end: inline yNRInv.MulByNonResidueInv(y)

	result.C0.B0.Mul(&x.C0.B1, y)
	result.C0.B1.Mul(&x.C0.B2, y)
	result.C0.B2.Mul(&x.C0.B0, &yNRInv)

	result.C1.B0.Mul(&x.C1.B1, y)
	result.C1.B1.Mul(&x.C1.B2, y)
	result.C1.B2.Mul(&x.C1.B0, &yNRInv)

	z.Set(&result)
	return z
}

// MulByVWNRInv set z to x*(y*v*w*(0,1)^{-1}) and return z
// here y*v*w means the E12 element with C1.B1=y and all other components 0
func (z *E12) MulByVWNRInv(x *E12, y *E2) *E12 {
	var result E12
	var yNRInv E2

	{ // begin: inline yNRInv.MulByNonResidueInv(y)
		buf := (y).A1
		{ // begin: inline MulByNonResidueInv(&(yNRInv).A1, &(y).A0)
			nrinv := fp.Element{
				330620507644336508,
				9878087358076053079,
				11461392860540703536,
				6973035786057818995,
				8846909097162646007,
				104838758629667239,
			}
			(&(yNRInv).A1).Mul(&(y).A0, &nrinv)
		} // end: inline MulByNonResidueInv(&(yNRInv).A1, &(y).A0)
		(yNRInv).A0 = buf
	} // end: inline yNRInv.MulByNonResidueInv(y)

	result.C0.B0.Mul(&x.C1.B1, y)
	result.C0.B1.Mul(&x.C1.B2, y)
	result.C0.B2.Mul(&x.C1.B0, &yNRInv)

	result.C1.B0.Mul(&x.C0.B2, y)
	result.C1.B1.Mul(&x.C0.B0, &yNRInv)
	result.C1.B2.Mul(&x.C0.B1, &yNRInv)

	z.Set(&result)
	return z
}

// MulByWNRInv set z to x*(y*w*(0,1)^{-1}) and return z
// here y*w means the E12 element with C1.B0=y and all other components 0
func (z *E12) MulByWNRInv(x *E12, y *E2) *E12 {
	var result E12
	var yNRInv E2

	{ // begin: inline yNRInv.MulByNonResidueInv(y)
		buf := (y).A1
		{ // begin: inline MulByNonResidueInv(&(yNRInv).A1, &(y).A0)
			nrinv := fp.Element{
				330620507644336508,
				9878087358076053079,
				11461392860540703536,
				6973035786057818995,
				8846909097162646007,
				104838758629667239,
			}
			(&(yNRInv).A1).Mul(&(y).A0, &nrinv)
		} // end: inline MulByNonResidueInv(&(yNRInv).A1, &(y).A0)
		(yNRInv).A0 = buf
	} // end: inline yNRInv.MulByNonResidueInv(y)

	result.C0.B0.Mul(&x.C1.B2, y)
	result.C0.B1.Mul(&x.C1.B0, &yNRInv)
	result.C0.B2.Mul(&x.C1.B1, &yNRInv)

	result.C1.B0.Mul(&x.C0.B0, &yNRInv)
	result.C1.B1.Mul(&x.C0.B1, &yNRInv)
	result.C1.B2.Mul(&x.C0.B2, &yNRInv)

	z.Set(&result)
	return z
}

// MulByNonResidue multiplies a E6 by ((0,0),(1,0),(0,0))
func (z *E6) MulByNonResidue(x *E6) *E6 {
	var result E6
	result.B1.Set(&(x).B0)
	result.B2.Set(&(x).B1)
	{ // begin: inline result.B0.MulByNonResidue(&(x).B2)
		buf := (&(x).B2).A0
		{ // begin: inline MulByNonResidue(&(result.B0).A0, &(&(x).B2).A1)
			buf := *(&(&(x).B2).A1)
			(&(result.B0).A0).Double(&buf).Double(&(result.B0).A0).AddAssign(&buf)
		} // end: inline MulByNonResidue(&(result.B0).A0, &(&(x).B2).A1)
		(result.B0).A1 = buf
	} // end: inline result.B0.MulByNonResidue(&(x).B2)
	z.Set(&result)
	return z
}

// Frobenius set z to Frobenius(x) in E12 and return z
func (z *E12) Frobenius(x *E12) *E12 {
	// Algorithm 28 from https://eprint.iacr.org/2010/354.pdf (beware typos!)
	var t [6]E2

	// Frobenius acts on fp2 by conjugation
	t[0].Conjugate(&x.C0.B0)
	t[1].Conjugate(&x.C0.B1)
	t[2].Conjugate(&x.C0.B2)
	t[3].Conjugate(&x.C1.B0)
	t[4].Conjugate(&x.C1.B1)
	t[5].Conjugate(&x.C1.B2)

	t[1].MulByNonResiduePower2(&t[1])
	t[2].MulByNonResiduePower4(&t[2])
	t[3].MulByNonResiduePower1(&t[3])
	t[4].MulByNonResiduePower3(&t[4])
	t[5].MulByNonResiduePower5(&t[5])

	z.C0.B0 = t[0]
	z.C0.B1 = t[1]
	z.C0.B2 = t[2]
	z.C1.B0 = t[3]
	z.C1.B1 = t[4]
	z.C1.B2 = t[5]

	return z
}

// FrobeniusSquare set z to Frobenius^2(x) in E12 and return z
func (z *E12) FrobeniusSquare(x *E12) *E12 {
	// Algorithm 29 from https://eprint.iacr.org/2010/354.pdf (beware typos!)
	var t [6]E2

	t[1].MulByNonResiduePowerSquarE2(&x.C0.B1)
	t[2].MulByNonResiduePowerSquare4(&x.C0.B2)
	t[3].MulByNonResiduePowerSquare1(&x.C1.B0)
	t[4].MulByNonResiduePowerSquare3(&x.C1.B1)
	t[5].MulByNonResiduePowerSquare5(&x.C1.B2)

	z.C0.B0 = x.C0.B0
	z.C0.B1 = t[1]
	z.C0.B2 = t[2]
	z.C1.B0 = t[3]
	z.C1.B1 = t[4]
	z.C1.B2 = t[5]

	return z
}

// FrobeniusCube set z to Frobenius^3(x) in E12 and return z
func (z *E12) FrobeniusCube(x *E12) *E12 {
	// Algorithm 30 from https://eprint.iacr.org/2010/354.pdf (beware typos!)
	var t [6]E2

	// Frobenius^3 acts on fp2 by conjugation
	t[0].Conjugate(&x.C0.B0)
	t[1].Conjugate(&x.C0.B1)
	t[2].Conjugate(&x.C0.B2)
	t[3].Conjugate(&x.C1.B0)
	t[4].Conjugate(&x.C1.B1)
	t[5].Conjugate(&x.C1.B2)

	t[1].MulByNonResiduePowerCubE2(&t[1])
	t[2].MulByNonResiduePowerCube4(&t[2])
	t[3].MulByNonResiduePowerCube1(&t[3])
	t[4].MulByNonResiduePowerCube3(&t[4])
	t[5].MulByNonResiduePowerCube5(&t[5])

	z.C0.B0 = t[0]
	z.C0.B1 = t[1]
	z.C0.B2 = t[2]
	z.C1.B0 = t[3]
	z.C1.B1 = t[4]
	z.C1.B2 = t[5]

	return z
}

// MulByNonResiduePower1 set z=x*(0,1)^(1*(p-1)/6) and return z
func (z *E2) MulByNonResiduePower1(x *E2) *E2 {
	// (0,1)^(1*(p-1)/6)
	// 92949345220277864758624960506473182677953048909283248980960104381795901929519566951595905490535835115111760994353
	b := fp.Element{
		7981638599956744862,
		11830407261614897732,
		6308788297503259939,
		10596665404780565693,
		11693741422477421038,
		61545186993886319,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePower2 set z=x*(0,1)^(2*(p-1)/6) and return z
func (z *E2) MulByNonResiduePower2(x *E2) *E2 {
	// (0,1)^(2*(p-1)/6)
	// 80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410946
	b := fp.Element{
		6382252053795993818,
		1383562296554596171,
		11197251941974877903,
		6684509567199238270,
		6699184357838251020,
		19987743694136192,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePower3 set z=x*(0,1)^(3*(p-1)/6) and return z
func (z *E2) MulByNonResiduePower3(x *E2) *E2 {
	// (0,1)^(3*(p-1)/6)
	// 216465761340224619389371505802605247630151569547285782856803747159100223055385581585702401816380679166954762214499
	b := fp.Element{
		10965161018967488287,
		18251363109856037426,
		7036083669251591763,
		16109345360066746489,
		4679973768683352764,
		96952949334633821,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePower4 set z=x*(0,1)^(4*(p-1)/6) and return z
func (z *E2) MulByNonResiduePower4(x *E2) *E2 {
	// (0,1)^(4*(p-1)/6)
	// 80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410945
	b := fp.Element{
		15766275933608376691,
		15635974902606112666,
		1934946774703877852,
		18129354943882397960,
		15437979634065614942,
		101285514078273488,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePower5 set z=x*(0,1)^(5*(p-1)/6) and return z
func (z *E2) MulByNonResiduePower5(x *E2) *E2 {
	// (0,1)^(5*(p-1)/6)
	// 123516416119946754630746545296132064952198520638002533875843642777304321125866014634106496325844844051843001220146
	b := fp.Element{
		2983522419010743425,
		6420955848241139694,
		727295371748331824,
		5512679955286180796,
		11432976419915483342,
		35407762340747501,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerSquare1 set z=x*(0,1)^(1*(p^2-1)/6) and return z
func (z *E2) MulByNonResiduePowerSquare1(x *E2) *E2 {
	// (0,1)^(1*(p^2-1)/6)
	// 80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410946
	b := fp.Element{
		6382252053795993818,
		1383562296554596171,
		11197251941974877903,
		6684509567199238270,
		6699184357838251020,
		19987743694136192,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerSquarE2 set z=x*(0,1)^(2*(p^2-1)/6) and return z
func (z *E2) MulByNonResiduePowerSquarE2(x *E2) *E2 {
	// (0,1)^(2*(p^2-1)/6)
	// 80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410945
	b := fp.Element{
		15766275933608376691,
		15635974902606112666,
		1934946774703877852,
		18129354943882397960,
		15437979634065614942,
		101285514078273488,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerSquare3 set z=x*(0,1)^(3*(p^2-1)/6) and return z
func (z *E2) MulByNonResiduePowerSquare3(x *E2) *E2 {
	// (0,1)^(3*(p^2-1)/6)
	// 258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458176
	b := fp.Element{
		9384023879812382873,
		14252412606051516495,
		9184438906438551565,
		11444845376683159689,
		8738795276227363922,
		81297770384137296,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerSquare4 set z=x*(0,1)^(4*(p^2-1)/6) and return z
func (z *E2) MulByNonResiduePowerSquare4(x *E2) *E2 {
	// (0,1)^(4*(p^2-1)/6)
	// 258664426012969093929703085429980814127835149614277183275038967946009968870203535512256352201271898244626862047231
	b := fp.Element{
		3203870859294639911,
		276961138506029237,
		9479726329337356593,
		13645541738420943632,
		7584832609311778094,
		101110569012358506,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerSquare5 set z=x*(0,1)^(5*(p^2-1)/6) and return z
func (z *E2) MulByNonResiduePowerSquare5(x *E2) *E2 {
	// (0,1)^(5*(p^2-1)/6)
	// 258664426012969093929703085429980814127835149614277183275038967946009968870203535512256352201271898244626862047232
	b := fp.Element{
		12266591053191808654,
		4471292606164064357,
		295287422898805027,
		2200696361737783943,
		17292781406793965788,
		19812798628221209,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerCube1 set z=x*(0,1)^(1*(p^3-1)/6) and return z
func (z *E2) MulByNonResiduePowerCube1(x *E2) *E2 {
	// (0,1)^(1*(p^3-1)/6)
	// 216465761340224619389371505802605247630151569547285782856803747159100223055385581585702401816380679166954762214499
	b := fp.Element{
		10965161018967488287,
		18251363109856037426,
		7036083669251591763,
		16109345360066746489,
		4679973768683352764,
		96952949334633821,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerCubE2 set z=x*(0,1)^(2*(p^3-1)/6) and return z
func (z *E2) MulByNonResiduePowerCubE2(x *E2) *E2 {
	// (0,1)^(2*(p^3-1)/6)
	// 258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458176
	b := fp.Element{
		9384023879812382873,
		14252412606051516495,
		9184438906438551565,
		11444845376683159689,
		8738795276227363922,
		81297770384137296,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerCube3 set z=x*(0,1)^(3*(p^3-1)/6) and return z
func (z *E2) MulByNonResiduePowerCube3(x *E2) *E2 {
	// (0,1)^(3*(p^3-1)/6)
	// 42198664672744474621281227892288285906241943207628877683080515507620245292955241189266486323192680957485559243678
	b := fp.Element{
		17067705967832697058,
		1855904398914139597,
		13640894602060642732,
		4220705945553435413,
		9604043198466676350,
		24145363371860877,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResiduePowerCube4 set z=x*(0,1)^(4*(p^3-1)/6) and return z
func (z *E2) MulByNonResiduePowerCube4(x *E2) *E2 {
	// (0,1)^(4*(p^3-1)/6)
	// the value is 1; nothing to do
	return z
}

// MulByNonResiduePowerCube5 set z=x*(0,1)^(5*(p^3-1)/6) and return z
func (z *E2) MulByNonResiduePowerCube5(x *E2) *E2 {
	// (0,1)^(5*(p^3-1)/6)
	// 216465761340224619389371505802605247630151569547285782856803747159100223055385581585702401816380679166954762214499
	b := fp.Element{
		10965161018967488287,
		18251363109856037426,
		7036083669251591763,
		16109345360066746489,
		4679973768683352764,
		96952949334633821,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

const tAbsVal uint64 = 9586122913090633729

// Expt set z to x^t in E12 and return z
// TODO make a ExptAssign method that assigns the result to self; then this method can assert fail if z != x
// TODO Expt is the only method that depends on tAbsVal.  The rest of the tower does not depend on this value.  Logically, Expt should be separated from the rest of the tower.
func (z *E12) Expt(x *E12) *E12 {
	// TODO what if x==0?
	// TODO make this match Element.Exp: x is a non-pointer?

	// tAbsVal in binary: 1000010100001000110000000000000000000000000000000000000000000001
	// drop the low 46 bits (all 0 except the least significant bit): 100001010000100011 = 136227
	// Shortest addition chains can be found at https://wwwhomes.uni-bielefeld.de/achim/addition_chain.html

	var result, x33 E12

	// a shortest addition chain for 136227
	result.Set(x)             // 0                1
	result.Square(&result)    // 1( 0)            2
	result.Square(&result)    // 2( 1)            4
	result.Square(&result)    // 3( 2)            8
	result.Square(&result)    // 4( 3)           16
	result.Square(&result)    // 5( 4)           32
	result.Mul(&result, x)    // 6( 5, 0)        33
	x33.Set(&result)          // save x33 for step 14
	result.Square(&result)    // 7( 6)           66
	result.Square(&result)    // 8( 7)          132
	result.Square(&result)    // 9( 8)          264
	result.Square(&result)    // 10( 9)          528
	result.Square(&result)    // 11(10)         1056
	result.Square(&result)    // 12(11)         2112
	result.Square(&result)    // 13(12)         4224
	result.Mul(&result, &x33) // 14(13, 6)      4257
	result.Square(&result)    // 15(14)         8514
	result.Square(&result)    // 16(15)        17028
	result.Square(&result)    // 17(16)        34056
	result.Square(&result)    // 18(17)        68112
	result.Mul(&result, x)    // 19(18, 0)     68113
	result.Square(&result)    // 20(19)       136226
	result.Mul(&result, x)    // 21(20, 0)    136227

	// the remaining 46 bits
	for i := 0; i < 46; i++ {
		result.Square(&result)
	}
	result.Mul(&result, x)

	z.Set(&result)
	return z
}

// FinalExponentiation computes the final expo x**((p**12 - 1)/r)
func (z *E12) FinalExponentiation(x *E12) *E12 {
	// For BLS curves use Section 3 of https://eprint.iacr.org/2016/130.pdf; "hard part" is Algorithm 1 of https://eprint.iacr.org/2016/130.pdf
	var result E12
	result.Set(x)

	// memalloc
	var t [6]E12

	// buf = x**(p^6-1)
	t[0].FrobeniusCube(&result).
		FrobeniusCube(&t[0])

	result.Inverse(&result)
	t[0].Mul(&t[0], &result)

	// x = (x**(p^6-1)) ^(p^2+1)
	result.FrobeniusSquare(&t[0]).
		Mul(&result, &t[0])

	// hard part (up to permutation)
	// performs the hard part of the final expo
	// Algorithm 1 of https://eprint.iacr.org/2016/130.pdf
	// The result is the same as p**4-p**2+1/r, but up to permutation (it's 3* (p**4 -p**2 +1 /r)), ok since r=1 mod 3)

	t[0].InverseUnitary(&result).Square(&t[0])
	t[5].Expt(&result)
	t[1].Square(&t[5])
	t[3].Mul(&t[0], &t[5])

	t[0].Expt(&t[3])
	t[2].Expt(&t[0])
	t[4].Expt(&t[2])

	t[4].Mul(&t[1], &t[4])
	t[1].Expt(&t[4])
	t[3].InverseUnitary(&t[3])
	t[1].Mul(&t[3], &t[1])
	t[1].Mul(&t[1], &result)

	t[0].Mul(&t[0], &result)
	t[0].FrobeniusCube(&t[0])

	t[3].InverseUnitary(&result)
	t[4].Mul(&t[3], &t[4])
	t[4].Frobenius(&t[4])

	t[5].Mul(&t[2], &t[5])
	t[5].FrobeniusSquare(&t[5])

	t[5].Mul(&t[5], &t[0])
	t[5].Mul(&t[5], &t[4])
	t[5].Mul(&t[5], &t[1])

	result.Set(&t[5])

	z.Set(&result)
	return z
}
